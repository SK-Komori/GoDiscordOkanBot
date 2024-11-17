package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

func main() {
	TOKEN := os.Getenv("DISCORD_BOT_TOKEN")
	GUILD_ID := os.Getenv("GUILD_ID")

	ses, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}

	// discord側に作るコマンド
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "bot-test",
			Description: "TEST Command",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "mojiretsu",
					Description: "試しになんか入力してくれ",
				},
			},
		},
		{
			Name:        "bubuduke",
			Description: "ぶぶ漬けでもいかがどす？",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionChannel,
					Name:        "channel",
					Description: "追い出したいボイスチャンネル",
				},
			},
		},
	}

	// コード側で動く処理内容
	commandHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"bot-test": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			var respStr string
			options := i.ApplicationCommandData().Options

			userNick := i.Member.DisplayName()

			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			if option, ok := optionMap["mojiretsu"]; ok {
				respStr += option.StringValue()
			} else {
				respStr += "何もないよ；；"
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: userNick + "さん こんにちは！！\n あなたが入力したのは「" + respStr + "」ですね！",
				},
			})
		},
		"bubuduke": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			respStr := ""
			options := i.ApplicationCommandData().Options

			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			g, err := s.State.Guild(i.GuildID)
			if err != nil {
				log.Panic(err)
			}

			if option, ok := optionMap["channel"]; ok {
				vc := option.ChannelValue(s)
				if vc.Type == discordgo.ChannelTypeGuildVoice {
					respStr = fmt.Sprintf("%sに参加している人を移動させます！", vc.Name)
					for _, vs := range g.VoiceStates {
						if vs.ChannelID == vc.ID {
							// respStr += fmt.Sprintf("\n%s", vs.UserID)
							if err := s.GuildMemberMove(i.GuildID, vs.UserID, nil); err != nil {
								log.Panic(err)
							}
						}
					}
				} else {
					respStr = "ボイスチャンネルの選択してね"
				}
			}
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: respStr,
				},
			})
		},
	}

	ses.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	ses.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err = ses.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := ses.ApplicationCommandCreate(ses.State.User.ID, GUILD_ID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer ses.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	isRemoveCommands := true
	if isRemoveCommands {
		log.Println("Removing commands...")
		for _, v := range registeredCommands {
			err := ses.ApplicationCommandDelete(ses.State.User.ID, GUILD_ID, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}

	log.Println("Gracefully shutting down.")
}
