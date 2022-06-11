package util

import (
	"testing"

	"github.com/marcosvidolin/mdschallenge/model"
)

func TestChunkSlice(t *testing.T) {
	prods := []model.Product{
		{Sku: "1"},
		{Sku: "2"},
		{Sku: "3"},
		{Sku: "4"},
		{Sku: "5"},
	}

	if len(ChunkSlice(prods, 2)) != 3 {
		t.Errorf("chunk len chould be 3")
	}

	if len(ChunkSlice(prods, 1)) != 5 {
		t.Errorf("chunk len chould be 5")
	}

	if len(ChunkSlice([]model.Product{}, 1)) != 0 {
		t.Errorf("chunk len chould be 0")
	}

}
