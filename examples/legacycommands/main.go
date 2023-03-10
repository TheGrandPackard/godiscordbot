package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"github.com/thegrandpackard/godiscordbot"
)

var (
	commands = map[string]func(d *godiscordbot.DiscordBot, m *discordgo.MessageCreate, arguments []string){
		// Help Command
		"!help": func(d *godiscordbot.DiscordBot, m *discordgo.MessageCreate, arguments []string) {
			log.Info().
				Str("member", m.Author.Username).
				Str("channel", m.ChannelID).
				Str("command", "!help").
				Msg("Legacy Command Received")

			d.SendMessageWithDeletionTimeout(m.ChannelID, "Help command response goes here.", d.LegacyCommandResponseTimeout)
		},
		// Info Command
		"!info": func(d *godiscordbot.DiscordBot, m *discordgo.MessageCreate, arguments []string) {
			log.Info().
				Str("member", m.Author.Username).
				Str("channel", m.ChannelID).
				Str("command", "!info").
				Msg("Legacy Command Received")

			d.SendMessageWithDeletionTimeout(m.ChannelID, "This is an example discord bot.", d.LegacyCommandResponseTimeout)
		},
	}
)

func main() {
	// Configure and initialize the bot
	bot, err := godiscordbot.New(godiscordbot.Options{
		DiscordToken:           "<DISCORD_BOT_TOKEN>",
		Intents:                      discordgo.IntentsGuildMessages,
		LegacyCommandDeletionTimeout: time.Second * 30,
		LegacyCommandResponseTimeout: time.Minute * 2,
		LegacyCommandPrefix:          "!",
	})
	if err != nil {
		panic(err)
	}

	// Map legacy commands (can be added after starting the bot)
	bot.MapLegacyCommands(commands)

	// Start the bot
	err = bot.Start()
	if err != nil {
		panic(err)
	}
	// Stop the bot when the application is exited
	defer bot.Stop()

	// Capture Ctrl-c to stop the bot
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Info().Msg("Gracefully shutting down")
}
