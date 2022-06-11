package usecase

import (
	"sync"

	"github.com/marcosvidolin/mdschallenge/database"
	"github.com/marcosvidolin/mdschallenge/model"
)

type FindProduct interface {
	FindByCountryAndSku(country string, sku string) (model.Product, error)
}

type mdsFindProduct struct {
	mtx sync.Mutex
	db  database.DB
}

func NewFindProduct(d database.DB) FindProduct {
	return &mdsFindProduct{
		mtx: sync.Mutex{},
		db:  d,
	}
}

// FindById gets a product by id
func (m *mdsFindProduct) FindByCountryAndSku(country string, sku string) (model.Product, error) {
	p, err := m.db.Get(country, sku)
	if err != nil {
		return model.Product{}, err
	}
	return p, nil
}
