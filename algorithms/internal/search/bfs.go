package search

import (
	c "github.com/Yuseonr/uav-h3-pathfinding/algorithms/internal/collections"
	g "github.com/Yuseonr/uav-h3-pathfinding/algorithms/internal/graph"
)

/*
Yuseonr

Deskripsi   :
 - bfs.go implementasi algoritma Breadth-First Search (BFS) untuk pencarian jalur pada graf
 - BFS menjelajahi graf level per level menggunakan FIFO Queue.
 - Menjamin jalur terpendek dalam satuan jumlah loncatan.
*/
type BFS struct{}

func (b *BFS) SearchPath(graf *g.Graph, start, goal int32) Result {

	// inisialisasi queue, asalDari, dan counter
	queue := c.CreateQueue()

	// Penjelasan asalDari :
	// untuk minimalisir overhead dari pengaksesan dan pengubahan hash map, digunakan array/slice of int32 yang dimana ini akan memenuhi 2 tugas yaitu :
	// 1. sebagai penanda apakah suatu node sudah dikunjungi
	// 2. sebagai penanda jika dikunjungi, dikunjungi dari node mana
	// hal ini dapat tercapai dengan cara menggunakan index serta value dari slice tersebut:
    // - indeks (Index) = representasi ID node/heksagon yang sedang diperiksa
    // - nilai (Value) = representasi ID node/heksagon sebelumnya (parent) yang membawanya ke sini
	// Mengunakan nilai awal = -1 / sentinel nya untuk menandai node belum di lihat siapa siapa
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

	// enqueue heksagon awal ke dalam queue
	queue.Enqueue(start)

	// selama antrian belum kosong terus cari
	for !queue.IsEmpty() {
		// ambil hexagon paling depan
		currentHex, _ := queue.Dequeue()
		// counter node yang di expand / dikeluarkan dari queue
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

		// curren bukan goal, jadi ambil node / hexagon aslinya terus periksa tetangganya di for loop bawah
		currentHexNode := graf.Nodes[currentHex]

		// untuk setiap tetangga, jika belum dikunjungi (asalDari == -1), tandai asalDarinya dan enqueue ke dalam queue
		// sesuai urutan tetangga yang ada di graf
		for _, neighborID := range currentHexNode.Neighbors {
			if asalDari[neighborID] == -1 {
				asalDari[neighborID] = currentHex 
				queue.Enqueue(neighborID)
				// node yang dilihat / dimasukan queue
				nodesGenerated++
			}
		}
	}
	// tidak ketemu goal nya
	return Result{ NodesExpanded: nodesExpanded, NodesGenerated: nodesGenerated, Found: false}
}