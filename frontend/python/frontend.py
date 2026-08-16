import matplotlib.pyplot as pl
import numpy as np
import matplotlib.animation as animation


import cartopy.crs as ccrs
import cartopy.io.img_tiles as cimgt
import cartopy.feature as cfeature



def main() -> None:


    fig = pl.figure(figsize=(17,10))
    osm = cimgt.OSM()
    ax = fig.add_subplot(1,1,1,projection=osm.crs)

    ax.set_extent([-6.3500,  -6.0000, 53.4500, 53.1000], crs=ccrs.PlateCarree())

    ax.add_image(osm,10)
    lons = [-6.27]
    lats = [53.33]
    dots, = ax.plot(
        lons, lats, 
        marker='o', color='red', linestyle='None', markersize=9, 
        transform=ccrs.PlateCarree()
    )


   # 3. Toggle visibility directly on the plot object (no global/nonlocal needed)
    # def toggle_dots(frame):
    #     dots.set_visible(not dots.get_visible())
    #     return dots,

    # # 4. Assign animation object
    # ani = animation.FuncAnimation(fig, toggle_dots, interval=1500, blit=False)
    pl.ion()
    pl.show()
    # return ani

    try:
        while pl.fignum_exists(fig.number):  # Continues while map window is open
            lons[0]+=0.02
            dots.set_data(lons,lats)
            dots.set_visible(not dots.get_visible())
            fig.canvas.draw_idle()  # Request canvas update
            pl.pause(1.5)           # Pause 1.5 seconds while keeping GUI responsive
    except KeyboardInterrupt:
        pass
            

if __name__ == "__main__":
   main()