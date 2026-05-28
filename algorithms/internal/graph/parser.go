/*
Yuseonr

Deskripsi   : 
 - parser.go bertugas untuk memuat file JSON yang berisi data graf dan mengubahnya menjadi struktur data yang dapat digunakan dalam memori.
*/

package graph

import (
	"encoding/json"
	"os"
)

/*
Yuseonr

Deskripsi   : 
 - representasi json node mentahnya
 - nama/id h3 tidak masuk karena pada json aslinya nama h3 adalah key dari object
 - : example :
 {
  "898d8c7913bffff": {
    "lat": -7.058443335807414,
    "lng": 110.46785224384097,
    "cost": 1,
    "neighbors": [
      "898d8c7910fffff",
      "898d8c79177ffff",
      "898d8c7912bffff",
      "898d8c79123ffff",
      "898d8c79133ffff",
      "898d8c79107ffff"
    ]
  },
  "898d8c6b503ffff": {
    "lat": -7.021706655291291,
    "lng": 110.40757255620079,
    "cost": 1,
    "neighbors": [
      "898d8c6b51bffff",
      "898d8c6b50bffff",
      "898d8c6b50fffff",
      "898d8c6b513ffff"
    ]
  },
  ...
 }
*/
type rawJsonNode struct {
	Lat       float64  `json:"lat"`
	Lng       float64  `json:"lng"`
	Cost      int32    `json:"cost"`
	Neighbors []string `json:"neighbors"` 
}


/*
Yuseonr

Deskripsi   : 
 - Menerima filepath ke json graf -> mereturn Graph yang sudah di load ke memori
*/
func LoadJsonGraph(jsonfilepath string) (*Graph, error) {
	// Buka file JSON
	graf_json, err := os.Open(jsonfilepath)
	if err != nil {
		return nil, err
	}
	defer graf_json.Close()

	// Map untuk menyimpan data mentah dari JSON
	var rawMap map[string]rawJsonNode
	if err := json.NewDecoder(graf_json).Decode(&rawMap); err != nil {
		return nil, err
	}

	// Buat IDMapper dan ngemap satu satu H3 ke ID(int32) untuk semua node yang ada di rawMap
	mapper := CreateIDMapper()
	for h3 := range rawMap {
		mapper.AddOrGetID(h3)
	}

	// init graph dengan ukuran sesuai jumlah node yang sudah di map, dan IDMapper yang sudah diinit
	graf_memori := CreateGraph(mapper.nextID, mapper)

	// Iterasi lagi untuk mengisi data node ke dalam graph berdasarkan mapping yang sudah dibuat
	for h3, raw := range rawMap {
		// ini harusnya sudah pasti ada karena kita sudah map semua H3 sebelumhya
		id := mapper.AddOrGetID(h3)

		// node yang akan dimasukkan ke dalam graph tapi lum masukin neightboornya
		node := &Node{
			ID:        id,
			Lat:       raw.Lat,
			Lng:       raw.Lng,
			Cost:      raw.Cost,
			Neighbors: make([]int32, 0, len(raw.Neighbors)), 
		}

		// untuk setiap neighbor H3, cek apakah sudah ada mapping ID(int32) dan tambahkan ke list neighbors node ini
		for _, neighborH3 := range raw.Neighbors {
			if neighborID, exists := mapper.h3ToID[neighborH3]; exists {
				node.Neighbors = append(node.Neighbors, neighborID)
			}
		}
		// simpan node ini ke dalam graph di index yang sesuai dengan ID(int32)nya
		graf_memori.Nodes[id] = node
	}

	return graf_memori, nil
}