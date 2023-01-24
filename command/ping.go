package command

import (
	"fmt"
	"github.com/Goscord/goscord/goscord/discord"
)

type PingCommand struct{}

func (c *PingCommand) Name() string {
	return "ping"
}

func (c *PingCommand) Description() string {
	return "Получить задержку бота"
}

func (c *PingCommand) Category() string {
	return "general"
}

func (c *PingCommand) Options() []*discord.ApplicationCommandOption {
	return make([]*discord.ApplicationCommandOption, 0)
}

func (c *PingCommand) Execute(ctx *Context) bool {
	_ = ctx.SendResponse(fmt.Sprintf("Pong! 🏓 (%dms)", ctx.client.Latency().Milliseconds()), true)

	return true
}
