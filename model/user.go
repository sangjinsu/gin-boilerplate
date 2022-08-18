package model

import (
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Age          uint8     `json:"age"`
	Birthday     time.Time `json:"birthday"`
	MemberNumber string    `json:"member_number"`
	ActivatedAt  time.Time `json:"activated_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
