package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/PMahdiDehghaniP/Resturan_Panel_Web_API/config"
	logging "github.com/PMahdiDehghaniP/Resturan_Panel_Web_API/logger"

	_ "github.com/lib/pq"
)

var dbClient *sql.DB
var logger = logging.NewLogger(config.GetConfig())

func InitPostgresDB(cfg *config.Config) {
	postgresConfig := cfg.Postgres

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		postgresConfig.Host,
		postgresConfig.Port,
		postgresConfig.User,
		postgresConfig.Password,
		postgresConfig.DbName,
		postgresConfig.SSLMode,
	)

	logger.Info(logging.Postgres, logging.StartUp, "Connecting To Postgres ...", nil)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Fatal(logging.Postgres, logging.StartUp, err.Error(), nil)
	}

	if err = db.Ping(); err != nil {
		logger.Fatal(logging.Postgres, logging.StartUp, err.Error(), nil)
	}

	db.SetMaxIdleConns(postgresConfig.MaxIdleConnections)
	db.SetMaxOpenConns(postgresConfig.MaxOpenConnections)
	db.SetConnMaxLifetime(postgresConfig.MaxLifetimeConnection * time.Second)
	db.SetConnMaxIdleTime(postgresConfig.MaxIdleTimeConnection * time.Minute)

	dbClient = db

	logger.Info(logging.Postgres, logging.StartUp, "Connected to Postgres", nil)
}

func GetPostgresDB() *sql.DB {
	if dbClient == nil {
		logger.Fatal(logging.Postgres, logging.Get, "Database is not initialized", nil)
	}
	return dbClient
}

func ClosePostgresDB() {
	if dbClient != nil {
		logger.Info(logging.Postgres, logging.Shutdown, "Closing Postgres DB", nil)

		if err := dbClient.Close(); err != nil {
			logger.Error(logging.Postgres, logging.Shutdown, err.Error(), nil)
		}
	}
}
