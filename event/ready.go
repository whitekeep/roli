package event

import (
	"log"

	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/gateway"
	"roli/command"
	"roli/config"
)

func OnReady(client *gateway.Session, config *config.Config, cmdMgr *command.CommandManager) func() {
	return func() {
		log.Printf("Logged in as %s\n", client.Me().Tag())

		// Initialize slash commands
		cmdMgr.Init()

		// Set activity
		var activityType discord.ActivityType

		switch config.ActivityType {
		case "playing":
			activityType = discord.ActivityPlaying
		case "streaming":
			activityType = discord.ActivityStreaming
		case "listening":
			activityType = discord.ActivityListening
		case "watching":
			activityType = discord.ActivityWatching
		}

		_ = client.SetActivity(&discord.Activity{Name: config.ActivityName, Type: activityType})
	}
}
