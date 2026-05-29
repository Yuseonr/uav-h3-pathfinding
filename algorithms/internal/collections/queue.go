package collections

/*
Yuseonr

Deskripsi   : 
 - queue.go implementasi struktur data queue ketika penelusuran graf
 - FIFO
*/
type Queue struct {
	hexes []int32
}

// Yuseonr
//
// membuat queue baru array dinamis dengan kapasitas awal 1024
func CreateQueue() *Queue {
	return &Queue{ hexes: make([]int32, 0 , 1024) }
}

// Yuseonr
//
// menambahkan elemen ke belakang queue
func (q *Queue) Enqueue(hex int32) {
	q.hexes = append(q.hexes, hex)
}

// Yuseonr
//
// menghapus elemen dari depan queue dan mengembalikan nilainya, jika kosong mengembalikan false
func (q *Queue) Dequeue() (int32, bool) {
	if len(q.hexes) == 0 {
		return 0, false
	}
	hex := q.hexes[0]
	q.hexes = q.hexes[1:]
	return hex, true
}

// Yuseonr
//
// isempty
func (q *Queue) IsEmpty() bool {
	return len(q.hexes) == 0
}