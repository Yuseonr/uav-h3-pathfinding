package search

import (
	h3 "github.com/uber/h3-go/v4"
	g "github.com/Yuseonr/uav-h3-pathfinding/algorithms/internal/graph"
)

/*
Yuseonr

Deskripsi   : 
 - Fungsi LineOfSight mengevaluasi apakah garis lurus antara dua node heksagon bebas dari rintangan
 - menggunakan library h3-go untuk menarik garis (GridPath) antar dua titik heksagon
 - jika ada heksagon yang dilewati garis tersebut namun tidak terdaftar di memori graf, maka LoS return false
 - else true bisa lewat
*/
func LineOfSight(graf *g.Graph, parentNode, neighborNode int32) bool {
	if parentNode == neighborNode {
		return true
	}

	h3ParentStr, _ := graf.Mapper.GetH3(parentNode)
	h3NeighborStr, _ := graf.Mapper.GetH3(neighborNode)
	h3ParentCell := h3.CellFromString(h3ParentStr)
	h3NeighborCell :=  h3.CellFromString(h3NeighborStr)

	intersectedCells, err := h3.GridPath(h3ParentCell, h3NeighborCell)
	if err != nil {
		return false
	}

	for _, cell := range intersectedCells {
		if _, exists := graf.Mapper.GetIDByCell(cell); !exists {
    		return false
		}
	}

	return true
}