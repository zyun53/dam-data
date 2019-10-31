import os
import sys
import shapefile

with shapefile.Reader('./W01-14_GML/W01-14-g_Dam',encoding="cp932") as shp:
    fields = shp.fields
    print(fields)
