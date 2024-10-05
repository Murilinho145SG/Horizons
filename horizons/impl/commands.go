package horizons

import "github.com/bwmarrin/discordgo"

func RegisterCommands(s *discordgo.Session) error {
	_, err := s.ApplicationCommandCreate(s.State.User.ID, "", &discordgo.ApplicationCommand{
		Name: "test",
		Description: "retorna o teste",
	})

	if err != nil {
		return err
	}

	return nil
}