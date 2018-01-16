package model

import (
	"time"
)

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}
type Receipt struct {
	Model
	Email         string  `gorm:"not null" json:"email"`
	FirstName     string  `gorm:"not null" json:"first_name"`
	LastName      string  `gorm:"not null" json:"last_name"`
	FirstNameRead string  `gorm:"not null" json:"first_name_read"`
	LastNameRead  string  `gorm:"not null" json:"last_name_read"`
	Phone         string  `gorm:"not null" json:"phone"`
	Address       string  `gorm:"not null" json:"address"`
	Sales         []Sales `gorm:"-" json:"sales"`
	TotalPrice    uint    `gorm:"not null" json:"total_price"`
}
type Sales struct {
	Model
	ReceiptID uint `gorm:"not null" json:"receipt_id"`
	ProductID uint `gorm:"not null" json:"product_id"`
	Price     uint `gorm:"not null" json:"price"`
	Count     uint `gorm:"not null" json:"count"`
}

type Seller struct {
	Model
	Name        string `gorm:"not null" json:"name"`
	InstagramID string `gorm:"not null" json:"instagram_id"`
}
type Product struct {
	Model
	Name      string          `gorm:"not null" json:"name"`
	Details   string          `gorm:"not null" json:"details"`
	Price     string          `gorm:"not null" json:"price"`
	SellerID  uint            `gorm:"not null" json:"seller_id"`
	Release   bool            `gorm:"not null" json:"release"`
	remaining uint            `gorm:"not null" json:"remaining"`
	Images    []ProductImages `gorm:"-" json:"images"`
}
type ProductImages struct {
	ProductID uint   `gorm:"not null" json:"product_id"`
	DisplayNo uint   `gorm:"not null" json:"display_no"`
	ImagePath string `gorm:"not null" json:"image_path"`
}
