package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       		uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt `gorm:"index"`
	Name    		string   	   `gorm:"not null"`
	Email    		string   	   `gorm:"not null;unique"`
	Password		string    	   `gorm:"not null"`
	Role     		string         `gorm:"not null"`
	
	Account  		[]Account 	   `gorm:"foreignKey:UserID"`
	// Payments 		[]Payment      `gorm:"foreignKey:UserID"`
	// Transaction 	[]Account      `gorm:"foreignKey:UserID"`
}

type Account struct {
	ID        		uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt `gorm:"index"`
	UserID          uuid.UUID      `gorm:"type:uuid;not null"`
	User            User   		   `gorm:"foreignKey:UserID"`
	AccountNumber   string   	   `gorm:"not null;unique"`
	Balance         float64 	   `gorm:"not null"`
	CreatedBy    	uuid.UUID      `gorm:"type:uuid;not null"`
	CreatedByUser   User           `gorm:"foreignKey:CreatedBy"`
	
}