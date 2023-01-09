package model

import (
	"time"
)

type User struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Created        time.Time `json:"created"`
	Updated        time.Time `json:"updated"`
	ManufacturerID string    `json:"manufacturer_id"`
	CreatedBy      string    `json:"created_by"`
	Password       string
}

type Manufacturer struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
