package service

import (
	"time"

	"github.com/SK-Komori/GoDiscordOkanBot/infra"
	"github.com/go-xorm/xorm"
)

type Service interface {
	Rating
	Coin
}

type service struct {
	Rating
	Coin
}

func NewService(e *xorm.Engine) Service {
	return &service{
		Rating: NewServiceRating(infra.NewPlayer(e)),
		Coin:   NewServiceCoin(infra.NewCoin(e)),
	}
}

func timeNow() time.Time {
	return time.Now()
}
