package model_test

import (
	"encoding/json"
	"fmt"
	"log"
	"sekolahbeta/hacker/mini_project3/config"
	"sekolahbeta/hacker/mini_project3/model"

	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Init()  {
	err := godotenv.Load("../.env")
	if err != nil{
		fmt.Println("env not found, using global .env")
	}
	config.OpenDB()
}

func TestDeleteAllBook(t *testing.T) {
    Init()

	bookData := model.Book{}

	err := bookData.DeleteAll(config.Mysql.DB)

	if(err != nil){
		log.Fatal(err)
	}

	assert.Nil(t, err)
}

func TestCreateBook(t *testing.T)  {
	Init()

	bookData := []model.Book{
				{
					ISBN : "137189391",
					Penulis : "Aoyama Gosho",
					Tahun : 1998,
					Judul : "Conan The Sniper",
					Gambar: "rusa.jpg",
					Stok: 10,
				},
				{
					ISBN : "137112314",
					Penulis : "Conan Indihome",
					Tahun : 1998,
					Judul : "Conan The Star",
					Gambar: "rusa.jpg",
					Stok: 13,
				},
			}
	
	var err error 

	for _,book := range bookData{
		err := book.Create(config.Mysql.DB)
		if err != nil{
			log.Fatal(err)
		}	
	}

	assert.Nil(t, err)
}

func TestUpdate(t *testing.T){
	Init()

	bookData := model.Book{
		ID : 1,
		ISBN : "137189391",
		Penulis : "Aoyama Gosho * Peety",
		Tahun : 1998,
		Judul : "Conan The Bright",
		Gambar: "rusa.jpg",
		Stok: 10,
	}

    err := bookData.UpdateOne(config.Mysql.DB)

	if(err != nil){
		log.Fatal(err)
	}

	assert.Nil(t, err)
}

func TestGetByAll(t *testing.T){
	Init()

	bookData := model.Book{}

	res, err := bookData.GetAll(config.Mysql.DB)
	assert.Nil(t, err)
	resJson, err := json.Marshal(res)

	if(err != nil){
		log.Fatal(err)
	}

	resJsonString := string(resJson)

	fmt.Println(resJsonString)
}