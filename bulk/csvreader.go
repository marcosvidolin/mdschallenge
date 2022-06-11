package bulk

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/marcosvidolin/mdschallenge/model"
)

type BulkStockSource interface {
	Read(path string) ([]model.Product, error)
}

type csvSource struct {
	opts *Options
}

type Options struct {
	SkipHeader bool
}

func NewCsvSource(o *Options) BulkStockSource {
	return csvSource{
		opts: o,
	}
}

// Read read a file from a given path
func (c csvSource) Read(path string) ([]model.Product, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, &model.InternalServerError{Message: err.Error()}
	}
	defer f.Close()

	skipfl := false
	if c.opts != nil && c.opts.SkipHeader {
		skipfl = true
	}

	prods := []model.Product{}

	reader := csv.NewReader(f)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if skipfl {
			skipfl = false
			continue
		}

		var p model.Product
		p, err = recordToProd(record)
		if err != nil {
			return nil, &model.InternalServerError{Message: err.Error()}
		}

		prods = append(prods, p)
	}

	return prods, nil
}

func recordToProd(record []string) (model.Product, error) {
	p := model.Product{}
	p.Country = record[0]
	p.Sku = record[1]
	p.Name = record[2]
	var err error
	p.Qtd, err = strconv.Atoi(record[3])
	if err != nil {
		return model.Product{}, err
	}
	return p, nil
}
