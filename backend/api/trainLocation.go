package api

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type ArrayOfObjTrainPositions struct {
	ObjTrainPositions []ObjTrainPositions `xml:"objTrainPositions"`
}

type ObjTrainPositions struct {
	TrainStatus    string  `xml:"TrainStatus"`
	TrainLatitude  float64 `xml:"TrainLatitude"`
	TrainLongitude float64 `xml:"TrainLongitude"`
	TrainCode      string  `xml:"TrainCode"`
	TrainDate      string  `xml:"TrainDate"`
	PublicMessage  string  `xml:"PublicMessage"`
	Direction      string  `xml:"Direction"`
}

func GetTrainPositionsLive() (*ArrayOfObjTrainPositions, error) {
	data, err := callTrainsData()
	if err != nil {
		return nil, err
	}
	trains, err := xmlToTrain(data)
	if err != nil {
		return nil, err
	}
	return trains, nil
}

func callTrainsData() ([]byte, error) {
	res, err := http.Get(TRAINS_LOCATION_GET)
	if err != nil {
		return nil, fmt.Errorf("failed to call get on %q: %w", TRAINS_LOCATION_GET, err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func xmlToTrain(data []byte) (*ArrayOfObjTrainPositions, error) {
	objTrainPos := new(ArrayOfObjTrainPositions)
	if err := xml.Unmarshal(data, &objTrainPos); err != nil {
		return nil, fmt.Errorf("failed in the unmarshal: %w", err)
	}
	return objTrainPos, nil
}
