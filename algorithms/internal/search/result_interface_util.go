package search

import "github.com/Yuseonr/uav-h3-pathfinding/algorithms/internal/graph"

/*
Yuseonr

Deskripsi   : 
 - Result adalah hasil dari eksekusi algoritma
*/
type Result struct {
	PathIDs 		[]int32 // rute yang ditemukan dalam bentuk ID int32
	NodesExpanded   int     // jumlah heksagon yang dikeluarkan dari antrean (Pop)
	NodesGenerated  int     // jumlah heksagon tetangga yang dievaluasi (Push)
	Found           bool    // status apakah rute berhasil ditemukan (harusnya ditemukan semua untuk testcase yang valid)
}

/*
Yuseonr

Deskripsi   :
 - Algorithm interface agar semua algoritma nya dapat mengimplementasikan ini.
*/
type Algorithm interface {
	SearchPath(graf_map *graph.Graph, startID, goalID int32) Result
}

/*
Yuseonr

Deskripsi : 
 - util procedure untuk merekonstruksi jalur dari array asalDari.
*/
func reconstructPath(asalDari []int32, start, goal int32) []int32 {

	// ngetrace back dari goal ke start menggunakan array asalDari untuk menemukan jalur yang dilalui
	// harus gini karena cara kita menyimpan path (dijelaskan di bfs.go) 
	// jadi untuk rekonstruksi bakal dari akhir -> dia di expand dari node mana -> ke node itu -> dia diexpand dari node mana ....
	// hingga sampai ke start kemudian di reverse
	var jalur []int32
	langkahSaatIni := goal

	for langkahSaatIni != start {
		jalur = append(jalur, langkahSaatIni)
		langkahSebelumnya := asalDari[langkahSaatIni]
		langkahSaatIni = langkahSebelumnya
	}
	jalur = append(jalur, start)

	// reverse jalur nya
	indeksKiri := 0
	indeksKanan := len(jalur) - 1

	for indeksKiri < indeksKanan {
		nilaiSementara := jalur[indeksKiri]
		jalur[indeksKiri] = jalur[indeksKanan]
		jalur[indeksKanan] = nilaiSementara
		indeksKiri++
		indeksKanan--
	}

	return jalur
}