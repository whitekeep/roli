package command

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/discord/embed"
	"strings"
)

type EmbedCommand struct{}

func (c *EmbedCommand) Name() string {
	return "embed"
}

func (c *EmbedCommand) Description() string {
	return "Отправить embed с твоим текстом"
}

func (c *EmbedCommand) Category() string {
	return "general"
}

func (c *EmbedCommand) Options() []*discord.ApplicationCommandOption {
	return []*discord.ApplicationCommandOption{
		{
			Name:        "title",
			Type:        discord.ApplicationCommandOptionString,
			Description: "Заголовок",
			Required:    true,
		},
		{
			Name:        "description",
			Type:        discord.ApplicationCommandOptionString,
			Description: "Описание используй -br для переноса строки",
			Required:    true,
		},
	}
}

func (c *EmbedCommand) Execute(ctx *Context) bool {
	e := embed.NewEmbedBuilder()

	title := ctx.interaction.Data.(discord.ApplicationCommandData).Options[0].Value.(string)
	description := ctx.interaction.Data.(discord.ApplicationCommandData).Options[1].Value.(string)

	e.AddField(title, strings.ReplaceAll(description, "-br", "\n"), false)
	e.SetColor(embed.Green)

	_ = ctx.SendResponse(e.Embed(), true)

	return true
}
