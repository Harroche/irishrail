package api

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type ArrayOfObjStationData struct {
	ObjStationsData []ObjStationData `xml:"objStationData"`
}

type ObjStationData struct {
	Servertime      string `xml:"Servertime"`
	Traincode       string `xml:"Traincode"`
	Stationfullname string `xml:"Stationfullname"`
	Stationcode     string `xml:"Stationcode"`
	Querytime       string `xml:"Querytime"`
	Traindate       string `xml:"Traindate"`
	Origin          string `xml:"Origin"`
	Destination     string `xml:"Destination"`
	Origintime      string `xml:"Origintime"`
	Destinationtime string `xml:"Destinationtime"`
	Status          string `xml:"Status"`
	Lastlocation    string `xml:"Lastlocation"`
	Duein           int32  `xml:"Duein"`
	Late            int32  `xml:"Late"`
	Exparrival      string `xml:"Exparrival"`
	Expdepart       string `xml:"Expdepart"`
	Scharrival      string `xml:"Scharrival"`
	Schdepart       string `xml:"Schdepart"`
	Direction       string `xml:"Direction"`
	Traintype       string `xml:"Traintype"`
	Locationtype    string `xml:"Locationtype"`
}

func GetDartStation(stationID string) (*ArrayOfObjStationData, error) {
	data, err := callApiStation(stationID)
	if err != nil {
		return nil, err
	}
	stationData, err := xmlToStationsData(data)
	if err != nil {
		return nil, err
	}
	return stationData, nil
}

func callApiStation(stationID string) ([]byte, error) {
	res, err := http.Get(fmt.Sprintf("%s%s", TRAIN_STATION_GET, stationID))
	if err != nil {
		return nil, fmt.Errorf("failed call on the speific station %q: %w", stationID, err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func xmlToStationsData(xmlString []byte) (*ArrayOfObjStationData, error) {
	stationsData := new(ArrayOfObjStationData)
	if err := xml.Unmarshal(xmlString, stationsData); err != nil {
		return nil, fmt.Errorf("failed unmarshal: %w", err)
	}
	return stationsData, nil
}
