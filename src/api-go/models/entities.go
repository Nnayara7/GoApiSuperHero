package models

type SuperType string

const (
	HeroType    SuperType = "good"
	VillainType SuperType = "bad"
)

type Super struct {
	ID           int64  `gorm:"PRIMARY_KEY"`
	UUID         int64  `gorm:"UNIQUE"`
	Name         string
	Intelligence int64
	strength int64
	speed int64
	durability int64
	power int64
	combat int64
}
