package constants

import (
	"github.com/Meta502/go-discord-boilerplate/src/command"
	"github.com/Meta502/go-discord-boilerplate/src/pingpong"
)

var Commands command.MessageCommands = []command.MessageCommand{
	pingpong.Ping,
	pingpong.Pong,
}
