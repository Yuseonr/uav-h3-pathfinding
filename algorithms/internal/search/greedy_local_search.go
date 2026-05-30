package search

import (
	"math"

	g "github.com/Yuseonr/uav-h3-pathfinding/algorithms/internal/graph"
)

/*
Yuseonr

Deskripsi   :
 - greedy_local_search.go implementasi algoritma greedy local search untuk pencarian jalur pada graf
 - greedy_local_search atau yang digunakan : Steepest-Descent Hill Climbing merupakan salah satu algoritma local search yang memilih jalan secara greedy yaitu :
 - dia akan memilih neighbor dengan jarak haversine terdekat dengan goal tanpa mempedulikan neighbor lain
 - tidak menjamin rute terpendek
 - source : Henri Tantyoko, Sandy Kurniawan, Helmie Arif Wibawa, "Kecerdasan Buatan / Sistem Cerdas 06: Local Search", S1 Informatika UNDIP.
*/
type GLS struct{}

func (gls *GLS) SearchPath(graf *g.Graph, start, goal int32) Result {

	// inisialisasi asalDari, dan counter
	// Penjelasan asalDari bisa cek : bfs.go
	asalDari := make([]int32, len(graf.Nodes))
	nodesExpanded := 0
	nodesGenerated := 0

	// inisialisasi asalDari dengan -1 untuk menandakan belum dikunjungi
	for i := range asalDari {
		asalDari[i] = -1
	}

	// tandai heksagon awal sudah dikunjungi
	asalDari[start] = start

	// inisialisasi currentHex pengganti antrian / stack
	currentHex := start

	// untuk perhitungan haversine
	goalHexNode := graf.Nodes[goal]

	// selama hex saat ini bukanlah goal
	for currentHex != goal {
		// counter node yang di expand / dievaluasi
		nodesExpanded++ 

		// inisialisasi untuk cari best step nya
		bestNeighbor := int32(-1)
		minJarak := math.Inf(1)
		
		// ambil node dari currenthex
		currentHexNode := graf.Nodes[currentHex]
		
		// evaluasi tiap tetangga / neighbor yang belum pernah dikunjungi
		for _, neighborID := range currentHexNode.Neighbors {
	
			// belum perna dikunjungi = -1
			if asalDari[neighborID] == -1 {
				// ambil node tetangga
				neighborHexNode := graf.Nodes[neighborID]
				// evaluasi jaraknya ke goal
				jarakNeighborToGoal := Haversine(neighborHexNode, goalHexNode)
				// nodes generated disini merujuk ke node yang kita hitung
				nodesGenerated++ 
	
				// cari jarak yang paling kecil (Steepest-Descent)
				if (jarakNeighborToGoal < minJarak) {
					minJarak = jarakNeighborToGoal
					bestNeighbor = neighborID
				}
			}
		}

		// local minimum problem tidak ada lagi tetangga yang bisa di evaluasi (udah dikunjungi semua)
		// tidak ketemu goalny
		if bestNeighbor == -1 {
			return Result{ NodesExpanded: nodesExpanded, NodesGenerated: nodesGenerated, Found: false}
		}

		// best neighgboor
		asalDari[bestNeighbor] = currentHex
		currentHex = bestNeighbor
	}

	// jika keluar dari loop, artinya currentHex == goal 
	// +1 ke nodesExpanded untuk node goal 
	nodesExpanded++
	
	path := reconstructPath(asalDari, start, goal)
	
	return Result{
		PathIDs:        path,
		NodesExpanded:  nodesExpanded,
		NodesGenerated: nodesGenerated,
		Found:          true,
	}
}