package command

import (
	"github.com/bwmarrin/discordgo"
)

type MessageCommand struct {
	Name    string
	Alias   map[string]bool
	Handler func(*discordgo.Session, *discordgo.MessageCreate) bool
}

type MessageCommands []MessageCommand
