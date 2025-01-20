package service

import (
	"github.com/SK-Komori/GoDiscordOkanBot/infra"
	"github.com/SK-Komori/GoDiscordOkanBot/service/model"
	"github.com/pkg/errors"
)

type Rating interface {
	CreatePlayer(id string) (bool, error)
}

type rating struct {
	infra infra.Rating
}

func NewServiceRating(i infra.Rating) Rating {
	return &rating{
		infra: i,
	}
}

func (r *rating) CreatePlayer(id string) (bool, error) {
	pl, err := r.infra.GetPlayer(id)
	if err != nil {
		return false, errors.Wrapf(err, "failed to get player info, id = %s", id)
	}
	if pl != nil {
		return true, nil
	}

	now := timeNow()
	registData := model.Player{
		ID:      id,
		Rate:    1000,
		Created: now,
		Updated: now,
	}

	if err := r.infra.CreatePlayer(registData); err != nil {
		return false, errors.Wrapf(err, "failed to create player, id = %s", id)
	}
	return false, nil
}
