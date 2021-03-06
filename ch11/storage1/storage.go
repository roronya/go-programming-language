package storage

import (
    "fmt"
    "log"
    "net/smtp"
)

func bytesInUse(username string) int64 { return 0 }

// 電子メールの送信者設定
const sender = "notifications@example.com"
const password = "correcthorsebatterystaple"
const hostname = "smtp.example.com"

const template = `Warning: yo are using %d bytes of storage, %d%% of your quote.`

func CheckQouta(username string) {
  used := bytesInUse(username)
  const quota = 1000000000 // 1GB
  percent := 100 * used / quota
  if percent < 90 {
    return
  }
  msg := fmt.Sprintf(template, used, percent)
  auth := smtp.PlainAuth("", sender, password, hostname)
  err := smtp.SendMail(hostname+":587", auth, sender, []string{username}, []byte(msg))
  if err != nil {
    log.Printf("smtp.SendMail(%s) failed: %s", username, err)
  }
}
