package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marcosvidolin/mdschallenge/bulk"
	"github.com/marcosvidolin/mdschallenge/controller"
	"github.com/marcosvidolin/mdschallenge/database"
	"github.com/marcosvidolin/mdschallenge/usecase"
)

func main() {

	db := &database.MdsDBClient{}

	fnd := usecase.NewFindProduct(db)

	cns := usecase.NewConsumeStock(db)

	ubk := usecase.NewBulkOperation(db,
		bulk.NewCsvSource(&bulk.Options{SkipHeader: true}),
	)

	ctlr := controller.New(cns, fnd, ubk)

	r := gin.Default()
	r.GET("/mds/countries/:country/products/:sku", ctlr.FindByCountryAndSku)
	r.PATCH("/mds/countries/:country/products/:sku/consume", ctlr.ConsumeStock)
	r.POST("/mds/orders/bulkimport", ctlr.BulkImport)
	r.Run()

}
