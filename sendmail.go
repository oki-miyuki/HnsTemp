package main

import (
	"fmt"
	"time"
	"net/smtp"
)

func sendMail(name string, title string, description string) error {
    from     := "daemon@foo dot com"
    to       := "tweet@foo dot com"
    user     := "smtp user@foo dot com"
    password := "your password"
    subject  := title
    server   := "foo dot com"
    port     := "587"

    emailTemplate := `To: %s
From: %s
Date: %s
Message-ID: <%dhnstemp@daemon.hns.temp>
Subject: %s

%s
`

  body := fmt.Sprintf(emailTemplate, to, from, time.Now().Format(time.RFC1123Z), time.Now().Unix(), subject, description)
  auth := smtp.PlainAuth("", user, password, server)
  err := smtp.SendMail(
    fmt.Sprintf("%s:%s", server, port),
    auth,
    from,
    []string{to},
    []byte(body),
  )
  return err
}

