import matplotlib.pyplot as pl
import numpy as np

import api
import cartopy.crs as ccrs
import cartopy.io.img_tiles as cimgt



def main() -> None:


    fig = pl.figure(figsize=(17,10))
    osm = cimgt.OSM()
    ax = fig.add_subplot(1,1,1,projection=osm.crs)

    ax.set_extent([-6.3500,  -6.0000, 53.4500, 53.1000], crs=ccrs.PlateCarree())

    ax.add_image(osm,10)
    lonsStation = []
    latsStation = []
    lonsTrains = []
    latsTrains = []
    stations =  api.GetAllStations()
    for station in stations:
        lonsStation.append(station.StationLongitude)
        latsStation.append(station.StationLatitude)
    
    dots_station, = ax.plot(
        lonsStation, latsStation, 
        marker='s', color='green', linestyle='None', markersize=3, 
        transform=ccrs.PlateCarree()
    )


    dots_trains, = ax.plot(
        lonsTrains,latsTrains,
        marker= 'o', color = 'red', linestyle='None', markersize = 4,
        transform=ccrs.PlateCarree()
    )

    pl.ion()
    pl.show()
            


    try:
        while pl.fignum_exists(fig.number): 
            trains = api.GetAllTrainsLiveLocation()
            lonsTrains.clear()
            latsTrains.clear()
            for train in trains:
                lonsTrains.append(train.TrainLongitude)
                latsTrains.append(train.TrainLatitude)
            dots_trains.set_data(lonsTrains,latsTrains)
            fig.canvas.draw_idle()  
            pl.pause(.5)           # Pause .5 seconds while keeping GUI responsive
    except KeyboardInterrupt:
        pass
if __name__ == "__main__":
   main()