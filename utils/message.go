package utils

import (
	"fmt"
	"github.com/Goscord/goscord/goscord/discord"
)

func GetMessageUrl(msg *discord.Message) string {
	return fmt.Sprintf("https://discord.com/channels/%s/%s/%s", msg.GuildId, msg.ChannelId, msg.Id)
}
