package collections

/*
Yuseonr

Deskripsi   : 
 - stack.go implementasi struktur data stack ketika penelusuran graf
 - LIFO
*/
type Stack struct {
	hexes []int32
}

// Yuseonr
//
// membuat stack baru array dinamis dengan kapasitas awal 1024
func CreateStack() *Stack {
	return &Stack{ hexes: make([]int32, 0 , 1024) }
}

// Yuseonr
//
// menambahkan elemen ke atas stack
func (s *Stack) Push(hex int32) {
	s.hexes = append(s.hexes, hex)
}


// Yuseonr
//
// menghapus elemen dari atas stack dan mengembalikan nilainya, jika kosong mengembalikan false
func (s *Stack) Pop() (int32, bool) {
	if len(s.hexes) == 0 {
		return 0, false
	}
	hex := s.hexes[len(s.hexes)-1]
	s.hexes = s.hexes[:len(s.hexes)-1]
	return hex, true
}


// Yuseonr
//
// isempty
func (s *Stack) IsEmpty() bool {
	return len(s.hexes) == 0
}