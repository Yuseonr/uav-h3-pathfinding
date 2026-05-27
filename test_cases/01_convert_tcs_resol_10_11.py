# Yuseonr
# Script untuk Mengkonversi test case pasangan titik Start-Goal yang valid dari resolusi 9 ke resolusi 10 dan 11
# untuk digunakan dalam evaluasi algoritma pathfinding UAV di atas peta Semarang

import json
import h3

def convert_tcs_resol_10_11():
    with open("test_cases/test-case-solvable-res9.json", "r") as f:
        test_cases = json.load(f)
    with open("data/processed/h3_res10_graph.json", "r") as f:
        graph_r10 = json.load(f)
    with open("data/processed/h3_res11_graph.json", "r") as f:
        graph_r11 = json.load(f)

    test_cases_r10 = []
    test_cases_r11 = []

    hilang_r10 = 0
    hilang_r11 = 0

    for case in test_cases:
        lat_start, lng_start = case["start"]["lat"], case["start"]["lng"]
        lat_goal, lng_goal = case["goal"]["lat"], case["goal"]["lng"]

        h3_r10_start = h3.latlng_to_cell(lat_start, lng_start, 10)
        h3_r10_goal = h3.latlng_to_cell(lat_goal, lng_goal, 10)

        # Validasi: Apakah masih aman dari obstacle di R10?
        if h3_r10_start in graph_r10 and h3_r10_goal in graph_r10:
            case_r10 = case.copy()
            case_r10["start"] = {"lat": lat_start, "lng": lng_start, "h3_r10": h3_r10_start}
            case_r10["goal"] = {"lat": lat_goal, "lng": lng_goal, "h3_r10": h3_r10_goal}
            test_cases_r10.append(case_r10)
        else:
            hilang_r10 += 1

        h3_r11_start = h3.latlng_to_cell(lat_start, lng_start, 11)
        h3_r11_goal = h3.latlng_to_cell(lat_goal, lng_goal, 11)

        # Validasi: Apakah masih aman dari obstacle di R11?
        if h3_r11_start in graph_r11 and h3_r11_goal in graph_r11:
            case_r11 = case.copy()
            case_r11["start"] = {"lat": lat_start, "lng": lng_start, "h3_r11": h3_r11_start}
            case_r11["goal"] = {"lat": lat_goal, "lng": lng_goal, "h3_r11": h3_r11_goal}
            test_cases_r11.append(case_r11)
        else:
            hilang_r11 += 1

    # Simpan hasil konversi
    with open("test_cases/test-case-solvable-res10.json", "w") as f:
        json.dump(test_cases_r10, f, indent=2)
    with open("test_cases/test-case-solvable-res11.json", "w") as f:
        json.dump(test_cases_r11, f, indent=2)

    print(f"Total awal (Res 9) : {len(test_cases)} rute")
    print(f"Berhasil Res 10    : {len(test_cases_r10)} rute ({hilang_r10} rute gagal karena presisi rintangan)")
    print(f"Berhasil Res 11    : {len(test_cases_r11)} rute ({hilang_r11} rute gagal karena presisi rintangan)")

if __name__ == "__main__":
    convert_tcs_resol_10_11()