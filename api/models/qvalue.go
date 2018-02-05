package models

// QValue represent quality of water measured in a Station
type QValue struct {
	StationCode  string `json:"stationcode"`
	DateOfSurvey string `json:"dateofsurvey"`
	QvalueScore  string `json:"qvaluescore"`
}

// Response represent a generic response for API
type Response struct {
	Success bool     `json:"success"`
	Total   int      `json:"total"`
	Data    []QValue `json:"data"`
}
