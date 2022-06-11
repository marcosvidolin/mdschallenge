package database

import (
	"sync"

	"github.com/marcosvidolin/mdschallenge/model"
)

// In memory database
type inMemoryDB struct {
	data map[string]model.Product
}

// DB database interface
type DB interface {
	Get(country string, sku string) (model.Product, error)
	Create(p model.Product) error
	Update(p model.Product) error
}

type MdsDBClient struct {
	mtx sync.Mutex
	db  inMemoryDB
}

// Get gets a product from database
func (d *MdsDBClient) Get(country string, sku string) (model.Product, error) {
	d.mtx.Lock()
	defer d.mtx.Unlock()

	if d.db.data == nil {
		d.db.data = make(map[string]model.Product)
	}

	p, ok := d.db.data[country+sku]
	if !ok {
		return model.Product{}, &model.ProductNotFoundError{Country: country, Sku: sku}
	}
	return p, nil
}

// Create creates a product in datebase
func (d *MdsDBClient) Create(p model.Product) error {
	d.mtx.Lock()
	defer d.mtx.Unlock()

	if d.db.data == nil {
		d.db.data = make(map[string]model.Product)
	}
	id := p.Country + p.Sku
	d.db.data[id] = p
	return nil
}

// Update updates a product in database
func (d *MdsDBClient) Update(p model.Product) error {
	d.mtx.Lock()
	defer d.mtx.Unlock()

	if d.db.data == nil {
		d.db.data = make(map[string]model.Product)
	}
	id := p.Country + p.Sku
	_, ok := d.db.data[id]
	if !ok {
		return &model.ProductNotFoundError{Country: p.Country, Sku: p.Sku}
	}
	d.db.data[id] = p
	return nil
}
