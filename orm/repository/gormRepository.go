package repository

import (
	model2 "github.com/geremde/genuine/orm/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type gormRepository struct {
	db *gorm.DB
}

func GormRepository() Repository {
	// todo: use settings for allowing another databases
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return &gormRepository{db}
}

func (g *gormRepository) Migrate(model interface{}) {
	g.db.AutoMigrate(model)
}

func (g *gormRepository) Create(model interface{}) {
	g.db.Table("users")
	g.db.Create(model)
}

func (g *gormRepository) RetrieveAll(model interface{}) {
	g.db.Model(&model)
	g.db.Find(&model)
}

func (g *gormRepository) RetrieveFirst(model interface{}, conditions ...interface{}) {
	g.db.First(model,conditions)
}

func (g *gormRepository) Update(model interface{}, column string, value interface{}) {
	g.db.Model(&model).Update(column, value)
}

func (g *gormRepository) Delete(model interface{}) {
	g.db.Delete(&model, model.(model2.Model).ID)
}
