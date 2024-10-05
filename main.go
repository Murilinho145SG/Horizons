package main

import (
	"HorizonsBot/impl"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	byte, berr := os.ReadFile("token")

	if berr != nil {
		fmt.Errorf("Erro ao ler o token: ", berr)
		return
	}

	token := string(byte)
	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Errorf("Erro ao criar sessão do bot: ", err)
		return
	}

	EventListeners(dg)

	err = dg.Open()

	if err != nil {
		fmt.Errorf("Erro ao conectar no discord: ", err)
		return
	}

	horizons.RegisterCommands(dg)

	fmt.Println("Bot Online com sucesso")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}

func EventListeners(s *discordgo.Session) {
	s.AddHandler(horizons.Command)
	s.AddHandler(horizons.EmbedCommand)
}
