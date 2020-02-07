package events

import (
	"strings"
	_ "vinibot/bot/commands"
	"vinibot/bot/commands/registry"
	"vinibot/data"
	"vinibot/utils"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	botSettings := new(data.BotSettings)
	messagePrefix := utils.GetBotPrefix(botSettings)
	message := strings.ToLower(m.Content)

	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(message, messagePrefix) {
		return
	}

	message = strings.TrimPrefix(message, messagePrefix)

	registry.Get(message)(s, m)
}
