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