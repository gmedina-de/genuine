package user

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func UserRepository() userRepository {
	// todo: use settings for allowing another databases
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&user{})
	return userRepository{db}
}

func (g *userRepository) Create(model interface{}) {
	g.db.Table("users")
	g.db.Create(model)
}

func (g *userRepository) RetrieveAll(model interface{}) {
	g.db.Table("users").Find(model)
}
