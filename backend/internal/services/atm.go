package services

import (
	"github.com/team-for-more-tech-5-0/opti-bank-backend.git/internal/database"
	"github.com/team-for-more-tech-5-0/opti-bank-backend.git/internal/models/atm"
	"log"
)

// Подсчет ближайщих банкоматов
func CalculateNearAtms(lat, lon float64, radius float64) ([]atm.Atm, error) {
	var result []atm.Atm
	radiusStep := 1.0

	//Поиск, пока не будет минимум 3 банкомата
	for len(result) < 3 {
		newResult, err := FindNearAtms(lat, lon, radius)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		result = newResult
		radius += radiusStep
		//Каждый раз увеличиваем радиус поиска
		radiusStep *= 1.5
	}
	return result, nil
}
func FindNearAtms(lat, lon float64, radius float64) ([]atm.Atm, error) {
	var result []atm.Atm

	banks, err := database.GetAtms()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for _, currentBank := range banks {
		dist := distance(lat, lon, currentBank.Latitude, currentBank.Longitude)
		if dist <= radius {
			result = append(result, currentBank)
		}
	}

	return result, nil
}
