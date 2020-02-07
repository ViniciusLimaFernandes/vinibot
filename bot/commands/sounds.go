package commands

import (
	"vinibot/bot/commands/registry"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func init() {
	registry.Register("bora", Bora)
}

func Bora(s *discordgo.Session, m *discordgo.MessageCreate) {
	filePath := "vinibot/assets/sounds/bora-cumpadi-ajuda-o-doente.mp3"

	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		log.Errorln("Cannot get channel:", err)
		return
	}

	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		log.Errorln("Cannot get the guild:", err)
		return
	}

	log.Infof("I god the Channel ID: %s , and the Guild ID: %s", g.ID, c.ID)

	dgv, err := s.ChannelVoiceJoin(g.ID, c.ID, false, true)
	if err != nil {
		log.Errorln("Failed joining in the channel:", err)
		return
	}

	dgvoice.PlayAudioFile(dgv, filePath, make(chan bool))

	dgv.Close()
}
