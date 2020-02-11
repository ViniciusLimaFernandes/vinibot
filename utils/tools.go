package utils

import (
	"fmt"
	"io/ioutil"
	"vinibot/data"

	log "github.com/sirupsen/logrus"

	"gopkg.in/yaml.v2"
)

func GetBotToken(st *data.BotSettings) string {
	yamlFile, err := ioutil.ReadFile("./data/config.yaml")
	if err != nil {
		log.Errorln("Error reading settings file:", err)
	}

	err = yaml.Unmarshal(yamlFile, st)
	if err != nil {
		log.Errorln("Error parsing settings file:", err)
	}

	return st.Bot.Auth

}

func GetBotPrefix(st *data.BotSettings) string {
	yamlFile, err := ioutil.ReadFile("./data/config.yaml")
	if err != nil {
		log.Println("Error reading settings file: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, st)
	if err != nil {
		log.Println("Error parsing settings file: %v", err)
	}

	return st.Bot.Prefix

}

func WelcomeMessage() {
	fmt.Println(`
	__     __ _         _  ____          _   
	\ \   / /(_) _ __  (_)| __ )   ___  | |_ 
	 \ \ / / | || '_ \ | ||  _ \  / _ \ | __|
	  \ V /  | || | | || || |_) || (_) || |_ 
	   \_/   |_||_| |_||_||____/  \___/  \__|  Version 1.0
											 
									  
	  `)
}
