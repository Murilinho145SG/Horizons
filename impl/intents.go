package implHorizons

import "github.com/bwmarrin/discordgo"

func SetIntents() discordgo.Intent {
	return discordgo.IntentsGuildMembers | discordgo.IntentsGuildMessages | discordgo.IntentsGuildMessageReactions
}