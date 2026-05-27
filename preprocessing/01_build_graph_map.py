# Yuseonr
# Script untuk mmbuat peta hexagonal graf dari peta semarang dan obstacle yang sudah dibuat 

import h3
import geopandas as gpd
from shapely.geometry import Polygon
import json

def buat_graph_map(batas_wilayah: Polygon, gdf_obstacle: gpd.GeoDataFrame, resolution: int):
    # Membangun grid hexagonal di dalam batas wilayah
    hexagons = h3.geo_to_cells(batas_wilayah.__geo_interface__, resolution)
    index_obst = gdf_obstacle.sindex

    # Menyiapkan wadah hexagon valid
    set_hex_valid = set()
    koordinat_hex = {}

    # Memeriksa setiap hexagon apakah beririsan dengan rintangan
    for hex_id in hexagons:
        # Mendapatkan batas hexagon sebagai polygon
        batas = h3.cell_to_boundary(hex_id)
        # Membuat polygon dari batas hexagon
        poly = Polygon([(lon, lat) for lat, lon in batas])
        # Mencari kandidat rintangan yang mungkin beririsan dengan hexagon
        candidates = gdf_obstacle.iloc[list(index_obst.intersection(poly.bounds))]
        # Jika ada kandidat yang beririsan, lewati hexagon ini
        if candidates.intersects(poly).any():
            continue
        # Jika tidak ada rintangan yang beririsan, tambahkan hexagon ini ke set valid
        set_hex_valid.add(hex_id)
        koordinat_hex[hex_id] = h3.cell_to_latlng(hex_id)

    # Membangun representasi graf dari hexagon valid
    graph = {}
    for hex_id, (lat, lng) in koordinat_hex.items():
        # Mendapatkan semua tetangga dari hexagon ini
        semua_tetangga = h3.grid_disk(hex_id, 1)
        # Tetangga adalah hexagon yang berjarak 1 langkah (grid_disk dengan radius 1) dan juga valid (ada di set_hex_valid)
        tetangga_valid = [n for n in semua_tetangga if n != hex_id and n in set_hex_valid]

        # Hanya buat node untuk hexagon yang valid
        graph[hex_id] = {
            "lat": lat,
            "lng": lng,
            "cost": 1,
            "neighbors": tetangga_valid
        }

    return graph


if __name__ == "__main__":

    batas_wilayah = gpd.read_file("data/raw/batas-wilayah-semarang.geojson").dissolve().geometry.iloc[0]
    obstacles = gpd.read_file("data/processed/master-obstacles.geojson")

    # Buat graph map untuk Resolusi 9, 10, dan 11
    list_resolusi = [9,10,11]
    for resolution in list_resolusi:
        graph_representation = buat_graph_map(batas_wilayah, obstacles, resolution)
        with open(f"data/processed/h3_res{resolution}_graph.json", "w") as f:
            json.dump(graph_representation, f, indent=2)
        print(f"Graf Resolusi {resolution} disimpan")







