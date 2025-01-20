package infra

import "github.com/go-xorm/xorm"

type Coin interface {
	CreateBettor()
}

type coin struct {
	db *xorm.Engine
}

func NewCoin(e *xorm.Engine) Coin {
	return &coin{
		db: e,
	}
}

func (c *coin) CreateBettor() {

}
