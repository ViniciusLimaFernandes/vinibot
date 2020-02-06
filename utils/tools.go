package utils

import (
	"io/ioutil"
	"log"
	"vinibot/data"

	"gopkg.in/yaml.v2"
)

func GetBotToken(st *data.BotToken) string {
	yamlFile, err := ioutil.ReadFile("./data/config.yaml")
	if err != nil {
		log.Println("Error reading settings file: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, st)
	if err != nil {
		log.Println("Error parsing settings file: %v", err)
	}

	return st.Bot.Auth

}
