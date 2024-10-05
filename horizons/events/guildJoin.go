package horizons

import (
	"HorizonsBot/horizons/utils"

	"github.com/bwmarrin/discordgo"
)

func EmbedJoinBuilder(s *discordgo.Session, userName string) *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{
		Title: "Novo Membro!",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: userName,
				Value: "Seja bem vindo!",
				Inline: true,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Horizons",
			IconURL: utils.GetImgLink(0),
		},
	}

	return embed
}

//func SendJoinMessage()