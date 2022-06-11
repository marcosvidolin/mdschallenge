package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcosvidolin/mdschallenge/model"
	"github.com/marcosvidolin/mdschallenge/usecase"
)

type Consume struct {
	Amount int `json:"amount"`
}

type RestController interface {
	FindByCountryAndSku(c *gin.Context)
	ConsumeStock(c *gin.Context)
	BulkImport(c *gin.Context)
}

type mdsController struct {
	cns usecase.ConsumeStock
	fnd usecase.FindProduct
	blk usecase.BulkOperation
}

func New(c usecase.ConsumeStock, f usecase.FindProduct, b usecase.BulkOperation) RestController {
	return &mdsController{
		cns: c,
		fnd: f,
		blk: b,
	}
}

// FindProductBySku http handler for search for products
func (m *mdsController) FindByCountryAndSku(c *gin.Context) {
	country := c.Param("country")
	if len(country) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "country required"})
		return
	}

	sku := c.Param("sku")
	if len(sku) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sku required"})
		return
	}

	p, err := m.fnd.FindByCountryAndSku(country, sku)
	if err != nil {
		c.JSON(mapError(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)
}

// ConsumeStock http handler for stock consuming
func (m *mdsController) ConsumeStock(c *gin.Context) {

	country := c.Param("country")
	if len(country) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "country required"})
		return
	}

	sku := c.Param("sku")
	if len(sku) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sku required"})
		return
	}

	body := Consume{}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	prd, err := m.cns.Consume(country, sku, body.Amount)
	if err != nil {
		c.JSON(mapError(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, prd)
}

// BulkImport http hrandler for bulk operations
func (m *mdsController) BulkImport(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(mapError(err), gin.H{"error": err.Error()})
		return
	}

	if file == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file required"})
		return
	}

	fname := "uploaded/" + file.Filename
	err = c.SaveUploadedFile(file, fname)
	if err != nil {
		c.JSON(mapError(err), gin.H{"error": err.Error()})
		return
	}

	// just to avoid the user waiting for the processing
	go func(fn string) {
		err = m.blk.Import(fname)
		if err != nil {
			c.JSON(mapError(err), gin.H{"error": err.Error()})
			return
		}
	}(fname)

	c.JSON(http.StatusAccepted, gin.H{
		"message": "proccessing file",
	})
}

func mapError(err error) int {
	switch err.(type) {
	case *model.NegativeStokError:
		return http.StatusUnprocessableEntity
	case *model.ProductNotFoundError:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
