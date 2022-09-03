package models

import "time"

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Password string `json:"-" gorm:"type: varchar(255)"`
	Status   string `json:"status" gorm:"type: varchar(255)"`
	PostCode string `json:"post_code"`
	Address  string `json:"address" gorm:"type: varchar(255)"`
	// Profile   ProfileResponse `json:"profile"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UsersProfileResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Status   string `json:"status"`
	PostCode string `json:"post_code"`
	Address  string `json:"address"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
