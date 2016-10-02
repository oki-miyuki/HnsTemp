// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package main

import (
	_ "fmt"
	"syscall"
	"unsafe"
)

var gInitialed bool = false
var gDll *syscall.DLL 
var gFindUsb *syscall.Proc
var gUsbDevice uintptr = 0
var gGetTempHumid *syscall.Proc

func connectUsb() bool {
    var usb_no int = 0
    gUsbDevice, _, _ = gFindUsb.Call(uintptr(unsafe.Pointer(&usb_no)))
    return (gUsbDevice != 0)
}

func initFuncs() error {
	if gInitialed {
	  return nil 
	}
	var err error
	gDll, err = syscall.LoadDLL("USBMeter.DLL")
	if err != nil {
	    //elog.Error(1, fmt.Sprintf("%s service failed: %v", "HnsTemp", err))
	    return err
	}
	gInitialed = true
    gFindUsb, err = gDll.FindProc("_FindUSB@4")
    if err != nil {
	    //elog.Error(1, fmt.Sprintf("%s service failed: %v", "HnsTemp", err))
	    return err
    }
    gGetTempHumid, err = gDll.FindProc("_GetTempHumid@12")
    if err != nil {
	    //elog.Error(1, fmt.Sprintf("%s service failed: %v", "HnsTemp", err))
	    return err
    }
    var _ = gFindUsb
    var _ = gGetTempHumid
    connectUsb()
    var _ = gDll
    return nil
}

func termFuncs() {
	if gInitialed {
		gDll.Release()
		gInitialed = false
	}
	//sendMail(name, "HnsTemp", "Service was stopped."
}

func getTemperture() float64 {
    if gUsbDevice == 0 {
      return 9999.0
    }
    var temp float64
    var hum float64
    gGetTempHumid.Call(gUsbDevice, uintptr(unsafe.Pointer(&temp)), uintptr(unsafe.Pointer(&hum)))
	return temp;
}

