import matplotlib.pyplot as pl
import numpy as np

import api
import cartopy.crs as ccrs
import cartopy.io.img_tiles as cimgt


NORTH = "Northbound"

def main() -> None:


    fig = pl.figure(figsize=(17,10))
    osm = cimgt.OSM()
    ax = fig.add_subplot(1,1,1,projection=osm.crs)

    ax.set_extent([-6.3500,  -6.0000, 53.4500, 53.1000], crs=ccrs.PlateCarree())

    ax.add_image(osm,10)
    lonsStation = []
    latsStation = []
    stations =  api.GetAllStations()
    for station in stations:
        lonsStation.append(station.StationLongitude)
        latsStation.append(station.StationLatitude)
    
    ax.plot(
        lonsStation, latsStation, 
        marker='s', color='black', linestyle='None', markersize=3, 
        transform=ccrs.PlateCarree()
    )


    dots_trains_n, = ax.plot(
        [],[],
        marker= 'o', color = 'red', linestyle='None', markersize = 4,
        transform=ccrs.PlateCarree()
    )
    dots_trains_s, = ax.plot(
        [],[],
        marker= 'o', color = 'green', linestyle='None', markersize = 4,
        transform=ccrs.PlateCarree()
    )

    pl.ion()
    pl.show()
            

    try:
        while pl.fignum_exists(fig.number): 
            trains = getTrains()
            dots_trains_n.set_data(trains[0],trains[1])
            dots_trains_s.set_data(trains[2],trains[3])
            fig.canvas.draw_idle()  
            pl.pause(.5)           # Pause .5 seconds while keeping GUI responsive
    except KeyboardInterrupt:
        pass

def getTrains() -> list[list[float]]:
    trains = api.GetAllTrainsLiveLocation()
    nLong = []
    nLat = []
    sLong = []
    sLat = []
    for train in trains:
        if train.Direction == NORTH:
            nLong.append(train.TrainLongitude)
            nLat.append(train.TrainLatitude)
        else:
            sLong.append(train.TrainLongitude)
            sLat.append(train.TrainLatitude)
    return [nLong,nLat,sLong,sLat]


if __name__ == "__main__":
   main()