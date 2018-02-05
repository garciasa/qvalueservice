package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/garciasa/qvalueservice/api/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// BiologyQvalueService defines services provided by this microservice
type BiologyQvalueService interface {
	GetQvalues(w http.ResponseWriter, r *http.Request)
	GetQvalueByMonitoringStation(w http.ResponseWriter, r *http.Request)
}

type biologyService struct {
	log *logrus.Entry
	db  *sql.DB
}

// NewBiologyQvalueService generates a new BiologyQvalueService
func NewBiologyQvalueService(log *logrus.Logger, db *sql.DB) BiologyQvalueService {
	return &biologyService{
		log: log.WithField("service", "qvalue"),
		db:  db,
	}
}

// GetQvalues get all values we have in database
func (b *biologyService) GetQvalues(w http.ResponseWriter, r *http.Request) {

	var results []models.QValue

	rows, err := b.db.Query(`
		SELECT [StationCode]
	  	,[DateOfSurvey]
	    ,[QValueScore]
		FROM [vector].[sde].[MON_QRECORDS_VERSION2]
		WHERE QValueScore IS NOT NULL
	`)

	defer rows.Close()

	if err != nil {
		b.log.Fatal(err.Error())
	}

	for rows.Next() {
		var d = models.QValue{}
		if err := rows.Scan(&d.StationCode, &d.DateOfSurvey, &d.QvalueScore); err != nil {
			b.log.Fatal(err.Error())
		}
		results = append(results, d)
	}
	b.log.Info("Returning data from GetQvalues endpoint")
	response := &models.Response{
		Success: true,
		Total:   len(results),
		Data:    results,
	}
	json.NewEncoder(w).Encode(response)

}

// GetQvalueByMonitoringStation get qvalues for a specified monitoring station
func (b *biologyService) GetQvalueByMonitoringStation(w http.ResponseWriter, r *http.Request) {
	var results []models.QValue

	code := mux.Vars(r)["stationcode"]

	rows, err := b.db.Query(`
		SELECT [StationCode]
	  	,[DateOfSurvey]
	    ,[QValueScore]
		FROM [vector].[sde].[MON_QRECORDS_VERSION2]
		WHERE [StationCode] = ?
		ORDER BY [DateOfSurvey]`, code)

	defer rows.Close()

	if err != nil {
		b.log.Fatal(err.Error())
	}

	for rows.Next() {
		var d = models.QValue{}
		if err := rows.Scan(&d.StationCode, &d.DateOfSurvey, &d.QvalueScore); err != nil {
			b.log.Fatal(err.Error())
		}
		results = append(results, d)
	}

	b.log.Info("Returning data from GetQvalueByMonitoringStation endpoint")
	response := &models.Response{
		Success: true,
		Total:   len(results),
		Data:    results,
	}
	json.NewEncoder(w).Encode(response)
}
