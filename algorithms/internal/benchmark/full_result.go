package benchmark

/*
Yuseonr

Deskripsi   : 
 - full_result.go menyimpan struct FullResult yang digunakan untuk menyimpan hasil lengkap benchmark algoritma pathfinding dalam 1 test case
 - dari struct ini bisa langsung expor ke .csv untuk analisa

*/
type FullResult struct {

	IdTestCase				string	// IdTestCase untuk saat ini generated dari main exec, belum melekat pada json test case nya
	Algoritma				string	// nama algoritma yang digunakan untuk tc ini 
	ResolusiMap 			int	  	// resolusi map untuk tc ini
	KategoriJarak 			string	// kategori jarak

	BanyakLoncatan          int	   	// len(jalurhexids)-1 -> banyak loncatan hexagon menuju goal
	TotalJarakMeter	       	float64 // total jarak fisik antar titik titik kordinat hexagon (pusat 1 ke pusat 2 hingga goal)
	EuclideanJarakMeter		float64 // jarak murni start ke goal ditarik lurus tanpa peduli obstacles

	WaktuEksekusiMicrosec 	int64	// waktu eksekusi
	NodesExpanded   		int     // banyak node yang diexpand
	NodesGenerated  		int     // banyak node yang dilihat

	Status					string	 // status 
	JalurHexIDs				[]string // Jalur hex id start ke goal 
}