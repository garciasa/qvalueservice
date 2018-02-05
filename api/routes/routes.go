package routes

import (
	"net/http"

	"github.com/garciasa/qvalueservice/api"
	"github.com/gorilla/mux"
)

// NewRouter creating new router
func NewRouter(api api.BiologyQvalueService) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/qvalue/all", api.GetQvalues).Methods(http.MethodGet)
	router.HandleFunc("/qvalue/{stationcode}", api.GetQvalueByMonitoringStation).Methods(http.MethodGet)

	return router
}
