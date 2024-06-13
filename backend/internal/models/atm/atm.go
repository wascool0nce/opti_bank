package atm

type Atm struct {
	ID        int64                  `json:"atm_id"`
	Address   string                 `json:"address"`
	Latitude  float64                `json:"latitude"`
	Longitude float64                `json:"longitude"`
	IsAllDay  bool                   `json:"allday"`
	Services  map[string]interface{} `json:"services"`
	Time      int64                  `json:"time"`
	Queue     int64                  `json:"queue"`
}
