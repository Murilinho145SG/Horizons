package horizons

import "github.com/bwmarrin/discordgo"

func RegisterCommands(s *discordgo.Session) error {
	commands := []*discordgo.ApplicationCommand{
		{
			Name: "test",
			Description: "retorna test",
		},
		{
			Name: "embed",
			Description: "testa embed",
		},
	}

	for _, command := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", command)
		if err != nil {
			return err
		}
	}



	return nil
}