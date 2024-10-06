package commandsHorizons

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

var stopAntiPessoas chan bool

func runAntiPessoas(s *discordgo.Session, user *discordgo.User, stopAntiPessoas chan bool) {
	for {
		select {
		case <-stopAntiPessoas:
			return
		default:
			guild, err := s.State.Guild("1147572092708073492")
			if err != nil {
				return
			}

			for _, vs := range guild.VoiceStates {
				if vs.UserID == user.ID {
					member, memberErr := s.GuildMember(guild.ID, vs.UserID)
					if memberErr != nil {
						channel, _ := s.UserChannelCreate("816098460121301062")
						s.ChannelMessageSend(channel.ID, "Erro ao encontrar o membro")
						return
					}

					if !member.Mute {
						s.GuildMemberMute(guild.ID, member.User.ID, true)
						s.GuildMemberDeafen(guild.ID, member.User.ID, true)
					}
					break
				}
			}

			time.Sleep(1 * time.Second)
		}
	}
}

func AntiPessoas(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name == "antipessoas" {
		option := i.ApplicationCommandData().Options[0]
		user := option.UserValue(s)

		if stopAntiPessoas != nil {
			stopAntiPessoas <- true
			stopAntiPessoas = nil

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   discordgo.MessageFlagsEphemeral,
					Content: "O " + user.Mention() + " vai parar de ser abusado",
				},
			})

		} else {
			stopAntiPessoas = make(chan bool)
			go runAntiPessoas(s, user, stopAntiPessoas)

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   discordgo.MessageFlagsEphemeral,
					Content: "O " + user.Mention() + " será abusado",
				},
			})
		}
	}
}
