# Yuseonr
# Script untuk menggabungkan data obstacle yang ada di data/raw menjadi satu file obstacle

import geopandas as gpd
import pandas as pd
from shapely.ops import unary_union

def load(path):
    gdf = gpd.read_file(path)
    gdf["geometry"] = gdf.geometry.buffer(0) 
    return gdf

kkop = load("data/raw/zona-penerbangan-ayani.geojson")
militer = load("data/raw/zona-militer-semarang.geojson")
sutt = load("data/raw/gap-tower-sutt-semarang.geojson")
gedung = load("data/raw/gedung-tinggi-semarang.geojson")

# Satukan semua geometri menjadi satu list 
semua_rintangan = pd.concat([kkop.geometry, militer.geometry, sutt.geometry, gedung.geometry])

# Gabungkan area yang sama menjadi satu kesatuan MultiPolygon
mega_obstacle = unary_union(semua_rintangan)

# Download hasilnya ke file GeoJSON
gpd.GeoSeries([mega_obstacle], crs="EPSG:4326").to_file("data/processed/master-obstacles.geojson", driver="GeoJSON")

print("Selesai menggabungkan data obs")