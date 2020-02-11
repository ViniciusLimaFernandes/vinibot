package registry

import (
	"sync"

	"github.com/bwmarrin/discordgo"
)

var mutex sync.Mutex

var commands map[string]func(*discordgo.Session, *discordgo.MessageCreate)

func Register(command string, handler func(*discordgo.Session, *discordgo.MessageCreate)) {
	mutex.Lock()
	defer mutex.Unlock()

	if commands == nil {
		commands = make(map[string]func(*discordgo.Session, *discordgo.MessageCreate))
	}

	commands[command] = handler
}

func Get(command string) func(*discordgo.Session, *discordgo.MessageCreate) {
	return commands[command]
}
