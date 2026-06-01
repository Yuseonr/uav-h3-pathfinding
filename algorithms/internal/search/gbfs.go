package search

import (
	c "github.com/Yuseonr/uav-h3-pathfinding/algorithms/internal/collections"
	g "github.com/Yuseonr/uav-h3-pathfinding/algorithms/internal/graph"
)

/*
Yuseonr

Deskripsi   :
 - gbfs.go implementasi algoritma greedy local search untuk pencarian jalur pada graf
 - Greedy Best First Search memilih jalan secara greedy yaitu :
 - dia akan memilih neighbor dengan jarak haversine terdekat dengan goal
 - berbeda dengan greedy biasa GBFS memiliki fallback priority queue sehingga tidak akan terjebak pada local minimu
 - menjamin menemukan rute
 - tidak menjamin rute terpendek
*/
type GBFS struct{}

func (gbfs *GBFS) SearchPath(graf *g.Graph, start, goal int32) Result {

	// inisialisasi min priority queue -> selalu memberikan node dengan akumulasi cost terkecil
	pqueue := c.CreatePriorityQueue()

	// inisilaisasi asalDari : penjelasan ada di bfs.go
	asalDari := make([]int32, len(graf.Nodes))
	nodesExpanded := 0
	nodesGenerated := 0

	// inisialisasi semua cost dengan nilai Tak Terhingga (Infinity) dan asalDari dengan -1
	for i := range asalDari {
		asalDari[i] = -1
	}

	// mulai gbfs
	// tandai hexagon awal udh dikunjungi
	asalDari[start] = start

	// mengambil node goal sekali untuk evaluasi haversine
	goalHexNode := graf.Nodes[goal]

	// push start dengan nilai 0
	pqueue.Push(start, 0)

	// selama antrian belum kosong terus cari
	for !pqueue.IsEmpty() {
		// ambil hex sekarang (tidak perlu value nya karna secara otomatis yang kita pop adalah jarak terdekat karena priority queue)
		currentHex, _ := pqueue.Pop()
		// node yang dipilih
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

		// current bukan goal, jadi ambil node / hexagon aslinya terus periksa tetangganya di for loop bawah
		currentHexNode := graf.Nodes[currentHex]

		for _, neighborID := range currentHexNode.Neighbors {
			// Apabila node belum pernah dievluasi
			if asalDari[neighborID] == -1 {
				// set kita menemukan neighbor ini dari mana
				asalDari[neighborID] = currentHex
				// ambil node nya
				neighborHexNode := graf.Nodes[neighborID]
				// hitung cost f(n) = g(n)
				neighborCost := Haversine(neighborHexNode, goalHexNode)
				// terhitung sebagai node yang kita masukan ke stack
				nodesGenerated++
				// push dan otomatis akan masuk ke dalam priority queue
				pqueue.Push(neighborID, neighborCost)
			}

		}
	}
	// queue habis dan belum ketemu goal => tidak ketemu path
	return Result{ NodesExpanded: nodesExpanded, NodesGenerated: nodesGenerated, Found: false}
}
