package client

import (
	"fmt"

	"github.com/xdevor/ginger/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(`
		host=%s
		user=%s
		password=%s
		dbname=%s
		port=%s
		sslmode=%s
		TimeZone=%s`,
		config.Database.Host,
		config.Database.User,
		config.Database.Password,
		config.Database.Database,
		config.Database.Port,
		config.Database.SSLMode,
		config.Database.Timezone,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return
}
