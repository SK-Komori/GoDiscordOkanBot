package model

import "time"

type Player struct {
	ID      string
	Rate    int
	Created time.Time
	Updated time.Time
}
