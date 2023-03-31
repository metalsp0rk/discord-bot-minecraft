package config

import(
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "log"
)

var(
  Token string
  BotPrefix string
  config *conf
)


type conf struct {
    Token string `yaml:"Token"`
    BotPrefix string `yaml:"BotPrefix"`
}

func ReadConfig() error {

  yamlFile, err := ioutil.ReadFile("/home/metalspork/code/discord-bot-minecraft/config.yaml")
  if err != nil {
    log.Printf("yamlFile.Get err   #%v ", err)
  }
  err = yaml.Unmarshal(yamlFile, &config)
  if err != nil {
    log.Fatalf("Unmarshal: %v", err)
  }

  Token     = config.Token
  BotPrefix = config.BotPrefix
  return nil
}


