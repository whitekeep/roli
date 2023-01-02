package command

import (
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/gateway"
	"roli/config"
)

type Context struct {
	config      *config.Config
	cmdMgr      *CommandManager
	client      *gateway.Session
	interaction *discord.Interaction
}

type Command interface {
	Name() string
	Description() string
	Category() string
	Options() []*discord.ApplicationCommandOption
	Execute(ctx *Context) bool
}
