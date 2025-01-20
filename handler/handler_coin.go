package handler

import (
	"github.com/SK-Komori/GoDiscordOkanBot/service"
	"github.com/bwmarrin/discordgo"
)

type CoinHandler interface {
	CreateBettor(s *discordgo.Session, opt interactOpt) (*discordgo.InteractionResponse, error)
}

type coinHandler struct {
	service service.Coin
}

func NewCoinHandler(c service.Coin) CoinHandler {
	return &coinHandler{
		service: c,
	}
}

func (c *coinHandler) CreateBettor(s *discordgo.Session, opt interactOpt) (*discordgo.InteractionResponse, error) {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "賭け人登録の方はまだ未実装ﾆｬ",
		},
	}, nil
}
