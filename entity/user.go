package entity

import (
	"time"
)

type User struct {
	Id        int64     `json:"id"`
	Alias     string    `json:"alias"`
	Name      string    `json:"name"`
	Gender    int       `json:"gender"`
	Birth     time.Time `json:"birth"`
	State     int       `json:"state"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//func (User) TableName() string {
//	return "zl_users"
//}
