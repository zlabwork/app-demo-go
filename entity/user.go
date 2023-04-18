package entity

import (
	"time"
)

type User struct {
	Id        int64
	Alias     string
	Name      string
	Gender    int
	Birth     time.Time
	State     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
