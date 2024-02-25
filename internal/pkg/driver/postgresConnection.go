package driver

import (
	"fmt"
	"vothings/internal/app/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnectionPostgres(cfg *domain.Configs) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host = %s user = %s password = %s dbname = %s port = %s sslmode=disable",
		cfg.Postgres.Host, cfg.Postgres.UserName, cfg.Postgres.Password, cfg.Postgres.DbName, cfg.Postgres.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("error when try to open connection: %s", err)
	}
	return db, nil
}
