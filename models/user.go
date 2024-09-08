package models

import (
	"time"
)

// type User struct {
// 	ID       int    `json:"id" gorm:"primaryKey"`
// 	Username string `json:"user_name" gorm:"unique;not null"`
// 	Password string `json:"password" gorm:"not null"`
// 	Role     string `json:"role" gorm:"not null"` // 'admin', 'waiter'
// }

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Username  string `json:"username" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	RoleID    uint   `json:"role_id" gorm:"not null"`
	Role      Role   `json:"role" gorm:"foreignKey:RoleID;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
