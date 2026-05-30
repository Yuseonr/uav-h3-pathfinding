/*
Yuseonr

Deskripsi   : 
 - mapper.go bertugas sebagai penerjemah dari data h3_res{x}_graph.json ke dalam bentuk struct yang telah dibuat.
 - untuk memenuhi tugas itu dibuat struct IDMapper yang menangani terjemahan dua arah antara string H3 (64bit) dan internal int32.
*/

package graph


/*
Yuseonr

Deskripsi   : 
 - IDMapper adalah struct yang menyimpan dua map untuk menerjemahkan antara string H3 dan int32 ID.
 - h3ToID menyimpan mapping dari string H3 ke int32 ID.
 - idToH3 menyimpan mapping dari int32 ID ke string H3.
 - nextID adalah counter untuk memberikan ID baru saat H3 baru ditemukan.
*/
type IDMapper struct {
	h3ToID map[string]int32 // mapping dari string H3 ke int32 ID
	idToH3 map[int32]string // mapping dari int32 ID ke string H3
	nextID int32			// counter untuk ID berikutnya	
}

/*
Yuseonr

Deskripsi   : 
 - Konstruktor untuk init IDMapper dengan map kosong dan nextID dimulai dari 0.
*/
func CreateIDMapper() *IDMapper {
	return &IDMapper{
		h3ToID: make(map[string]int32),
		idToH3: make(map[int32]string),
		nextID: 0, // ximulai dari indeks 0
	}
}

/*
Yuseonr

Deskripsi   : 
 - AddOrGetID mengecek apakah H3 sudah terdaftar. Jika belum, buat ID(int32) baru + kembalikan ID(int32)nya
 - jika sudah, kembalikan ID(int32) yang sudah ada.
*/
func (idm *IDMapper) AddOrGetID(idh3 string) int32 {
	// Cek apakah H3 ini sudah memiliki ID(int32) yang terdaftar
	if id, exists := idm.h3ToID[idh3]; exists {
		// sudah ada, kembalikan ID(int32) yang sudah terdaftar
		return id
	}
	
	// H3 baru, buat ID baru dan simpan di kedua map
	newid := idm.nextID
	idm.h3ToID[idh3] = newid
	idm.idToH3[newid] = idh3
	// incr untuk ID berikutnya
	idm.nextID++
	
	return newid
}

/*
Yuseonr

Deskripsi   : 
 - GetID mengecek apakah H3 sudah terdaftar kembalikan ID(int32)nya jika belum retunr 0
*/
func (idm *IDMapper) GetID(idh3 string) (int32, bool) {
	id, exists := idm.h3ToID[idh3]
	return  id, exists
}

/*
Yuseonr

Deskripsi   : 
 - GetH3 menerima ID(int32) dan mengembalikan string H3 / IDH3 asli yang sesuai jika ada ata false jika tidak ditemukan
*/
func (idm *IDMapper) GetH3(id int32) (string, bool) {
	h3, exists := idm.idToH3[id]
	return h3, exists
}