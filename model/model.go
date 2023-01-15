package model

import (
	"time"
)

type User struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
	Password string
}
