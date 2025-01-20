package infra

import (
	"github.com/SK-Komori/GoDiscordOkanBot/service/model"
	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
)

type Rating interface {
	CreatePlayer(player model.Player) error
	GetPlayer(id string) (*model.Player, error)
}

type rating struct {
	db *xorm.Engine
}

func NewPlayer(e *xorm.Engine) Rating {
	return &rating{
		db: e,
	}
}

func (r *rating) CreatePlayer(player model.Player) error {
	if _, err := r.db.Table("players").Insert(player); err != nil {
		return errors.Wrap(err, "failed to insert player records")
	}
	return nil
}

func (r *rating) GetPlayer(id string) (*model.Player, error) {
	var player model.Player
	has, err := r.db.Table("players").Where("id = ?", id).Get(&player)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get player records")
	}
	if !has {
		return nil, nil
	}
	return &player, nil
}
