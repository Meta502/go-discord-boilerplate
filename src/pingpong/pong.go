package pingpong

import (
	"github.com/Meta502/go-discord-pingpong/src/command"
	"github.com/bwmarrin/discordgo"
)

var Pong = command.MessageCommand{
	Name: "pong",
	Alias: map[string]bool{
		"tong": true,
	},
	Handler: func(s *discordgo.Session, m *discordgo.MessageCreate) bool {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
		return true
	},
}
