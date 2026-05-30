package search

import (
	c "github.com/Yuseonr/uav-h3-pathfinding/algorithms/internal/collections"
	g "github.com/Yuseonr/uav-h3-pathfinding/algorithms/internal/graph"
)

/*
Yuseonr

Deskripsi   :
 - dfs.go implementasi algoritma Depth-First Search (DFS) untuk pencarian jalur pada graf
 - DFS menjelajahi graf dengan cara mencari secara mendalam pada satu node dulu
 - Tidak menjamin rute terpendek
*/
type DFS struct{}

func (d *DFS) SearchPath(graf *g.Graph, start, goal int32) Result {
	
	// inisialisasi stack, asalDari, dan counter
	stack := c.CreateStack()

	// Penjelasan asalDari bisa cek : bfs.go
	asalDari := make([]int32, len(graf.Nodes))
	nodesExpanded := 0
	nodesGenerated := 0

	// inisialisasi asalDari dengan -1 untuk menandakan belum dikunjungi
	for i := range asalDari {
		asalDari[i] = -1
	}

	// mulai BFS
	// tandai heksagon awal sudah dikunjungi
	asalDari[start] = start

	// push heksagon awal ke dalam stack
	stack.Push(start)

	// selama stack belum kosong terus cari
	for !stack.IsEmpty() {
		// ambil hexagon paling depan
		currentHex ,_ := stack.Pop()
		// counter node yang di expand / di pop dari stack
		nodesExpanded++

		// // jika ini adalah goal, return hasilnya dengan merekonstruksi jalur dari asalDari
		if currentHex == goal {
			path := reconstructPath(asalDari, start, goal)
			return Result {
				PathIDs: path,
				NodesExpanded: nodesExpanded,
				NodesGenerated: nodesGenerated,
				Found: true,
			}
		}

		// curren bukan goal, jadi ambil node / hexagon aslinya terus periksa tetangganya di for loop bawah
		currentHexNode := graf.Nodes[currentHex]

		// untuk setiap tetangga, jika belum dikunjungi (asalDari == -1), tandai asalDarinya dan psuh ke stack
		// sesuai urutan tetangga yang ada di graf
		for _, neighborID := range currentHexNode.Neighbors {
			if asalDari[neighborID] == -1 {
				asalDari[neighborID] = currentHex
				stack.Push(neighborID)
				// node yang dilihat / dipush ke stack
				nodesGenerated++
			}

		}
	}
	// tidak ketemu goalnya
	return Result{ NodesExpanded: nodesExpanded, NodesGenerated: nodesGenerated, Found: false}
}