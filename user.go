package entity

import (
	"time"
	"gorm.io/gorm"
)

type Users struct{
	gorm.Model
	Username	string
	FirstName	string
	LastName	string
	Email		string
	Password	string
	BirthDay	time.Time 
	Profilepic	string	`gorm:"type:longtext"`
}
