package graph

/*
Yuseonr

Deskripsi   :
Struct yang akan merepresentasikan node pada graph, ini merupakan satu heksagon H3 di dalam memori.

Optimisasi:
 - ID disimpan dalam bentuk integer agar lebih efisien dibanding string.
 - urutan attribute di optimalisasi untuk meminimalisir padding, sehingga ukuran struct lebih kecil
 - semua attribut diset public
*/
type Node struct {
	Lat       float64
	Lng       float64
	ID        int32
	Cost      int32
	Neighbors []int32
}