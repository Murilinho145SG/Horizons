package horizons

import (
	horizons "HorizonsBot/horizons/events"

	"github.com/bwmarrin/discordgo"
)


func EmbedCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name == "embed" {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					horizons.EmbedJoinBuilder(s, "Murilinho145"),
				},
			},
		})
	}
}