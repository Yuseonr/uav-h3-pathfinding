# Yuseonr
# Script untuk menghasilkan test case pasangan titik Start-Goal yang valid
# untuk digunakan dalam evaluasi algoritma pathfinding UAV di atas peta Semarang

import json
import random
import heapq
import h3
import geopandas as gpd
from shapely.geometry import Point
from math import radians, sin, cos, sqrt, atan2

KATEGORI_JARAK = {
    "short":  (500,   3000),
    "medium": (3000,  8000),
    "long":   (8000,  50000),
}

JUMLAH_PER_KATEGORI = 50

def haversine(lat1, lng1, lat2, lng2) -> float:
    # Menghitung jarak dua koordinat dalam meter menggunakan formula Haversine
    R = 6371000
    lat1, lng1, lat2, lng2 = map(radians, [lat1, lng1, lat2, lng2])
    dlat = lat2 - lat1
    dlng = lng2 - lng1
    a = sin(dlat/2)**2 + cos(lat1) * cos(lat2) * sin(dlng/2)**2
    return R * 2 * atan2(sqrt(a), sqrt(1-a))


def random_koordinat_dalam_wilayah(batas_wilayah) -> tuple:
    # Mengacak koordinat (lat, lng) yang jatuh di dalam polygon batas wilayah
    minx, miny, maxx, maxy = batas_wilayah.bounds
    while True:
        lng = random.uniform(minx, maxx)
        lat = random.uniform(miny, maxy)
        if batas_wilayah.contains(Point(lng, lat)):
            return lat, lng


def is_jarak_sesuai_kategori(lat1, lng1, lat2, lng2, kategori: str) -> bool:
    # Mengecek apakah jarak dua titik masuk ke dalam rentang kategori yang diminta
    jarak = haversine(lat1, lng1, lat2, lng2)
    min_j, max_j = KATEGORI_JARAK[kategori]
    return min_j <= jarak <= max_j

def cek_jalur_astar(graph: dict, start: str, goal: str) -> bool:
    # Memvalidasi apakah jalur dari start ke goal MUNGKIN ADA menggunakan A*
    # Tidak menyimpan jalur — hanya mengembalikan True/False (reachability check)
    # Dijalankan pada resolusi 9 (paling restriktif) sebagai ground truth validator
    if start not in graph or goal not in graph:
        return False

    goal_lat = graph[goal]["lat"]
    goal_lng = graph[goal]["lng"]

    def heuristik(hex_id):
        # Estimasi jarak dari hex_id ke goal menggunakan Haversine
        return haversine(graph[hex_id]["lat"], graph[hex_id]["lng"], goal_lat, goal_lng)

    g_score = {start: 0}
    open_set = [(heuristik(start), start)]
    visited  = set()

    while open_set:
        _, current = heapq.heappop(open_set)

        if current == goal:
            return True

        if current in visited:
            continue

        visited.add(current)

        for tetangga in graph[current]["neighbors"]:
            if tetangga in visited or tetangga not in graph:
                continue

            # Setiap langkah antar hexagon dianggap berbiaya sama (cost = 1)
            g_baru = g_score[current] + 1

            if tetangga not in g_score or g_baru < g_score[tetangga]:
                g_score[tetangga] = g_baru
                f_score = g_baru + heuristik(tetangga)
                heapq.heappush(open_set, (f_score, tetangga))

    return False

def generate_test_cases(graph_r9: dict, batas_wilayah, jumlah: int, kategori: str) -> list:
    # Menghasilkan pasangan Start-Goal yang sudah tervalidasi:
    # 1. Jarak sesuai kategori
    # 2. Kedua titik berada di hexagon yang bisa dilalui (bukan obstacle)
    # 3. Terbukti ada jalur di resolusi 9 
    test_cases  = []
    attempts    = 0
    max_attempts = jumlah * 100

    print(f"Generating {jumlah} test cases kategori '{kategori}'...")

    while len(test_cases) < jumlah and attempts < max_attempts:
        attempts += 1

        lat1, lng1 = random_koordinat_dalam_wilayah(batas_wilayah)
        lat2, lng2 = random_koordinat_dalam_wilayah(batas_wilayah)

        # Jarak tidak sesuai kategori
        if not is_jarak_sesuai_kategori(lat1, lng1, lat2, lng2, kategori):
            continue

        start_hex = h3.latlng_to_cell(lat1, lng1, 9)
        goal_hex  = h3.latlng_to_cell(lat2, lng2, 9)

        # Start dan goal tidak boleh hex yang sama
        if start_hex == goal_hex:
            continue

        # Kedua hex harus berada di area yang bisa dilalui (bukan obstacle)
        if start_hex not in graph_r9 or goal_hex not in graph_r9:
            continue

        # Validasi topologi jalur benar-benar ada di R9
        if not cek_jalur_astar(graph_r9, start_hex, goal_hex):
            continue

        test_cases.append({
            "start":       {"lat": lat1, "lng": lng1, "h3_r9": start_hex},
            "goal":        {"lat": lat2, "lng": lng2, "h3_r9": goal_hex},
            "kategori":    kategori,
            "jarak_meter": round(haversine(lat1, lng1, lat2, lng2), 2),
        })

        print(f"  [{kategori}] {len(test_cases)}/{jumlah} valid (attempt {attempts})")

    print(f"Selesai: {len(test_cases)} test cases dari {attempts} percobaan\n")
    return test_cases


if __name__ == "__main__":
    batas_wilayah = gpd.read_file("data/raw/batas-wilayah-semarang.geojson").dissolve().geometry.iloc[0]

    with open("data/processed/h3_res9_graph.json") as f:
        graph_r9 = json.load(f)

    semua_test_cases = []

    for kategori in ["short", "medium", "long"]:
        hasil = generate_test_cases(graph_r9, batas_wilayah, JUMLAH_PER_KATEGORI, kategori)
        semua_test_cases.extend(hasil)

    with open("test_cases/test-case-solvable-res9.json", "w") as f:
        json.dump(semua_test_cases, f, indent=2)

    print("Test cases disimpan ke test_cases/test-case-solvable-res9.json")