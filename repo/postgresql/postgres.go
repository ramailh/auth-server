package postgresql

import (
	"fmt"
	"log"

	"github.com/ramailh/auth-server/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type pq struct {
	db *gorm.DB
}

const (
	host     = "localhost"
	user     = "postgres"
	password = "password"
	db       = "postgres"
	port     = 5432
)

func NewPostgresClient() (*pq, error) {
	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, user, password, db, port)
	dbConn, err := gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	return &pq{db: dbConn}, err
}

func (pq *pq) Migrate() *pq {
	if exist := pq.db.Migrator().HasTable(&model.User{}); !exist {
		if err := pq.db.AutoMigrate(&model.User{}); err != nil {
			log.Println(err)
		}
	}

	return pq
}

func (pq *pq) Find(condition string, args ...interface{}) (model.User, error) {
	var data model.User
	err := pq.db.Table("users").Where(condition, args...).Take(&data).Error
	return data, err
}
