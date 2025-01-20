package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {
	TOKEN := os.Getenv("DISCORD_BOT_TOKEN")
	var (
		channelID string
		guildID   string
	)

	flag.StringVar(&guildID, "guild_id", "", "guild id")
	flag.StringVar(&channelID, "channel_id", "", "channel id")

	flag.Parse()

	if guildID == "" {
		fmt.Println("unset paramter: guild id")
		os.Exit(1)
	}

	if channelID == "" {
		fmt.Println("unset parameter: channel id")
		os.Exit(1)
	}

	fmt.Println(guildID)
	fmt.Println(channelID)

	ses, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		fmt.Printf("failed to creating Discord session: %+v\n", err)
		os.Exit(1)
	}

	err = ses.Open()
	if err != nil {
		fmt.Printf("failed to opening connection: %+v\n", err)
		os.Exit(1)
	}
	defer ses.Close()

	g, err := ses.State.Guild(guildID)
	if err != nil {
		fmt.Printf("failed to get guild state: %+v\n", err)
		os.Exit(1)
	}

	fmt.Printf("guild id: %s\n", g.ID)
	fmt.Printf("guild name: %s\n", g.Name)
	for _, vs := range g.VoiceStates {
		fmt.Println("-----")
		fmt.Println(vs.ChannelID)
		fmt.Println(vs.UserID)
		if vs.ChannelID == channelID {
			if err := ses.GuildMemberMove(guildID, vs.UserID, nil); err != nil {
				fmt.Printf("failed to remove member: %+v\n", err)
				os.Exit(1)
			}
		}
	}
}
