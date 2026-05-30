package benchmark

import (
	"encoding/json"
	"fmt"
	"os"

	g "github.com/Yuseonr/uav-h3-pathfinding/algorithms/internal/graph"
)

/*
Yuseonr

Deskripsi   :
 - test_mapper.go berfungsi untuk ngemapping testcases dari json ke in memory yang akan di evaluasi
 - RawTestCase -> struktur asli dari file JSON yang dihasilkan oleh Python.
 - String H3 masih berwujud teks mentah.
*/
type RawTestCase struct {
    Start struct {
        Lat   float64 `json:"lat"`
        Lng   float64 `json:"lng"`
        H3R9  string  `json:"h3_r9"`
        H3R10 string  `json:"h3_r10"`
        H3R11 string  `json:"h3_r11"`
    } `json:"start"`
    Goal struct {
        Lat   float64 `json:"lat"`
        Lng   float64 `json:"lng"`
        H3R9  string  `json:"h3_r9"`
        H3R10 string  `json:"h3_r10"`
        H3R11 string  `json:"h3_r11"`
    } `json:"goal"`
    Kategori   string  `json:"kategori"`
    JarakMeter float64 `json:"jarak_meter"`
}

func (r *RawTestCase) StartH3(resolusi int) string {
    switch resolusi {
    case 9:
        return r.Start.H3R9
    case 10:
        return r.Start.H3R10
    case 11:
        return r.Start.H3R11
    default:
        return ""
    }
}

func (r *RawTestCase) GoalH3(resolusi int) string {
    switch resolusi {
    case 9:
        return r.Goal.H3R9
    case 10:
        return r.Goal.H3R10
    case 11:
        return r.Goal.H3R11
    default:
        return ""
    }
}


/*
Yuseonr

Deskripsi   : 
 - TestCase berfungsi untuk menyimpan data test case yang udah di load
*/
type TestCase struct {
    StartID    int32
    GoalID     int32
    Kategori   string
    JarakMeter float64
    Resolusi   int
}

/*
Yuseonr

Deskripsi	:
 - LoadTestCase berfungsi untuk meload test case dari json file untuk dievaluasi
*/
func LoadTestCases(jsonpath string, mapper *g.IDMapper, resolusi int) ([]TestCase, error) {
    f, err := os.Open(jsonpath)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    var raws []RawTestCase
    if err := json.NewDecoder(f).Decode(&raws); err != nil {
        return nil, err
    }

    var hasilTestCases []TestCase
    for _, raw := range raws {
        startID, startOK := mapper.GetID(raw.StartH3(resolusi)) 
		if !startOK {
			return nil, fmt.Errorf("Gagal mendapatkan Start id")
		}
        goalID, goalOK   := mapper.GetID(raw.GoalH3(resolusi))  
		if !goalOK {
			return nil, fmt.Errorf("Gagal mendapatkan Goal id")
		}

        hasilTestCases = append(hasilTestCases, 
		TestCase{
            StartID:    startID,
            GoalID:     goalID,
            Kategori:   raw.Kategori,
            JarakMeter: raw.JarakMeter,
            Resolusi:   resolusi,
        })
    }
    return hasilTestCases, nil
}