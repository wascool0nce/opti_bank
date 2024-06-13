package app

import (
	"github.com/team-for-more-tech-5-0/opti-bank-backend.git/internal/database"
	"github.com/team-for-more-tech-5-0/opti-bank-backend.git/internal/transport"
)

// Здесь мы собираем все модули вместе
func Start() {
	transport.Transport()
	database.CloseConnection()
}
