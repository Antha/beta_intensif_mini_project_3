package utils

import (
	"fmt"
	"sekolahbeta/hacker/mini_project3/model"
	"strconv"
)

func CsvToStruct(records [][]string) [] model.Book {
	books := [] model.Book{}

	for _, book := range records {
		idUint64, err := strconv.ParseUint(book[0], 10, 64)
		if err != nil {
			fmt.Println("Gagal mengonversi ID:", err)
			continue // lanjutkan ke iterasi berikutnya jika terjadi kesalahan
		}

		id := uint(idUint64)

		tahunUint64, err := strconv.ParseUint(book[3], 10, 64)
		if err != nil {
			fmt.Println("Gagal mengonversi Tahun:", err)
			continue // lanjutkan ke iterasi berikutnya jika terjadi kesalahan
		}

		tahun := uint(tahunUint64)

		stokUint64, err := strconv.ParseUint(book[6], 10, 64)
		if err != nil {
			fmt.Println("Gagal mengonversi Stok:", err)
			continue // lanjutkan ke iterasi berikutnya jika terjadi kesalahan
		}

		stok := uint(stokUint64)

		books = append(books, model.Book{
			ID:            id,
			ISBN:          book[1],
			Penulis:       book[2],
			Tahun:         tahun,
			Judul:         book[4],
			Gambar:        book[5],
			Stok:          stok,
		})
	}

	return books
}
