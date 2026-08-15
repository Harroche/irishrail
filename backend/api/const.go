package api

const (
	TRAIN_STATIONS_GET  = "http://api.irishrail.ie/realtime/realtime.asmx/getAllStationsXML_WithStationType?StationType=D"
	TRAIN_STATION_GET   = "https://api.irishrail.ie/realtime/realtime.asmx/getStationDataByCodeXML?StationCode="
	TRAINS_LOCATION_GET = "https://api.irishrail.ie/realtime/realtime.asmx/getCurrentTrainsXML_WithTrainType?TrainType=D"
)
