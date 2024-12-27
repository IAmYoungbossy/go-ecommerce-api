package models

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User represents the user model in the application.
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Role      string    `json:"role" gorm:"default:user"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// BeforeSave is a GORM hook to hash the password before saving.
func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("Error hashing password:", err)
			return err
		}
		user.Password = string(hashedPassword)
	}
	return nil
}

// CheckPassword compares a plain text password with the hashed password in the database.
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
