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
    stations =  api.GetAllStations()
    for station in stations:
        lonsStation.append(station.StationLongitude)
        latsStation.append(station.StationLatitude)
    
    dots, = ax.plot(
        lonsStation, latsStation, 
        marker='s', color='green', linestyle='None', markersize=3, 
        transform=ccrs.PlateCarree()
    )

    pl.show()
            

if __name__ == "__main__":
   main()