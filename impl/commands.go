package implHorizons

import "github.com/bwmarrin/discordgo"

func updateCommands(s *discordgo.Session) error {
	commands, _ := s.ApplicationCommands(s.State.User.ID, "")
	for _, command := range commands {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", command.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

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
		{
			Name: "antipessoas",
			Description: "basicamente evita a pessoa mencionada",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type: discordgo.ApplicationCommandOptionUser,
					Name: "user",
					Description: "Usuário para ser abusado",
					Required: true,
				},
			},
		},
	}

	err := updateCommands(s)
	
	if err != nil {
		return err
	}

	for _, command := range commands {

		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", command)

		if err != nil {
			return err
		}
	}

	return nil
}