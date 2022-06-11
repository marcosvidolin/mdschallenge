package util

import "github.com/marcosvidolin/mdschallenge/model"

func ChunkSlice(slice []model.Product, chunkSize int) [][]model.Product {
	var chunks [][]model.Product
	for {
		if len(slice) == 0 {
			break
		}

		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}

		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}

	return chunks
}
