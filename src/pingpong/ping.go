package pingpong

import (
	"github.com/Meta502/go-discord-pingpong/src/command"
	"github.com/bwmarrin/discordgo"
)

var Ping = command.MessageCommand{
	Name: "ping",
	Alias: map[string]bool{
		"ting": true,
	},
	Handler: func(s *discordgo.Session, m *discordgo.MessageCreate) bool {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
		return true
	},
}
