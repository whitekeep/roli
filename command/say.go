package command

import (
	"github.com/Goscord/goscord/goscord/discord"
	"roli/utils"
)

type SayCommand struct{}

func (c *SayCommand) Name() string {
	return "say"
}

func (c *SayCommand) Description() string {
	return "Повторить сообщение в определённом канале"
}

func (c *SayCommand) Category() string {
	return "general"
}

func (c *SayCommand) Options() []*discord.ApplicationCommandOption {
	return []*discord.ApplicationCommandOption{
		{
			Name:        "message_id",
			Type:        discord.ApplicationCommandOptionString,
			Description: "ID сообщения, которое нужно отправить",
			Required:    true,
		},
		{
			Name:        "channel_id",
			Type:        discord.ApplicationCommandOptionString,
			Description: "ID канала, в который нужно отправить сообщение",
			Required:    true,
		},
	}
}

func (c *SayCommand) Execute(ctx *Context) bool {
	// Check if the user have the permission to use this command
	permission := utils.HavePermission(ctx.interaction.Member.Roles, ctx.config.Roles, utils.Admin)
	if !permission {
		// If the user don't have the permission, send a message to the user
		_ = ctx.SendResponse("У вас недостаточно прав для использования этой команды!", true)
		return false
	}

	// Get message and channel by provided ID
	msg, err := ctx.client.Channel.GetMessage(ctx.interaction.ChannelId, ctx.interaction.Data.(discord.ApplicationCommandData).Options[0].Value.(string))
	channel, err := ctx.client.Channel.GetChannel(ctx.interaction.Data.(discord.ApplicationCommandData).Options[1].Value.(string))

	// Send target message in target channel
	_, err = ctx.client.Channel.SendMessage(channel.Id, msg.Content)

	// TODO Улучшить отработку ошибок
	// Send all attachments
	for _, attachment := range msg.Attachments {
		_, err = ctx.client.Channel.SendMessage(ctx.interaction.ChannelId, attachment.URL)
	}

	// Send ephemeral message to user
	if err != nil {
		_ = ctx.SendResponse("Произошла ошибка! ❌", true)
	} else {
		_ = ctx.SendResponse("Сообщение успешно отправлено! ✅", true)
	}

	return true
}
