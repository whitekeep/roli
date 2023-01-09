package command

import (
	"errors"
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/discord/embed"
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

func (ctx *Context) SendResponse(content interface{}, ephemeral bool) (*discord.InteractionResponse, error) {

	var isEphemeral discord.MessageFlag
	if ephemeral {
		isEphemeral = discord.MessageFlagEphemeral
	}

	switch currentContent := content.(type) {
	case string:
		return ctx.client.Interaction.CreateResponse(
			ctx.interaction.Id,
			ctx.interaction.Token,
			&discord.InteractionCallbackMessage{
				Content: currentContent,
				Flags:   isEphemeral,
			})

	case *embed.Embed:
		return ctx.client.Interaction.CreateResponse(
			ctx.interaction.Id,
			ctx.interaction.Token,
			currentContent,
		)

	// TODO: Add os.File support

	default:
		return nil, errors.New("massage type not supported")
	}

}
