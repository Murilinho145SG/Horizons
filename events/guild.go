package eventsHorizons

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func MemberJoining(s *discordgo.Session, guildMemberAdd *discordgo.GuildMemberAdd) {
	member := guildMemberAdd.Member
	fmt.Println("enviando mensagem de bem vindo")
	if member != nil && !member.User.Bot {
		_, err := s.ChannelMessageSend("1147582211374186566", "Seja bem vindo(a) "+member.Mention())
		if err != nil {
			log.Fatalf("Erro ao enviar mensagem: %v", err)
		}
	}
}

func MemberLeaving(s *discordgo.Session, guildMemberRemove *discordgo.GuildMemberRemove) {
	member := guildMemberRemove.Member
	fmt.Println("enviando mensagem de despedida")
	if member != nil && !member.User.Bot {

		userId := member.User.ID
		var msg string

		if userId == "1154538983355842641" {
			msg = member.Mention() + " Você é um merda"
		} else {
			msg = member.Mention() + " foi embora"
		}
		
		_, err := s.ChannelMessageSend("1147582627075866844", msg)

		if err != nil {
			log.Fatalf("Erro ao enviar mensagem: %v", err)
		}
	}
}