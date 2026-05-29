package collections

import (
	"container/heap"
)

/*
Yuseonr

Deskripsi   : 
 - priority_queue.go implementasi struktur data Priority Queue menggunakan Min-Heap.
 - sumber referensi kode : https://medium.com/swlh/min-heaps-in-go-golang-fb3d9666c03f
*/
type PriorityQueue struct {
	mh *MinHeap
}

// CreatePriorityQueue membuat instance antrean baru yang aman digunakan
func CreatePriorityQueue() *PriorityQueue {
	mh := &MinHeap{}
	*mh = make(MinHeap, 0, 1024)
	heap.Init(mh)
	return &PriorityQueue{mh: mh}
}

// Push menambahkan ID Heksagon ke dalam antrean
func (pq *PriorityQueue) Push(nodeID int32, priority float64) {
	heap.Push(pq.mh, &HeapHex{NodeID: nodeID, Priority: priority})
}

// Pop mengambil elemen terkecil
func (pq *PriorityQueue) Pop() (int32, float64) {
	item := heap.Pop(pq.mh).(*HeapHex)
	return item.NodeID, item.Priority
}

// Peek melihat elemen terkecil tanpa mengeluarkannya
func (pq *PriorityQueue) Peek() (int32, float64, bool) {
	item := pq.mh.Peek()
	if item == nil {
		return 0, 0, false
	}
	hItem := item.(*HeapHex)
	return hItem.NodeID, hItem.Priority, true
}

// IsEmpty memeriksa apakah antrean kosong
func (pq *PriorityQueue) IsEmpty() bool {
	return pq.mh.Len() == 0
}