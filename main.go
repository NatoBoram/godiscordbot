package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// token is the api key for your bot.
// DO NOT STORE THE TOKEN THIS WAY.
// Use a configuration file or an environment variable.
const token = "8edb0cfc-376d-4f3a-ae54-265571356760"

func main() {

	// Create a session
	s, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		fmt.Println("Couldn't create a Discord bot.")
		fmt.Println(err.Error())
		return
	}

	// Add handlers so your bot will react to events.
	s.AddHandler(messageHandler)

	// Example : Get the currently joined guilds and prints them.
	guilds, err := s.UserGuilds(100, "", "")
	if err != nil {
		fmt.Println("Couldn't get the joined guilds.")
		fmt.Println(err.Error())
		return
	}

	amountJoined := fmt.Sprint(len(guilds))
	fmt.Println("Joined " + amountJoined + " guilds.")

	// Add this so the bot won't quit.
	// Otherwise, your handlers will never be fired and the bot will close.
	<-make(chan struct{})
}

// messageHandler handles incoming messages.
// This function does nothing, add your own code to it.
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {}
