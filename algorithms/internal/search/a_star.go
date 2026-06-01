package search

import (
	"math"
	c "github.com/Yuseonr/uav-h3-pathfinding/algorithms/internal/collections"
	g "github.com/Yuseonr/uav-h3-pathfinding/algorithms/internal/graph"
)

/*
Yuseonr

Deskripsi   :
 - a_star.go implementasi algoritma A* (A-Star) untuk pencarian jalur pada graf
 - A* menggunakan fungsi evaluasi f(n) = g(n) + h(n)
 - g(n) dihitung menggunakan konstanta JarakPerStep
 - h(n) dihitung menggunakan jarak spasial Euclidean/Haversine absolut ke titik tujuan.
 - menjamin penemuan rute terpendek secara geometris (Optimal & Complete).
*/
type ASTAR struct {
	JarakPerStep float64
}

func (astar *ASTAR) SearchPath(graf *g.Graph, start, goal int32) Result {
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

	// mulai A*
	// tandai hexagon awal udh dikunjungi dan cost di titik awal adalah 0
	asalDari[start] = start
	costTerbaik[start] = 0

	// kalkulasi f(n) awal = g(start) + h(start)
	startHexNode := graf.Nodes[start]
	fStart := costTerbaik[start] + Haversine(startHexNode, goalHexNode)

	// push heksagon awal ke antrian prioritas dengan prioritas f(n)
	pqueue.Push(start, fStart)

	// selama antrian prioritas belum kosong terus cari
	for !pqueue.IsEmpty() {
		// ambil heksagon dengan f(n) terkecil
		currentHex, current_F_Score := pqueue.Pop()

		// ambil nodenya
		currentHexNode := graf.Nodes[currentHex]

		// Skip entri lama yang sudah tidak merepresentasikan cost terbaik.
        // karena priority queue tidak mendukung update prioritas, entri lama
        // tetap dibiarkan di antrian tapi langsug si skip saat dipop.
		h_current := Haversine(currentHexNode, goalHexNode)
		f_score_valid := costTerbaik[currentHex] + h_current
		if current_F_Score > f_score_valid {
			continue
		}

		// counter node yang diekspansi /  dikeluarkan dari antrian prioritas
        nodesExpanded++

		// jika ini adalah goal, return hasilnya dengan merekonstruksi jalur dari asalDari
		if currentHex == goal {
			path := reconstructPath(asalDari, start, goal)
			return Result{
				PathIDs:        path,
				NodesExpanded:  nodesExpanded,
				NodesGenerated: nodesGenerated,
				Found:          true,
			}
		}

		// current node bukan goal
		// untuk setiap tetangga, evaluasi f(n) baru
		for _, neighborID := range currentHexNode.Neighbors {

			// kalkulasi g(n) : jarak dari start ke currentHex + jarak 1 langkah ke depan
			temp_gScore := costTerbaik[currentHex] + astar.JarakPerStep

			// jika rute via currentHex ini lebih pendek dari rekor gScore sebelumnya
			if temp_gScore < costTerbaik[neighborID] {

				// perbarui rekor jarak 
				costTerbaik[neighborID] = temp_gScore

				// catat dari mana kita datang
				asalDari[neighborID] = currentHex

				// kalkulasi h(n): estimasi sisa jarak menggunakan Haversine
				neighborHexNode := graf.Nodes[neighborID]
				h_score := Haversine(neighborHexNode, goalHexNode)

				// kalkulasi f(n) = g(n) + h(n)
				f_score_baru := temp_gScore + h_score

				// Push tetangga beserta nilai f(n)nya ke antrian prioritas
				pqueue.Push(neighborID, f_score_baru)

				// counter node yang dihitung
				nodesGenerated++
			}
		}
	}

	// antrian habis => tidak ketemu goalnya
	return Result{NodesExpanded: nodesExpanded, NodesGenerated: nodesGenerated, Found: false}
}