package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	UserID   string `gorm:"unique" json:"userID"`
	Email   string  `gorm:"unique" json:"email"`
	Nickname   string `json:"nickname"`
	Password string `json:"password"`
	Active 	bool	`json:"active"`
}

func (e *User) Disable() {
	e.Active = false
}

func (p *User) Enable() {
	p.Active = true
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrateUser(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	return db
}