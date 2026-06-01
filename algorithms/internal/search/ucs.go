package search

import (
    "math"
    c "github.com/Yuseonr/uav-h3-pathfinding/algorithms/internal/collections"
    g "github.com/Yuseonr/uav-h3-pathfinding/algorithms/internal/graph"
)

/*
Yuseonr

Deskripsi   :
 - ucs.go implementasi algoritma Uniform Cost Search (UCS) untuk pencarian jalur pada graf.
 - UCS menjelajahi graf dengan memprioritaskan ekspansi pada node yang memiliki total biaya (cost) terendah dari titik awal.
 - Menggunakan Min-Priority Queue untuk mengurutkan antrian berdasarkan biaya.
 - Menjamin rute terpendek secara geometris dengan menggunakan konstanta JarakPerStep sebagai pengganti fungsi Haversine agar komputasi O(1).
*/
type UCS struct{
    JarakPerStep float64
}

func (u *UCS) SearchPath(graf *g.Graph, start, goal int32) Result {

    // inisialisasi min-priority queue -> selalu memberikan node dengan akumulasi cost terkecil
    pqueue := c.CreatePriorityQueue()

    // Penjelasan costTerbaik dan asalDari :
    // 1. costTerbaik: mencatat gscore, menyimpan nilai biaya termurah dari start ke node index.
    // 2. asalDari: bisa cek bfs.go
    costTerbaik := make([]float64, len(graf.Nodes))
    asalDari    := make([]int32, len(graf.Nodes))
    
    // inisialisasi semua cost dengan nilai Tak Terhingga (Infinity) dan asalDari dengan -1 
    for i := range costTerbaik {
        costTerbaik[i] = math.Inf(1) 
        asalDari[i]    = -1
    }

    // counter
    nodesExpanded  := 0
    nodesGenerated := 0

    // mulai UCS
    // tandai hexagon awal udh dikunjungi dan cost di titik awal adalah 0
    asalDari[start]    = start
    costTerbaik[start] = 0
    
    
    // push heksagon awal ke antrian prioritas dengan prioritas cost 0
    pqueue.Push(start, 0)

    // selama antrian prioritas belum kosong terus cari
    for !pqueue.IsEmpty() {
        // ambil hexagon dengan cost prioritas terendah saat ini
        currentHex, currentCost := pqueue.Pop()

        // Skip entri lama yang sudah tidak merepresentasikan cost terbaik.
        // karena priority queue tidak mendukung update prioritas, entri lama
        // tetap dibiarkan di antrian tapi langsug si skip saat dipop.
        if currentCost > costTerbaik[currentHex] {
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

        // current bukan goal, jadi ambil node / hexagon aslinya terus periksa tetangganya di for loop bawah
        currentHexNode := graf.Nodes[currentHex]

        // untuk setiap tetangga, evaluasi apakah melewati tetangga dari currentHex ini lebih murah
        for _, neighborID := range currentHexNode.Neighbors {
            // hitung biaya baru: biaya ke current + biaya langkah ke tetangga pake const resolusi
            newCost := currentCost + u.JarakPerStep
            
            // jika biaya baru lebih kecil dari rekor biaya terbaik tetangga tersebut sebelumnya
            if newCost < costTerbaik[neighborID] {

                // perbarui rekor cost terbaik tetangga
                costTerbaik[neighborID] = newCost

                // catat bahwa rute termurah ke tetangga ini berasal dari currentHex
                asalDari[neighborID] = currentHex

                // push hexid tetangga beserta cost baru ke antrian prioritas
                pqueue.Push(neighborID, newCost)

                // node yang generate / dimasukkan ke antrian
                nodesGenerated++
            }
        }
    }
    // tidak ketemu goal nya
    return Result{ NodesExpanded:  nodesExpanded, NodesGenerated: nodesGenerated, Found: false }
}