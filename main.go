package main

import (
	"github.com/Goscord/goscord/goscord"
	"github.com/Goscord/goscord/goscord/gateway"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"roli/command"
	"roli/config"
	"roli/event"
)

var (
	client *gateway.Session
	Config *config.Config
	cmdMgr *command.CommandManager
)

func main() {
	// Load environment variables
	_ = godotenv.Load()

	Config, _ = config.GetConfig()

	// Create client instance
	client = goscord.New(&gateway.Options{
		Token: os.Getenv("BOT_TOKEN"),
		Intents: gateway.IntentsGuild |
			gateway.IntentGuildMessages |
			gateway.IntentGuildMembers |
			gateway.IntentDirectMessages |
			gateway.IntentMessageContent,
	})

	// Load command manager
	cmdMgr = command.NewCommandManager(client, Config)

	// Load events
	_ = client.On("ready", event.OnReady(client, Config, cmdMgr))
	_ = client.On("interactionCreate", cmdMgr.Handler(client, Config))

	// Login client
	if err := client.Login(); err != nil {
		panic(err)
	}

	// Wait here until term signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session
	client.Close()
}
