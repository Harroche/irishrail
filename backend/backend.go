package main

import (
	"encoding/json"
	"fmt"
	"irishrail/backend/api"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/irishrail/stations", GetAllStations)
	http.HandleFunc("/irishrail/livetrains", GetAllLiveTrains)
	http.HandleFunc("/irishrail/stations/", GetStation)

	fmt.Println("Server is running localhost:12345")
	log.Fatal(http.ListenAndServe(":12345", nil))

}

// "/irishrail/stations"
func GetAllStations(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed needs to be get", http.StatusMethodNotAllowed)
		return
	}
	stations, err := api.GetDartStations()
	if err != nil {
		http.Error(w, fmt.Sprintf("error occured with getting train station data: %v", err), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"count": len(stations.ObjStations),
		"data":  stations.ObjStations,
	})
}

// "/irishrail/livetrains"
func GetAllLiveTrains(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed needs to be get", http.StatusMethodNotAllowed)
		return
	}
	liveTrains, err := api.GetTrainPositionsLive()
	if err != nil {
		http.Error(w, fmt.Sprintf("error occured with getting all live trains data: %v", err), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"count": len(liveTrains.ObjTrainPositions),
		"data":  liveTrains.ObjTrainPositions,
	})
}

// "/irishrail/stations/{ID}
func GetStation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed needs to be get", http.StatusMethodNotAllowed)
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/irishrail/stations/")
	if id == "" {
		http.Error(w, "failed to dervie an id from the call", http.StatusBadRequest)
	}
	stationInfo, err := api.GetDartStation(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error occured with getting all station data: %v", err), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"count": len(stationInfo.ObjStationsData),
		"data":  stationInfo.ObjStationsData,
	})
}
