package database

// Закрываем соединение при завершении приложения
func CloseConnection() {
	db, err := GetDatabase()
	if err != nil {
		panic(err)
	}

	err = db.Close()
	if err != nil {
		panic(err)
	}
}
