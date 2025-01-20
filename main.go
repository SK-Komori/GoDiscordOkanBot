package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/SK-Komori/GoDiscordOkanBot/config"
	"github.com/SK-Komori/GoDiscordOkanBot/handler"
	"github.com/SK-Komori/GoDiscordOkanBot/service"
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

func initDB() *xorm.Engine {
	// 接続するDBの情報を文字列で生成
	db := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Asia%%2FTokyo", config.DB.UserName, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.DataBase)

	// ↑の情報をmysqlとして扱いDBのエンジンを生成。これで、engineを通してDBの操作が可能になる。
	engine, err := xorm.NewEngine("mysql", db)
	if err != nil {
		log.Fatalf("failed to connect db: %s", err)
	}

	engine.SetMapper(core.GonicMapper{})
	return engine
}

func main() {

	if err := config.ReadConfig(); err != nil {
		log.Fatal("failed to read config")
	}

	ses, err := discordgo.New("Bot " + config.Bot.Token)
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}

	ses.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	e := initDB()
	svc := service.NewService(e)
	h := handler.NewHandler(svc)

	handler.NewCommandHandler()
	handler.NewInteractionHandler(h)

	ses.AddHandler(handler.CommandHandler)
	ses.AddHandler(handler.InteractionHandler)

	err = ses.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	if err := ses.UpdateListeningStatus("necq help"); err != nil {
		fmt.Println("failed to update listening status, ", err)
		return
	}

	defer func() {
		if err := ses.Close(); err != nil {
			fmt.Printf("falied to session close: %+v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	log.Println("Gracefully shutting down.")
}
