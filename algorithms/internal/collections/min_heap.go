package collections

/*
Yuseonr

Deskripsi   : 
 - min_heap.go implementasi struktur data Min-Heap untuk digunakan dalam PriorityQueue.
 - Min-Heap adalah struktur data yang memungkinkan kita untuk selalu mengakses elemen dengan prioritas tertinggi (nilai terkecil) dalam waktu O(1) dan melakukan penambahan atau penghapusan elemen dalam waktu O(log N).
 - sumber referensi kode : https://medium.com/swlh/min-heaps-in-go-golang-fb3d9666c03f
*/
// HeapHex merepresentasikan satu id heksagon beserta priority nya (bisa dari harga g(n) atau f(n))
type HeapHex struct {
	NodeID   int32
	Priority float64 
}

// List of pointer to HeapHex
type MinHeap []*HeapHex

// Implementasi interface heap.Interface untuk MinHeap
func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Priority < h[j].Priority }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i]}

// Push menambahkan elemen baru ke heap, di sini kita menambahkan pointer ke HeapHex
func (h *MinHeap) Push(x interface{}) { 
	*h = append(*h, x.(*HeapHex))
}

// Pop menghapus dan mengembalikan elemen dengan prioritas terkecil dari heap
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Peek mengembalikan elemen dengan prioritas terkecil tanpa menghapusnya dari heap
func (h *MinHeap) Peek() interface{} {
	old := *h
	if len(old) == 0 {
		return nil
	}
	x := old[0] 
	return x
}

