package model

import (
	"time"
)

type Role struct {
	ID uint
	Name string
	Users []User `gorm:"many2many:role_user"`
}


type  User struct {
	ID uint
	UserName string `gorm:"type:varchar(255);not null"`
	PhoneNo string `gorm:"type:varchar(255);not null"`
	Password string `gorm:"type: varchar(255)"`
	Email string `gorm:"type:varchar(255);not null; unique"`
	TinNo string `gorm:"type:varchar(255);not null; unique"`
	Bid []Bid `gorm:"many2many:users_bid"`
//
}


type Session struct {
	ID uint
	UUID string `gorm:"type:varchar(255); not null"`
	Expires int64 `gorm:"type:varchar(255); not null"`
	SigninKey []byte `gorm:"type:varchar(255); not null"`
}


type Stock struct {
	ID uint
	Name string `gorm:"type:varchar(255);not null"`
	Desc string  `gorm:"type:varchar(255)`
	Turnoff bool
	MinValue float64
	MaxValue float64
	StartTime time.Time
}


type Category struct {
	ID uint
	Name string
	Stock []Stock  `gorm:"many2many:stock_category"`
}


type Bid struct {
	ID uint
	Name string
	Price string
	Status bool
	StockID uint
	SentTime time.Time
}
