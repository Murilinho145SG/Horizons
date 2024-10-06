package implHorizons

import "github.com/bwmarrin/discordgo"

func SetIntents() discordgo.Intent {
	return discordgo.IntentsGuilds | discordgo.IntentsGuildMembers |discordgo.IntentsGuildVoiceStates |discordgo.IntentsGuildMessages |discordgo.IntentsGuildMessageReactions
}