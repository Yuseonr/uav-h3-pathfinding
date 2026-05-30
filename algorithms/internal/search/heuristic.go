package search

import (
	"math"
	"github.com/Yuseonr/uav-h3-pathfinding/algorithms/internal/graph"
)

/*
Yuseonr

Deskripsi	:
 - heuristic.go berisi fungsi fungsi yang dibutuh kan untuk menghitung heuristic value atau pun untuk menghitung jarak real
 = Haversine merupakan fungsi perhitungan jarak antar 2 titik koordinat pada bidang bulat
 - rumus mengkutip dari : K. Y. Chen, “An Improved A* Search Algorithm for Road Networks Using New Heuristic Estimation,” Jul. 2022, [Online]. Available: http://arxiv.org/abs/2208.00312
*/
func Haversine(n1, n2 *graph.Node) float64 {
	r := 6371000.00
	lat1, lng1 := n1.Lat*math.Pi/180, n1.Lng*math.Pi/180
	lat2, lng2 := n2.Lat*math.Pi/180, n2.Lng*math.Pi/180
	dlat, dlon := lat2-lat1, lng2-lng1
	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlon/2), 2)
	return 2 * r * math.Asin(math.Sqrt(a))
}


/*
Yuseonr

Deskripsi	:
 - HitungJarakRute menggunakan Haversine untuk mentotal jarak dari rute yang ditemukan dari start ke goal
 - menerima graph map nya, list of ID path nya mereturn jarak dalam meter
*/
func HitungJarakRute(g *graph.Graph, path []int32) float64 {
	var totalJarakMeter float64
	for i := 0; i < len(path)-1; i++ {
		totalJarakMeter += Haversine(g.Nodes[path[i]], g.Nodes[path[i+1]])
	}
	return totalJarakMeter
}