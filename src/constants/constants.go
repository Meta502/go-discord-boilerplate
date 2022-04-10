package constants

import (
	"github.com/Meta502/go-discord-pingpong/src/command"
	"github.com/Meta502/go-discord-pingpong/src/pingpong"
)

var Commands command.MessageCommands = []command.MessageCommand{
	pingpong.Ping,
	pingpong.Pong,
}
