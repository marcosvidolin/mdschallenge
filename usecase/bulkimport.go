package usecase

import (
	"log"
	"sync"

	"github.com/marcosvidolin/mdschallenge/bulk"
	"github.com/marcosvidolin/mdschallenge/database"
	"github.com/marcosvidolin/mdschallenge/model"
	"github.com/marcosvidolin/mdschallenge/util"
)

type BulkOperation interface {
	Import(path string) error
}

type mdsBulkOperation struct {
	mtx     sync.Mutex
	db      database.DB
	blk     bulk.BulkStockSource
	maxProc int
}

func NewBulkOperation(d database.DB, b bulk.BulkStockSource) BulkOperation {
	return &mdsBulkOperation{
		mtx:     sync.Mutex{},
		db:      d,
		blk:     b,
		maxProc: 5,
	}
}

// BulkImport update a list of product
func (s *mdsBulkOperation) Import(path string) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	pdrs, err := s.blk.Read(path)
	if err != nil {
		return &model.InternalServerError{Message: err.Error()}
	}

	// calculates the chunk size according to the max num of process
	csize := int(len(pdrs) / s.maxProc)
	pchunks := util.ChunkSlice(pdrs, csize)

	wg := sync.WaitGroup{}

	// process the items using chunks for better performance
	for _, pc := range pchunks {
		wg.Add(1)
		go func(t *mdsBulkOperation, pdrs []model.Product) {
			defer wg.Done()
			for _, p := range pdrs {
				err := s.upinsert(p)
				if err != nil {
					log.Printf("upinsert error: %v", err)
					continue
				}
			}
		}(s, pc)
	}

	// wg.Wait()

	return nil
}

func (s *mdsBulkOperation) upinsert(p model.Product) error {
	prod, _ := s.db.Get(p.Country, p.Sku)

	// check if the product exist, if not... must be created
	if len(prod.Sku) == 0 {
		err := s.db.Create(p)
		if err != nil {
			return err
		}
		return nil
	}

	prod.Qtd += p.Qtd

	err := s.db.Update(prod)
	if err != nil {
		return err
	}

	return nil
}
