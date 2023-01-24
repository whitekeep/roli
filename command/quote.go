package command

import (
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/discord/embed"
)

type QuoteCommand struct{}

func (c *QuoteCommand) Name() string {
	return "цитата"
}

func (c *QuoteCommand) Description() string {
	return "Процитировать сообщение"
}

func (c *QuoteCommand) Category() string {
	return "general"
}

func (c *QuoteCommand) Options() []*discord.ApplicationCommandOption {
	return []*discord.ApplicationCommandOption{
		{
			Name:        "message_id",
			Type:        discord.ApplicationCommandOptionString,
			Description: "ID сообщения",
			Required:    true,
		},
		{
			Name:        "channel_id",
			Type:        discord.ApplicationCommandOptionString,
			Description: "Куда отправить",
			Required:    false,
		},
	}
}

func (c *QuoteCommand) Execute(ctx *Context) bool {
	// Get target message
	targetMassage, err := ctx.client.Channel.GetMessage(ctx.interaction.ChannelId, ctx.interaction.Data.Options[0].String())
	if err != nil {
		_, _ = ctx.SendResponse("Не удалось найти сообщение!", true)
		return true
	}

	// Check if target message is not empty and not embed
	if targetMassage.Content == "" && targetMassage.Embeds == nil {
		_, _ = ctx.SendResponse("Сообщение пустое!", true)
		return true
	}

	// Build embed
	e := embed.NewEmbedBuilder()

	e.AddField(targetMassage.Content, "_", false)
	e.SetTimestamp(targetMassage.Timestamp)
	e.SetColor(embed.Purple)
	e.SetFooter(targetMassage.Author.Username+"#"+targetMassage.Author.Discriminator, targetMassage.Author.AvatarURL())

	// Add all attachments
	for _, attachment := range targetMassage.Attachments {
		// TODO Set image not work!
		e.SetImage(attachment.URL)
	}

	// Set target channel
	targetChannel := ctx.interaction.ChannelId
	if len(ctx.interaction.Data.Options) >= 2 {
		chn, err := ctx.client.Channel.GetChannel(ctx.interaction.Data.Options[1].String())
		targetChannel = chn.Id
		if err != nil {
			_, _ = ctx.SendResponse("Не удалось найти канал!", true)
		}
	}

	// Send embed
	_, err = ctx.client.Channel.SendMessage(targetChannel, e.Embed())
	if err != nil {
		_, _ = ctx.SendResponse("Не удалось отправить сообщение!", true)
	}

	// Send Done response
	_, _ = ctx.SendResponse("Сообщение отправлено!", false)

	return true
}
