package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // module to connect to postgres
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// DB represent sqlx connection
var DB *sqlx.DB

// SetupDatabase used to configure database
func SetupDatabase() *sqlx.DB {
	// Build a DSN e.g. postgres://username:password@url.com:5432/db
	DSN := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		viper.GetString("database.postgres.username"),
		viper.GetString("database.postgres.password"),
		viper.GetString("database.postgres.host"),
		viper.GetInt("database.postgres.port"),
		viper.GetString("database.postgres.db"),
	)

	var err error

	DB, err = sqlx.Open("postgres", DSN)
	if err != nil {
		zap.S().Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		zap.S().Fatal(err)
	}

	return DB
}

// GetInstance returns connection
func GetInstance() *sqlx.DB { return DB }
