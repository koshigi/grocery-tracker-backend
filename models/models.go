package models

import (
	"gorm.io/gorm"
)

// Table Names: By default, GORM converts struct names to snake_case and pluralizes them for table names. For instance, a User struct becomes users in the database.

// Column Names: GORM automatically converts struct field names to snake_case for column names in the database.

// gorm.Model definition
type Item struct {
	gorm.Model
	ID          uint    `gorm:"primaryKey"` // Standard field for the primary key
	Title       string  // A regular string field
	Description *string // A pointer to a string, allowing for null values
	Price       uint8   // An unsigned 8-bit integer
}

type List struct {
	gorm.Model
	ID          uint    `gorm:"primaryKey"` // Standard field for the primary key
	Title       string  // A regular string field
	Description *string // A pointer to a string, allowing for null values
}

type ItemsList struct {
	gorm.Model
	ID     uint `gorm:"primaryKey"` // Standard field for the primary key
	ItemId uint
	Item   Item `gorm:"foreignKey:ItemId"`
	ListId uint
	List   List `gorm:"foreignKey:ListId"`
}
