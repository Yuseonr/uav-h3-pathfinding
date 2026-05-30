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