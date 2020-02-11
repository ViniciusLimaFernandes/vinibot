package commands

import ( 
	"vinibot/bot/commands/registry" 
	"github.com/bwmarrin/discordgo" 
)

func init() {
	registry.Register("ping",Ping)
	registry.Register("pong",Pong)
}

func Ping(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID,"Pong!")
}

func Pong(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID,"Ping!")
}