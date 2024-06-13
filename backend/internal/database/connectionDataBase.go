package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/team-for-more-tech-5-0/opti-bank-backend.git/internal/config"
	"log"
	"sync"
)

var (
	once sync.Once
	db   *sql.DB
	err  error
)

func GetDatabase() (*sql.DB, error) {
	once.Do(func() {
		// Строка подключения к базе данных
		connectionString := config.ConnectionDataBaseString

		// Устанавливаем соединение с базой данных
		db, err = sql.Open(config.DbName, connectionString)
		if err != nil {
			log.Fatal(err)
		}

		// Проверяем соединение
		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}
	})

	return db, err
}
