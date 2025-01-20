package config

import (
	"os"
)

var (
	Bot bot
	DB  db
)

type bot struct {
	Token string
}

type db struct {
	UserName string
	Password string
	Host     string
	DataBase string
	Port     string
}

func ReadConfig() error {

	// string系の環境変数はfor文でやっちゃう
	envKeys := []string{"DISCORD_BOT_TOKEN", "DB_USERNAME", "DB_PASSWORD", "DB_HOST", "DB_DATABASE", "DB_PORT"}
	env := map[string]string{}
	for i := range envKeys {
		env[envKeys[i]] = os.Getenv(envKeys[i])
	}

	Bot = bot{
		Token: env["DISCORD_BOT_TOKEN"],
	}

	DB = db{
		UserName: env["DB_USERNAME"],
		Password: env["DB_PASSWORD"],
		Host:     env["DB_HOST"],
		DataBase: env["DB_DATABASE"],
		Port:     env["DB_PORT"],
	}

	return nil
}
