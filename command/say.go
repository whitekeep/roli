package command

import (
	"github.com/Goscord/goscord/discord"
	"roli/utils"
)

type SayCommand struct{}

func (c *SayCommand) Name() string {
	return "say"
}

func (c *SayCommand) Description() string {
	return "Say all after the command in the channel!"
}

func (c *SayCommand) Category() string {
	return "general"
}

func (c *SayCommand) Options() []*discord.ApplicationCommandOption {
	return []*discord.ApplicationCommandOption{
		{
			Name:        "msgid",
			Type:        discord.ApplicationCommandOptionString,
			Description: "ID of the message to say",
			Required:    true,
		},
		{
			Name:        "chnid",
			Type:        discord.ApplicationCommandOptionString,
			Description: "ID of the channel where say the message",
			Required:    true,
		},
	}
}

func (c *SayCommand) Execute(ctx *Context) bool {
	// Check if the user have the permission to use this command
	permission, err := utils.HavePermission(ctx.interaction.Member.Roles, ctx.config.Roles, utils.Admin)
	if err != nil || !permission {
		// If the user don't have the permission, send a message to the user
		_, _ = ctx.client.Interaction.CreateResponse(
			ctx.interaction.Id,
			ctx.interaction.Token,
			&discord.InteractionCallbackMessage{
				Content: "You don't have the permission to use this command!",
				Flags:   discord.MessageFlagEphemeral,
			})
		return false
	}

	// Get message and channel by provided ID
	msg, err := ctx.client.Channel.GetMessage(ctx.interaction.ChannelId, ctx.interaction.Data.Options[0].String())
	channel, err := ctx.client.Channel.GetChannel(ctx.interaction.Data.Options[1].String())

	// Send target message in target channel
	_, err = ctx.client.Channel.SendMessage(channel.Id, msg.Content)

	// Send ephemeral message to user
	if err != nil {
		_, _ = ctx.client.Interaction.CreateResponse(
			ctx.interaction.Id,
			ctx.interaction.Token,
			&discord.InteractionCallbackMessage{
				Content: "Error!❌",
				Flags:   discord.MessageFlagEphemeral,
			})
	} else {
		_, _ = ctx.client.Interaction.CreateResponse(
			ctx.interaction.Id,
			ctx.interaction.Token,
			&discord.InteractionCallbackMessage{
				Content: "Done!✅",
				Flags:   discord.MessageFlagEphemeral,
			})
	}

	return true
}
