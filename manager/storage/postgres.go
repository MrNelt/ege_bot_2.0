package storage

import (
	"time"

	log "github.com/bearatol/lg"

	"github.com/kappaprideonly/ege_bot_2.0/manager/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgStorage struct {
	db *gorm.DB
}

func NewPgStorage(dsl string) *PgStorage {
	log.Debug("[postgres] %s", dsl)
	db, err := gorm.Open(postgres.Open(dsl), &gorm.Config{})
	for err != nil {
		log.Info("Trying to connect to pg")
		time.Sleep(2 * time.Second)
		db, err = gorm.Open(postgres.Open(dsl), &gorm.Config{})
	}
	log.Info("Connect to pg")
	db.AutoMigrate(&model.User{})
	return &PgStorage{db}
}

func (p *PgStorage) existUser(id uint) bool {
	user := model.User{}
	result := p.db.First(&user, id)
	return result.Error == nil
}

func (p *PgStorage) addUser(id, record uint, name string) {
	DB := p.db
	user := model.User{Name: name, ID: id, Record: record, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	result := DB.Create(&user)
	if result.Error != nil {
		log.Panic("Can't create user")
	}
}

func (p *PgStorage) findUser(id uint, name string) model.User {
	db := p.db
	user := model.User{}
	if result := db.First(&user, id); result.Error != nil {
		p.addUser(id, 0, name)
		db.First(&user, id)
	}
	return user
}

func (p *PgStorage) updateRecordUser(id, record uint, name string) {
	db := p.db
	user := model.User{ID: id}
	if result := db.Model(&user).Updates(model.User{UpdatedAt: time.Now(), Record: record}); result.RowsAffected == 0 {
		p.addUser(id, record, name)
	}
}

func (p *PgStorage) getUsersOrderedByRecord() []model.User {
	db := p.db
	users := []model.User{}
	db.Limit(10).Order("record desc").Find(&users)
	return users
}

func (p *PgStorage) countOfUsers() int64 {
	db := p.db
	var count int64
	db.Model(model.User{}).Count(&count)
	return count
}
