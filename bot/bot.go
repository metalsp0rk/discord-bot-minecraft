package bot

import(
  "fmt"
  "github.com/bwmarrin/discordgo"
  "github.com/metalsp0rk/discord-bot-minecraft/config"
)

var BotID string
var goBot *discordgo.Session

func Start() {
  goBot, err := discordgo.New("Bot " + config.Token)
  if err != nil {
    fmt.Println(err.Error())
    return
  }
  // u, err := goBot.User("@me")
  // if err != nil {
  //   fmt.Println(err.Error())
  // }
  // BotID = u.ID

  goBot.AddHandler(messageCreate)
  goBot.Identify.Intents = discordgo.IntentsGuildMessages

  err = goBot.Open()
  if err != nil {
    fmt.Println(err.Error())
    return
  }

  fmt.Println("Bot is running")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
  if m.Author.ID == s.State.User.ID {
    return
  }
  if m.Content == "ping" {
    fmt.Println("ping received, sending pong")
    s.ChannelMessageSend(m.ChannelID, "Pong!")
  }
  if m.Content == "pong" {
    fmt.Println("pong received, sending ping")
    s.ChannelMessageSend(m.ChannelID, "Ping!")
  }
}
