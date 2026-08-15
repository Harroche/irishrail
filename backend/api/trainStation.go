package api

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type ArrayOfObjStation struct {
	ObjStations []ObjStation `xml:"objStation"`
}

type ObjStation struct {
	StationDesc      string  `xml:"StationDesc"`
	StationAlias     string  `xml:"StationAlias"`
	StationLatitude  float64 `xml:"StationLatitude"`
	StationLongitude float64 `xml:"StationLongitude"`
	StationCode      string  `xml:"StationCode"`
	StationId        int32   `xml:"StationId"`
}

func GetDartStations() (*ArrayOfObjStation, error) {
	data, err := callAPIStations()
	if err != nil {
		return nil, err
	}
	arrayOfObjStation, err := xmlToStations(data)
	if err != nil {
		return nil, err
	}
	return arrayOfObjStation, nil
}

func callAPIStations() ([]byte, error) {
	res, err := http.Get(TRAIN_STATIONS_GET)
	if err != nil {
		return nil, fmt.Errorf("failed to in get operation for %q: %w", TRAIN_STATION_GET, err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func xmlToStations(xmlString []byte) (*ArrayOfObjStation, error) {
	stations := new(ArrayOfObjStation)
	err := xml.Unmarshal(xmlString, stations)
	if err != nil {
		return nil, fmt.Errorf("failed unmarsahl: %w", err)
	}
	return stations, nil
}
