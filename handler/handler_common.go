package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

const (
	registryRatingID = "registry rating"
	registryCoinID   = "registry coin"
)

func RegistryMenu(s *discordgo.Session, opt cmdOpt) error {
	registryRatingBtn := discordgo.Button{
		Label:    "レーティング登録",
		Style:    discordgo.PrimaryButton,
		CustomID: registryRatingID,
	}

	registryCoinBtn := discordgo.Button{
		Label:    "コイン賭け人登録",
		Style:    discordgo.SuccessButton,
		CustomID: registryCoinID,
	}

	actions := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{registryRatingBtn, registryCoinBtn},
	}

	msg := &discordgo.MessageEmbed{
		Title:       "登録メニュー",
		Description: "レーティングバトルの登録 か コイン使用者の登録 ができるよ",
	}

	data := &discordgo.MessageSend{
		Components: []discordgo.MessageComponent{actions},
		Embed:      msg,
	}

	if _, err := s.ChannelMessageSendComplex(opt.ChannelID, data); err != nil {
		return errors.Wrapf(err, "failed to send registry menu message with channelID: %s", opt.ChannelID)
	}
	return nil
}
