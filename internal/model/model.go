package model

import (
	"time"
)

type User struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	LastName string
	UserName string
	Email    string
	Password string
	Posts    []Post
}

type Post struct {
	Id        int       `gorm:"primaryKey"`
	Title     string    `validate:"required,max=20,min=5"`
	Content   string    `validate:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt time.Time
	UserID    int
	User      User `gorm:"constraint:OnDelete:CASCADE;"`
}
