package main

import (
	"fmt"
	"log"
	"time"
	"runtime"

	b "github.com/Yuseonr/uav-h3-pathfinding/algorithms/internal/benchmark"
	g "github.com/Yuseonr/uav-h3-pathfinding/algorithms/internal/graph"
	s "github.com/Yuseonr/uav-h3-pathfinding/algorithms/internal/search"
)

func selectJarakPerStep(resolusi int) float64 {
	var jarak float64
	switch resolusi {
	case 9:
		jarak = 347.76 // jarak antar centeroid hexagon resolusi 9 ( 200.78 * sqrt(3))
	case 10:
		jarak = 131.39 // jarak antar centeroid hexagon resolusi 10 ( 75.86 * sqrt(3))
	case 11:
		jarak = 49.64 // jarak antar centeroid hexagon resolusi 11 ( 28.66 * sqrt(3))
	default:
		jarak = 1.0 
	}
	return jarak
}

func main() {
	resolusi := 11
	grafPath := fmt.Sprintf("../../../data/processed/h3_res%d_graph.json", resolusi)
	testCasePath := fmt.Sprintf("../../../test_cases/test-case-solvable-res%d.json", resolusi)

	graf, err := g.LoadJsonGraph(grafPath)
	if err != nil {
		log.Fatalf("Gagal muat graf: %v", err)
	}

	testCases, err := b.LoadTestCases(testCasePath, graf.Mapper, resolusi)
	if err != nil {
		log.Fatalf("Gagal muat test cases: %v", err)
	}

	// dummy run untuk pemanasan JIT compiler dan cache
	dummyAlgo := &s.BFS{}
	dummyAlgo.SearchPath(graf, testCases[0].StartID, testCases[0].GoalID)
	runtime.GC()

	algoritma := map[string]s.Algorithm{
		"BFS": &s.BFS{},
		"DFS" : &s.DFS{},
		"GLS" : &s.GLS{},
		"UCS" : &s.UCS{ JarakPerStep: selectJarakPerStep(resolusi)},
		"GBFS": &s.GBFS{},
		"A*" : &s.ASTAR{ JarakPerStep: selectJarakPerStep(resolusi)},
		"Theta*" : &s.THETASTAR{ JarakPerStep: selectJarakPerStep(resolusi)},
	}

	tcID := 1

	for _, tc := range testCases {
		for namaAlgo, algo := range algoritma {
			// panggil garbage collector sebelum setiap pencarian untuk mengurangi noise pada pengukuran waktu
			runtime.GC()
			waktuMulai := time.Now()
			res := algo.SearchPath(graf, tc.StartID, tc.GoalID)
			durasi := time.Since(waktuMulai).Microseconds()

			jarakRute := 0.0
			if res.Found {
				jarakRute = s.HitungJarakRute(graf, res.PathIDs)
			}

			fmt.Printf("[%s] TC: %-3d | Kategori: %-6s | Sukses: %-5v | Waktu: %6d µs | Expanded: %5d | Jarak: %8.2f m | Ref: %8.2f m | Hops : %d\n",
				namaAlgo, tcID, tc.Kategori, res.Found, durasi, res.NodesExpanded, jarakRute, tc.JarakMeter, len(res.PathIDs))
		}
		tcID++
	}
}