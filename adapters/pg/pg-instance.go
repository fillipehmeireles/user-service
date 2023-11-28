package pg

import (
	"log"

	"github.com/fillipehmeireles/user-service/adapters/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgInstance struct {
	DB *gorm.DB
}

func NewPGInstance(dsn string) (*PgInstance, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("[PgInstance:NewPGInstance] error on connecting to database: %s", err.Error())
		return nil, err
	}

	return &PgInstance{
		DB: db,
	}, nil
}

func (pgInstance *PgInstance) Close() {
	sqlDB, _ := pgInstance.DB.DB()
	sqlDB.Close()
}

func (pgInstance *PgInstance) Migrate() {
	pgInstance.DB.Debug().AutoMigrate(
		&schemas.User{},
	)
}
