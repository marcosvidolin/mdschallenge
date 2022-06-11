package model

import (
	"testing"
)

func TestUpdateQtd(t *testing.T) {

	var tests = []struct {
		prd     Product
		order   int
		want    int
		message string
	}{
		{Product{Sku: "AAAAAAAA", Name: "Fake Prod A", Qtd: 10}, 10, 20, "should add 10"},
		{Product{Sku: "BBBBBBBB", Name: "Fake Prod B", Qtd: 10}, -20, 10, "should return error and keep the original qtd"},
		{Product{Sku: "CCCCCCCC", Name: "Fake Prod C", Qtd: 10}, -5, 5, "should subtract -5"},
	}

	for _, tt := range tests {
		t.Run(tt.message, func(t *testing.T) {
			ans, _ := tt.prd.UpdateQtd(tt.order)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

}
