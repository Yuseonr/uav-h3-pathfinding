package graph

/*
Yuseonr

Deskripsi   : 
 - graph.go bertugas sebagai wadah yang menyimpan semua struct yang diperlukan untuk merepresentasikan graph.
 - Jadi nanti 1 object ini akan menyimpan keseluruhan map grpah yang sudah di load dari file json.
 - Graph menyimpan array pointer ke Node, dan IDMapper untuk menerjemahkan antara string H3 dan int32 ID.
*/
type Graph struct {
	Nodes  []*Node 	
	Mapper *IDMapper
}

/*
Yuseonr

Deskripsi   : 
 - Konstruktor membuat grpah baru dengan ukuran tertentu dan IDMapper yang sudah diinit.
 - menerima size untuk mengalokasikan array Nodes dengan kapasitas tetap, mencegah overhead relokasi memori saat menambahkan node baru.
 - menerima mapper untuk menyimpan referensi ke IDMapper yang sudah diinit.
 - mengembalikan pointer ke Graph yang sudah dibuat.
*/
func CreateGraph(size int32, mapper *IDMapper) *Graph {
	return &Graph{
		Nodes:  make([]*Node, size), 
		Mapper: mapper,
	}
}