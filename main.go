package main

import (
  "fmt"
  "os"
  "os/signal"
  "github.com/metalsp0rk/discord-bot-minecraft/config"
  "github.com/metalsp0rk/discord-bot-minecraft/bot"
)

func main() {
  err := config.ReadConfig()
  if err != nil {
    fmt.Println(err.Error())
  }
  bot.Start()

  // Keep open till dead
  stop := make(chan os.Signal, 1)
  signal.Notify(stop, os.Interrupt)
  <-stop
  return
}
