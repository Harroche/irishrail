import requests
from dataclasses import dataclass

LOCALHOST = "http://localhost:12345/irishrail/"
GETALLSTATIONS = "stations"
GETALLTRAINSLIVELOCATION = "livetrains"

@dataclass
class Station:
    StationDesc:      str 
    StationAlias:     str
    StationLatitude:  float
    StationLongitude: float
    StationCode:      str 
    StationId:        int  

@dataclass
class Train:
    TrainStatus:    str  
    TrainLatitude:  float 
    TrainLongitude: float 
    TrainCode:      str 
    TrainDate:      str 
    PublicMessage:  str  
    Direction:      str 


def GetAllStations()-> list[Station]:
   response  = requests.get((LOCALHOST+GETALLSTATIONS))

   resp = response.json()
   return [Station(**item) for item in resp["data"]]

def GetAllTrainsLiveLocation() -> list[Train]:
    response = requests.get((LOCALHOST+GETALLTRAINSLIVELOCATION))
    resp = response.json()
    return [Train(**item) for item in resp["data"]]
    