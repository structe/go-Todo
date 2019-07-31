package model

import "time"

type BaseModel struct {
	ID        int `json:"id",gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
