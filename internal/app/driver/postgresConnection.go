package driver

import (
	"fmt"
	"log"
	"vothings/internal/app/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnectionPostgres(cfg *domain.Configs) *gorm.DB {
	dsn := fmt.Sprintf("host = %s user = %s password = %s dbname = %s port = %s sslmode=disable",
		cfg.Postgres.Host, cfg.Postgres.UserName, cfg.Postgres.Password, cfg.Postgres.DbName, cfg.Postgres.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when try to open connection: %s", err)
	}
	return db
}
