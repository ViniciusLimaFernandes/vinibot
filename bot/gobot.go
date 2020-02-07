package bot

import (
	"os"
	"os/signal"
	"syscall"
	"vinibot/bot/events"
	"vinibot/data"
	"vinibot/utils"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func Run() {
	botToken := new(data.BotSettings)
	token := utils.GetBotToken(botToken)

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Errorln("Error creating Discord session,", err)
	}

	discord.AddHandler(events.MessageCreate)

	err = discord.Open()
	if err != nil {
		log.Errorln("Error opening connection,", err)
	}

	utils.WelcomeMessage()
	log.Infoln("Bot is now running!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
	log.Infoln("\rRIP ViniBot =(")

}
