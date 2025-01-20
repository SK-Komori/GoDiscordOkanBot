package handler

import (
	"fmt"
	"log"
	"strings"

	"github.com/SK-Komori/GoDiscordOkanBot/service"
	"github.com/bwmarrin/discordgo"
)

type Handler interface {
	RatingHandler
	CoinHandler
}

type handler struct {
	RatingHandler
	CoinHandler
}

func NewHandler(s service.Service) Handler {
	return &handler{
		RatingHandler: NewRatingHandler(s),
		CoinHandler:   NewCoinHandler(s),
	}
}

type cmdOpt struct {
	ChannelID string
}

type interactOpt struct {
	From *discordgo.Member
}

var (
	commands  map[string]func(s *discordgo.Session, opt cmdOpt) error                                        = map[string]func(s *discordgo.Session, opt cmdOpt) error{}
	interacts map[string]func(s *discordgo.Session, opt interactOpt) (*discordgo.InteractionResponse, error) = map[string]func(s *discordgo.Session, opt interactOpt) (*discordgo.InteractionResponse, error){}
)

func NewCommandHandler() {
	commands["help"] = Help
	commands["touroku"] = RegistryMenu
}

func NewInteractionHandler(h Handler) {
	interacts[registryRatingID] = h.CreatePlayer
	interacts[registryCoinID] = h.CreateBettor
}

func CommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmds := strings.Split(m.Content, " ")
	if cmds[0] == "necq" {
		fmt.Println("necq ok")
		fmt.Println(cmds)
		if h, ok := commands[cmds[1]]; ok {
			fmt.Println("ok")
			h(s, cmdOpt{
				ChannelID: m.ChannelID,
			})
		}
	}
}

func InteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	id := i.MessageComponentData().CustomID
	if h, ok := interacts[id]; ok {
		res, err := h(s, interactOpt{
			From: i.Member,
		})
		if err != nil {
			fmt.Printf("failed to interaction handler: %s\n", err.Error())
		}
		if err := s.InteractionRespond(i.Interaction, res); err != nil {
			log.Panicf("failed to response interaction: %s\n", err.Error())
		}
	}
}
