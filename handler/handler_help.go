package handler

import "github.com/bwmarrin/discordgo"

func Help(s *discordgo.Session, opt cmdOpt) error {
	s.ChannelMessageSendEmbed(opt.ChannelID, &discordgo.MessageEmbed{
		Title: "Help Window",
		Color: 0x11FFFF,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Item 1",
				Value:  "Vaaaaalue1",
				Inline: true,
			},
			{
				Name:   "Item 2",
				Value:  "Vaaaaalue02",
				Inline: true,
			},
		},
	})
	return nil
}
