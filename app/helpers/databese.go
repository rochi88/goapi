package helpers

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabase() {
	var err error

	dbType := GetEnv("DB_CONNECTION", "pgsql")

	dsn, err := ConnectionURLBuilder(dbType)

	if err != nil {
		panic(err)
	}

	switch dbType {
	case "pgsql":
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "mysql":
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	default:
		panic("Unknown database type: " + dbType)
	}

	if err != nil {
		panic(err)
	}
}

func ConnectionURLBuilder(n string) (string, error) {
	// Define URL to connection.
	var url string

	// Switch given names.
	switch n {
	case "pgsql":
		// URL for PostgreSQL connection.
		url = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			GetEnv("DB_HOST", "127.0.0.1"),
			GetEnv("DB_PORT", "5432"),
			GetEnv("DB_USERNAME", "postgres"),
			GetEnv("DB_PASSWORD", "postgres"),
			GetEnv("DB_DATABASE", "webcore"),
			GetEnv("DB_SSL_MODE", "disable"),
		)
	case "mysql":
		// URL for Mysql connection.
		url = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			GetEnv("DB_USERNAME", "root"),
			GetEnv("DB_PASSWORD", ""),
			GetEnv("DB_HOST", "127.0.0.1"),
			GetEnv("DB_PORT", "3306"),
			GetEnv("DB_DATABASE", "webcore"),
		)
	case "redis":
		// URL for Redis connection.
		url = fmt.Sprintf(
			"%s:%s",
			GetEnv("REDIS_HOST", "127.0.0.1"),
			GetEnv("REDIS_PORT", "6379"),
		)
	case "fiber":
		// URL for Fiber connection.
		url = fmt.Sprintf(
			"%s:%s",
			GetEnv("SERVER_HOST", "127.0.0.1"),
			GetEnv("SERVER_PORT", "8080"),
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	// Return connection URL.
	return url, nil
}
