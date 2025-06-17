package model

import "time"

// Product represents a product in the e-commerce system.
type Product struct {
	ID          string  `json:"id" gorm:"PrimaryKey; not null; unique; index"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Currency    string  `json:"currency"`
	SKU         string  `json:"sku"`
	Stock       int     `json:"stock"`
	// CategoryIDs []string  `json:"category_ids"`
	// ImageURLs   []string  `json:"image_urls"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:current_timestamp"`
	IsAvailable bool      `json:"is_available"`
}
