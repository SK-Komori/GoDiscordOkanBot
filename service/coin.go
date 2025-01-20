package service

import "github.com/SK-Komori/GoDiscordOkanBot/infra"

type Coin interface {
	CreateBettor()
}

type coin struct {
	infra infra.Coin
}

func NewServiceCoin(i infra.Coin) Coin {
	return &coin{
		infra: i,
	}
}

func (c *coin) CreateBettor() {

}
