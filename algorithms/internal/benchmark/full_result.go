package benchmark

import "fmt"

/*
Yuseonr

Deskripsi   : 
 - full_result.go menyimpan struct FullResult yang digunakan untuk menyimpan hasil lengkap benchmark algoritma pathfinding dalam 1 test case
 - dari struct ini bisa langsung expor ke .csv untuk analisa

*/
type FullResult struct {

	IdTestCase				int		// IdTestCase untuk saat ini generated dari main exec, belum melekat pada json test case nya
	Algoritma				string	// nama algoritma yang digunakan untuk tc ini 
	ResolusiMap 			int	  	// resolusi map untuk tc ini
	KategoriJarak 			string	// kategori jarak

	StartHexID          	string  
	GoalHexID             	string  

	BanyakLoncatan          int	   	// len(jalurhexids)-1 -> banyak loncatan hexagon menuju goal
	TotalJarakMeter	       	float64 // total jarak fisik antar titik titik kordinat hexagon (pusat 1 ke pusat 2 hingga goal)
	EuclideanJarakMeter		float64 // jarak murni start ke goal ditarik lurus tanpa peduli obstacles

	WaktuEksekusiMicrosec 	int64	// waktu eksekusi
	NodesExpanded   		int     // banyak node yang diexpand
	NodesGenerated  		int     // banyak node yang dilihat

	Status					string	 // status 
	JalurHexIDs				[]string // Jalur hex id start ke goal 
}

// Fungsi helper untuk mengubah struct menjadi baris CSV (array of string)
func (r *FullResult) ToCSVRow() []string {
	return []string{
		fmt.Sprintf("%d", r.IdTestCase),
		r.Algoritma,
		fmt.Sprintf("%d", r.ResolusiMap),
		r.KategoriJarak,
		r.StartHexID,
		r.GoalHexID,
		fmt.Sprintf("%d", r.BanyakLoncatan),
		fmt.Sprintf("%.4f", r.TotalJarakMeter),
		fmt.Sprintf("%.4f", r.EuclideanJarakMeter),
		fmt.Sprintf("%d", r.WaktuEksekusiMicrosec),
		fmt.Sprintf("%d", r.NodesExpanded),
		fmt.Sprintf("%d", r.NodesGenerated),
		r.Status,
		fmt.Sprintf("%v", r.JalurHexIDs),
	}
}

// Fungsi helper untuk Header CSV
func GetCSVHeader() []string {
	return []string{
		"id_test_case", "algoritma", "resolusi", "kategori_jarak",
		"start_hex", "goal_hex", "hops", "jarak_meter", "referensi_meter",
		"waktu_microsec", "nodes_expanded", "nodes_generated", "status", "path_hex_ids",
	}
}