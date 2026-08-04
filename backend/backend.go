package main

import (
	"fmt"
	"irishrail/backend/api"
)

func main() {
	arrayOfObjStation, err := api.GetDartStations()
	if err != nil {
		fmt.Printf("produced an err %w", err)
		return
	}
	for i, station := range arrayOfObjStation.ObjStations {
		fmt.Printf("i(%d), station %v", i, station)
		fmt.Println()
	}

	arrayOfObjTrainPositions, err := api.GetTrainPositionsLive()
	if err != nil {
		fmt.Printf("produced an err %w", err)
		return
	}
	for i, trains := range arrayOfObjTrainPositions.ObjTrainPositions {
		fmt.Printf("i(%d), trains %v", i, trains)
		fmt.Println()
	}

}

func printErr(err error) {
	if err != nil {
		fmt.Printf("%w", err)
	}
}
