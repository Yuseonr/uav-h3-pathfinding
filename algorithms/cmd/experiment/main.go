package main

import (
	"fmt"
	"log"
	"time"

	b "github.com/Yuseonr/uav-h3-pathfinding/algorithms/internal/benchmark"
	g "github.com/Yuseonr/uav-h3-pathfinding/algorithms/internal/graph"
	s "github.com/Yuseonr/uav-h3-pathfinding/algorithms/internal/search"
)

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

	algoritma := map[string]s.Algorithm{
		"BFS": &s.BFS{},
		"DFS" : &s.DFS{},
		"GLS" : &s.GLS{},
	}

	tcID := 1

	for _, tc := range testCases {
		for namaAlgo, algo := range algoritma {
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