package main

import (
	"log"
	"os/exec"

	"github.com/bwmarrin/discordgo"
)

func main() {
	dg, err := discordgo.New("Bot " + "BOT_TOKEN")
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
	}

	dg.AddHandler(handleMessage)

	err = dg.Open()
	if err != nil {
		log.Fatal("Error opening Discord session: ", err)
	}

	defer dg.Close()

	log.Println("Bot is now running. Press CTRL-C to exit.")
	select {}
}

// func handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
//	if m.Author.Bot {
//		return
//	}
//
//	if m.Content == "!sysinfo" {
//		output, err := exec.Command("sh", "-c", "uname -a").Output()
//		if err != nil {
//			s.ChannelMessageSend(m.ChannelID, "Error executing command: "+err.Error())
//		}
//		s.ChannelMessageSend(m.ChannelID, "```"+string(output)+"```")
//	}

// func handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
//	if m.Author.Bot {
//		return
//	}
//
//	if m.Content == "" {
//		return
//	}
//
//	output, err := exec.Command("sh", "-c", m.Content).Output()
//	if err != nil {
//		s.ChannelMessageSend(m.ChannelID, "Error executing command: "+err.Error())
//	}
//
//	s.ChannelMessageSend(m.ChannelID, "```"+string(output)+"```")
//}
//

func handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if m.Content == "" {
		return
	}

	output, err := exec.Command("bash", "-c", "source ~/.zshrc && "+m.Content).Output()
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error executing command: "+err.Error())
	}

	s.ChannelMessageSend(m.ChannelID, "```"+string(output)+"```")
}
