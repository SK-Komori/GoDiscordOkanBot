package handler

import (
	"fmt"

	"github.com/SK-Komori/GoDiscordOkanBot/service"
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

type RatingHandler interface {
	CreatePlayer(s *discordgo.Session, opt interactOpt) (*discordgo.InteractionResponse, error)
	RatingBattle(s *discordgo.Session, opt interactOpt) (*discordgo.InteractionResponse, error)
}

type ratingHandler struct {
	service service.Rating
}

func NewRatingHandler(r service.Rating) RatingHandler {
	return &ratingHandler{service: r}
}

func (r *ratingHandler) CreatePlayer(s *discordgo.Session, opt interactOpt) (*discordgo.InteractionResponse, error) {
	already, err := r.service.CreatePlayer(opt.From.User.ID)
	if err != nil {
		return &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "登録中にエラっちゃったよ…",
			},
		}, errors.Wrap(err, "failed to create player")
	}

	var msg string
	if already {
		msg = fmt.Sprintf("%s さんは既にレーティングに登録済みでした！", opt.From.Mention())
	} else {
		msg = fmt.Sprintf("%s さんをレーティングに登録しました！", opt.From.Mention())
	}
	res := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	}
	return res, nil
}

func (r *ratingHandler) RatingBattle(s *discordgo.Session, opt interactOpt) (*discordgo.InteractionResponse, error) {
	var msg string
	res := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	}
	return res, nil
}
