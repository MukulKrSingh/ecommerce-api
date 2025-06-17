package model

import "time"

type User struct {
	ID               string    `json:"id" gorm:"PrimaryKey"`
	Email            string    `json:"email" gorm:"index;unique;not null"`
	Password         string    `json:"password"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	Phone            string    `json:"phone"`
	Address          string    `json:"address"`
	ProfilePicture   string    `json:"profile_picture"`
	IsSeller         bool      `json:"is_seller"`
	VerificationCode string    `json:"verification_code"`
	Verified         bool      `json:"verified" gorm:"default:false"`
	Role             string    `json:"role" gorm:"default:user"` // e.g., "user", "admin", "seller"
	CreatedAt        time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"default:current_timestamp"`

	// Cart             []CartItem `json:"cart"`
	// Orders           []Order    `json:"orders"`
}
