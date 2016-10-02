// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package main

import (
	"fmt"
	"time"
	
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/debug"
	"golang.org/x/sys/windows/svc/eventlog"
)

var elog debug.Log
var tempStatus int32 = stableTemp
var lastTime time.Time = time.Now()


const (
  stableTemp = iota
  warningTemp
  emergeTemp
)

const (
  warningLimit  = 33.5
  emergeLimit   = 35.0
)

type myservice struct{}

func (m *myservice) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown | svc.AcceptPauseAndContinue
	changes <- svc.Status{State: svc.StartPending}
	fasttick := time.Tick(500 * time.Millisecond)
	slowtick := time.Tick(2 * time.Second)
	tick := fasttick
	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
loop:
	for {
		select {
		case <-tick:
			//beep()
			invokeTemperture("HnsTemp")
			//elog.Info(1, "beep")
		case c := <-r:
			switch c.Cmd {
			case svc.Interrogate:
				changes <- c.CurrentStatus
				// Testing deadlock from https://code.google.com/p/winsvc/issues/detail?id=4
				time.Sleep(100 * time.Millisecond)
				changes <- c.CurrentStatus
			case svc.Stop, svc.Shutdown:
				break loop
			case svc.Pause:
				changes <- svc.Status{State: svc.Paused, Accepts: cmdsAccepted}
				tick = slowtick
			case svc.Continue:
				changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
				tick = fasttick
			default:
				elog.Error(1, fmt.Sprintf("unexpected control request #%d", c))
			}
		}
	}
	changes <- svc.Status{State: svc.StopPending}
	return
}

func log_error(name string, err error) {
  if err == nil {
    return
  }
  elog.Error(1, fmt.Sprintf("%s service failed: %v", name, err))
}

func getStatusString(state int32) string {
  if state == warningTemp {
    return "WARNING"
  } else if state == emergeTemp {
    return "FATAL"
  }
  return "SAFE"
}

func getStatus(temp float64) int32 {
 if temp > emergeLimit {
    return emergeTemp
  } else if temp > warningLimit {
    return warningTemp
  }
  return stableTemp
}

func sendTempertureMail(name string, status int32, temp float64) {
  err := sendMail(name,
    fmt.Sprintf("[%s] >>> Server ROOM %s <<<", name, getStatusString(status)),
    fmt.Sprintf(" Temperture: %.3f", temp))
  log_error(name, err)
}

func invokeTemperture(name string) {
  curTemp := getTemperture()
  curStatus := getStatus(curTemp)
  // ステータスが変更されたので通知
  if tempStatus != curStatus {
    //elog.Info(1,fmt.Sprintf("%s service: invokeTemperture 2", name))
    tempStatus = curStatus
    lastTime = time.Now()
    sendTempertureMail(name,curStatus,curTemp)
    var _ = tempStatus
  }
  duration := time.Now().Sub( lastTime )
  if duration.Hours() > 2.0 {
    //elog.Info(1,fmt.Sprintf("%s service: invokeTemperture 3", name))
    if tempStatus != stableTemp {
      //elog.Info(1,fmt.Sprintf("%s service: invokeTemperture 4", name))
      sendTempertureMail(name,curStatus,curTemp)
      lastTime = time.Now()
    }
  }
}

func runService(name string, isDebug bool) {
	var err error
	if isDebug {
		elog = debug.New(name)
	} else {
		elog, err = eventlog.Open(name)
		if err != nil {
			return
		}
	}
	defer elog.Close()
	elog.Info(1, fmt.Sprintf("starting %s service", name))
	err = initFuncs()
	log_error( name, err )
	defer termFuncs()
	startTemp := getTemperture()
	err = sendMail(name, "[HnsTemp] Start Notice",fmt.Sprintf("Service was started.\r\n\r\nTemperture : %.3f", startTemp)) 
	log_error(name, err)
	run := svc.Run
	if isDebug {
		run = debug.Run
	}
	err = run(name, &myservice{})
	log_error(name, err)
	elog.Info(1, fmt.Sprintf("%s service stopped", name))
	err = sendMail(name, "[HnsTemp] End Notice","Service was stopped") 
	log_error(name, err)
}
