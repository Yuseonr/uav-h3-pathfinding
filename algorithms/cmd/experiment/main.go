package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	b "github.com/Yuseonr/uav-h3-pathfinding/algorithms/internal/benchmark"
	g "github.com/Yuseonr/uav-h3-pathfinding/algorithms/internal/graph"
	s "github.com/Yuseonr/uav-h3-pathfinding/algorithms/internal/search"
)

func selectJarakPerStep(resolusi int) float64 {
	var jarak float64
	switch resolusi {
	case 9:
		jarak = 347.76
	case 10:
		jarak = 131.39
	case 11:
		jarak = 49.64
	default:
		jarak = 1.0
	}
	return jarak
}

func main() {
	resultDir := "../../../data/result"
	os.MkdirAll(resultDir, os.ModePerm)

	csvPath := filepath.Join(resultDir, "Allbenchmark.csv")
	csvFile, err := os.Create(csvPath)
	if err != nil {
		log.Fatalf("Gagal membuat file hasil: %v", err)
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()
	csvWriter.Write(b.GetCSVHeader())

	resolusiList := []int{9, 10, 11}

	for _, res := range resolusiList {
		resolusi := res
		grafPath := fmt.Sprintf("../../../data/processed/h3_res%d_graph.json", resolusi)
		testCasePath := fmt.Sprintf("../../../test_cases/test-case-solvable-res%d.json", resolusi)

		graf, err := g.LoadJsonGraph(grafPath)
		if err != nil {
			log.Fatalf("Gagal muat graf res %d: %v", resolusi, err)
		}

		testCases, err := b.LoadTestCases(testCasePath, graf.Mapper, resolusi)
		if err != nil {
			log.Fatalf("Gagal muat test cases res %d: %v", resolusi, err)
		}

		dummyAlgo := &s.BFS{}
		dummyAlgo.SearchPath(graf, testCases[0].StartID, testCases[0].GoalID)
		runtime.GC()

		algoritma := map[string]s.Algorithm{
			"BFS":    &s.BFS{},
			"DFS":    &s.DFS{},
			"GLS":    &s.GLS{},
			"UCS":    &s.UCS{JarakPerStep: selectJarakPerStep(resolusi)},
			"GBFS":   &s.GBFS{},
			"A*":     &s.ASTAR{JarakPerStep: selectJarakPerStep(resolusi)},
			"Theta*": &s.THETASTAR{JarakPerStep: selectJarakPerStep(resolusi)},
		}

		tcID := 1

		for _, tc := range testCases {
			for namaAlgo, algo := range algoritma {
				runtime.GC()
				waktuMulai := time.Now()
				resPath := algo.SearchPath(graf, tc.StartID, tc.GoalID)
				durasi := time.Since(waktuMulai).Microseconds()

				status := "Failed"
				jarakRute := 0.0
				banyakLoncatan := 0
				pathRelativeLoc := ""

				startH3, _ := graf.Mapper.GetH3(tc.StartID)
				goalH3, _ := graf.Mapper.GetH3(tc.GoalID)

				if resPath.Found {
					status = "Success"
					jarakRute = s.HitungJarakRute(graf, resPath.PathIDs)
					banyakLoncatan = len(resPath.PathIDs) - 1
					if banyakLoncatan < 0 {
						banyakLoncatan = 0
					}

					pathDir := filepath.Join(resultDir, "paths", fmt.Sprintf("res%d", resolusi))
					os.MkdirAll(pathDir, os.ModePerm)

					pathFileName := fmt.Sprintf("tc%d_%s.csv", tcID, namaAlgo)
					pathFilePath := filepath.Join(pathDir, pathFileName)
					pathRelativeLoc = filepath.Join("paths", fmt.Sprintf("res%d", resolusi), pathFileName)

					pf, _ := os.Create(pathFilePath)
					for _, pid := range resPath.PathIDs {
						h3Str, _ := graf.Mapper.GetH3(pid)
						fmt.Fprintln(pf, h3Str)
					}
					pf.Close()
				}

				fr := b.FullResult{
					IdTestCase:            tcID,
					Algoritma:             namaAlgo,
					ResolusiMap:           resolusi,
					KategoriJarak:         tc.Kategori,
					StartHexID:            startH3,
					GoalHexID:             goalH3,
					BanyakLoncatan:        banyakLoncatan,
					TotalJarakMeter:       jarakRute,
					EuclideanJarakMeter:   tc.JarakMeter,
					WaktuEksekusiMicrosec: durasi,
					NodesExpanded:         resPath.NodesExpanded,
					NodesGenerated:        resPath.NodesGenerated,
					Status:                status,
					JalurHexIDs:           []string{pathRelativeLoc},
				}

				csvWriter.Write(fr.ToCSVRow())

				fmt.Printf("Res[%d] [%-6s] TC: %-3d | Kategori: %-6s | Sukses: %-5v | Waktu: %7d µs | Expanded: %5d | Jarak: %8.2f m | Ref: %8.2f m | Hops : %d\n",
					resolusi, namaAlgo, tcID, tc.Kategori, resPath.Found, durasi, resPath.NodesExpanded, jarakRute, tc.JarakMeter, banyakLoncatan)
			}
			tcID++
		}
	}
}