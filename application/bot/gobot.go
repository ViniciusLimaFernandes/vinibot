package bot

import (
	"os"
	"os/signal"
	"strings"
	"syscall"
	"vinibot/data"
	"vinibot/utils"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func Run() {
	botToken := new(data.BotToken)
	token := utils.GetBotToken(botToken)

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Errorln("Error creating Discord session,", err)
	}

	discord.AddHandler(messageCreate)

	err = discord.Open()
	if err != nil {
		log.Errorln("Error opening connection,", err)
	}

	log.Infoln("Bot is now running!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
	log.Infof("\rRIP ViniBot =(")

}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := strings.ToLower(m.Content)

	if m.Author.ID == s.State.User.ID {
		return
	}

	if message == "!ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if message == "!pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
