package usecase

import (
	"sync"

	"github.com/marcosvidolin/mdschallenge/database"
	"github.com/marcosvidolin/mdschallenge/model"
)

type ConsumeStock interface {
	Consume(country string, sku string, amount int) (model.Product, error)
}

type mdsConsumeStock struct {
	mtx sync.Mutex
	db  database.DB
}

func NewConsumeStock(d database.DB) ConsumeStock {
	return &mdsConsumeStock{
		mtx: sync.Mutex{},
		db:  d,
	}
}

// Consume consumes stock from a product
func (m *mdsConsumeStock) Consume(country string, sku string, amount int) (model.Product, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	prd, err := m.db.Get(country, sku)
	if err != nil {
		return model.Product{}, err
	}

	if amount > 0 {
		amount *= -1
	}

	_, err = prd.UpdateQtd(amount)
	if err != nil {
		return model.Product{}, err
	}

	err = m.db.Update(prd)
	if err != nil {
		return model.Product{}, err
	}

	return prd, nil
}
