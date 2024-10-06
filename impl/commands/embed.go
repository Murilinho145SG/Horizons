package commandsHorizons

import (
	"github.com/bwmarrin/discordgo"
)

func EmbedCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name == "embed" {
		member := i.Member

		if member != nil {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Seja bem vindo(a) " + i.Member.Mention(),
				},
			})
		}
	}
}