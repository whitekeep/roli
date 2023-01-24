package command

import (
	"errors"
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/discord/embed"
	"github.com/Goscord/goscord/goscord/gateway"
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

func (ctx *Context) SendResponse(content interface{}, ephemeral bool) error {

	// Check if this message is ephemeral
	var isEphemeral discord.MessageFlag
	if ephemeral {
		isEphemeral = discord.MessageFlagEphemeral
	}

	// Check the type of the content
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
			// TODO add ephemeral support
		)

	// TODO: Add os.File support

	default:
		return errors.New("massage type not supported")
	}

}
