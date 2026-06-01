package search

import (
	"math"
	c "github.com/Yuseonr/uav-h3-pathfinding/algorithms/internal/collections"
	g "github.com/Yuseonr/uav-h3-pathfinding/algorithms/internal/graph"
)

/*
Yuseonr

Deskripsi   :
 - theta_star.go implementasi algoritma Theta* (Any-Angle Pathfinding) pada graf
 - merupakan perluasan dari A* yang memungkinkan pergerakan di luar batasan edge grid
 - membutuhkan fungsi Line of Sight (LoS) untuk mengecek rintangan antar dua titik
 - jika LoS bebas rintangan, algoritma akan membuat rute shortcut dari Parent(Current) langsung ke Tetangga.
 - source : K. Daniel, A. Nash, S. Koenig, and A. Felner, “Theta*: Any-Angle Path Planning on Grids,” Journal of Artificial Intelligence Research, vol. 39, pp. 533–579, Jan. 2014, doi: 10.1613/jair.2994.
  
*/
type THETASTAR struct {
	JarakPerStep float64
}

func (theta *THETASTAR) SearchPath(graf *g.Graph, start, goal int32) Result {
	// inisialisasi min-priority queue -> selalu memberikan node dengan nilai f(n)+g(n) terkecil
	pqueue := c.CreatePriorityQueue()

	// Penjelasan costTerbaik dan asalDari :
	// 1. costTerbaik: mencatat gScore,  menyimpan nilai biaya termurah dari start ke node index.
	// 2. asalDari: cek bfs.go
	costTerbaik := make([]float64, len(graf.Nodes))
	asalDari := make([]int32, len(graf.Nodes))

	// counter
	nodesExpanded := 0
	nodesGenerated := 0

	// inisialisasi semua gScore dengan nilai Tak Terhingga (Infinity) dan asalDari dengan -1
	for i := range costTerbaik {
		costTerbaik[i] = math.Inf(1)
		asalDari[i] = -1
	}

	// ambil node goal untuk hitung heuristik
	goalHexNode := graf.Nodes[goal]

	// mulai Theta*
	// tandai hexagon awal udh dikunjungi dan cost di titik awal adalah 0
	asalDari[start] = start
	costTerbaik[start] = 0

	// kalkulasi f(n) awal = g(start) + h(start)
	startHexNode := graf.Nodes[start]
	fStart := costTerbaik[start] + Haversine(startHexNode, goalHexNode)
	pqueue.Push(start, fStart)

	// selama antrian prioritas belum kosong terus cari
	for !pqueue.IsEmpty() {
		// ambil heksagon dengan f(n)+g(n) terkecil
		currentHex, current_F_Score := pqueue.Pop()
		
		// ambil node
		currentHexNode := graf.Nodes[currentHex]
        // validasi f(n) + g(n) untuk menghindari memproses node yang buruk
		h_current := Haversine(currentHexNode, goalHexNode)
		f_score_valid := costTerbaik[currentHex] + h_current
		if current_F_Score > f_score_valid {
			continue
		}
		// hitung jumlah node yang diekspansi
		nodesExpanded++

		// jika current adalah goal, rekonstruksi path dan return hasil
		if currentHex == goal {
			path := reconstructPath(asalDari, start, goal)
			return Result{
				PathIDs:        path,
				NodesExpanded:  nodesExpanded,
				NodesGenerated: nodesGenerated,
				Found:          true,
			}
		}

		// ambil parent dari current untuk cek Line of Sight
		parentOfCurrent := asalDari[currentHex]
		parentHexNode := graf.Nodes[parentOfCurrent]

		// iterasi semua tetangga current
		for _, neighborID := range currentHexNode.Neighbors {
			neighborHexNode := graf.Nodes[neighborID]
			// cek Line of Sight antara parent(current) dan neighbor
			if LineOfSight(graf, parentOfCurrent, neighborID) {
				jarakLoS := Haversine(parentHexNode, neighborHexNode)
				temp_gScore := costTerbaik[parentOfCurrent] + jarakLoS
				// jika rute shortcut lebih murah, update cost dan parent
				if temp_gScore < costTerbaik[neighborID] {
					costTerbaik[neighborID] = temp_gScore
					asalDari[neighborID] = parentOfCurrent

					h_score := Haversine(neighborHexNode, goalHexNode)
					pqueue.Push(neighborID, temp_gScore + h_score)
					nodesGenerated++
				}
				
			// jika tidak ada LoS, fallback ke perhitungan A* biasa
			} else {
				
				temp_gScore := costTerbaik[currentHex] + theta.JarakPerStep

				if temp_gScore < costTerbaik[neighborID] {
					costTerbaik[neighborID] = temp_gScore
					asalDari[neighborID] = currentHex 

					h_score := Haversine(neighborHexNode, goalHexNode)
					pqueue.Push(neighborID, temp_gScore + h_score)
					nodesGenerated++
				}
			}
		}
	}
	// jika antrian prioritas kosong dan goal belum ditemukan, return hasil tidak ditemukan
	return Result{NodesExpanded: nodesExpanded, NodesGenerated: nodesGenerated, Found: false}
}


