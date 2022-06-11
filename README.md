# mdschallenge

1. Provide a products API
  a. Get a product by SKU
	b. Consume stock from a product.
		- Should validate if the stock requested is available first, and then decrease it.

2. Provide an API that allows a bulk update of orders from a CSV.
  a. For each CSV line, the stock update could be positive or negative
	b. If a product doesn’t exist, it should be created.

## How to run

```shell
go run main.go
```

## API Request examples

- To upload a CSV file

```bash
curl -X POST http://localhost:8080/mds/orders/bulkimport -F 'file=@file_2.csv'
```

- To get a Product

```bash
curl http://localhost:8080/mds/countries/ke/products/da8ef851e075
```

- To consume the stock

```bash
curl -X PATCH http://localhost:8080/mds/countries/ke/products/da8ef851e075/consume -d '{"amount": 10}'
```
