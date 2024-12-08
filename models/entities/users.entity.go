package entities

import (
	"gorm.io/gorm"

	"time"

	"github.com/google/uuid"
)

type User struct {
    IdUser    uuid.UUID     `gorm:"primaryKey;type:uuid" json:"id_user"`
    FirstName string        `json:"first_name"`
    LastName  string        `json:"last_name"`
    Email     string        `json:"email"`
    Phone     string        `json:"phone"`
    Username  string        `json:"username" gorm:"unique"`
    Password  string        `json:"password"`
    Role      string        `json:"role"`
    CreatedAt time.Time     `json:"created_at"`
    UpdatedAt time.Time     `json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

