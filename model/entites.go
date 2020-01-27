package model

import (
	"time"
)

type Role struct {
	ID uint
	Name string
	Users []User
}


type  User struct {
	ID uint
	userName string `gorm:"type:varchar(255);not null"`
	phoneNo string `gorm:"type:varchar(255);not null"`
	password string `gorm:"type: varchar(255)"`
	email string `gorm:"type:varchar(255);not null; unique"`
	tin_no string `gorm:"type:varchar(255);not null; unique"`
	Bids []Bid


}
type Session struct {
	ID uint
	UUID string `gorm:"type:varchar(255); not null"`
	Expires int64 `gorm:"type:varchar(255); not null"`
	SigninKey []byte `gorm:"type:varchar(255); not null"`
}
type Stock struct {
	ID uint
	name string `gorm:"type:varchar(255);not null"`
	desc string  `gorm:"type:varchar(255)`
	turnoff bool
	min_value float64
	max_value float64
	start_time time.Time
	end_time time.Time
	CategoryID uint

}
type Category struct {
	ID uint
	name string
	Stocks []Stock
}
type Bid struct {
	ID uint
	name string
	price string
	status bool
	StockID uint
	UserID uint
	sent_time time.Time
}