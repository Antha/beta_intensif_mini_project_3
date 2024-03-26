package model_test

import (
	"encoding/csv"
	"fmt"
	"os"
	"sekolahbeta/hacker/mini_project3/config"
	"sekolahbeta/hacker/mini_project3/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImport(t *testing.T){
	Init()
	fileCsv, err := os.Open("D:\\Work\\HackerGolangClass\\Meeting - Project 3\\file\\sample_book.csv")
	
	defer fileCsv.Close()

	reader := csv.NewReader(fileCsv)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	books := utils.CsvToStruct(records)

	fmt.Println(books)

	for _, book := range books {
		err = book.Create(config.Mysql.DB)
	}
	

	assert.Nil(t, err)
}
