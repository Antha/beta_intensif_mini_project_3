package main

import (
	// "bufio"
	// "encoding/json"
	"bufio"
	"fmt"
	"log"
	"os"
	"sekolahbeta/hacker/mini_project3/config"
	"sekolahbeta/hacker/mini_project3/model"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	// "strconv"
	// "strings"
	// "sync"
)

func Init()  {
	err := godotenv.Load(".env")
	if err != nil{
		fmt.Println("env not found, using global .env")
	}
	config.OpenDB()
}

func main() {
	Init()
	// ListOfBook = append(ListOfBook, MyLibrary{
	// 	isbn:    "1001",
	// 	judul:        "Conan Movie",
	// 	penulis:       "Aoyama Gosho",
	// 	publisher:    "Toei Animation",
	// 	page_number:  99,
	// 	tahun:         2019,
	// })
	
	Menu()
}

func Menu()  {
	// pilihanMenu := 0
	var pilihanMenu int
	fmt.Println("=================================")
	fmt.Println("Sistem Manajemen Perpustakaan")
	fmt.Println("=================================")
	fmt.Println("Silahkan Pilih : ")
	fmt.Println("1. Tambah Buku")
	fmt.Println("2. Lihat Daftar Buku")
	fmt.Println("3. Edit Buku")
	fmt.Println("4. Hapus Buku")
	fmt.Println("5. Keluar")
	fmt.Println("=================================")
	fmt.Print("Masukan Pilihan : ")
	_, err := fmt.Scanln(&pilihanMenu)
	if err != nil {
		fmt.Println("Terjadi error:", err)
	}

	fmt.Println("")

	switch pilihanMenu {
	case 1:
		AddBook()
	case 2:
		ViewBook()
	case 3:
		EditBook()
	case 4:
		//DeleteBook()
	case 5:
		os.Exit(0)
	}

	Menu()
}

func ViewBook() {
	//panic("unimplemented")
	fmt.Println("=================================")
	fmt.Println("Daftar Buku")
	fmt.Println("=================================")
	fmt.Println("Memuat data ...")

	bookData := model.Book{}

	res, err := bookData.GetAll(config.Mysql.DB)

	if(err != nil){
		log.Fatal(err)
	}

	for urutan, book := range res {
		fmt.Printf("%d. ISBN: %s, Penulis : %s, Tahun : %d, Judul : %s, Gambar : %s, Stok : %d\n",
			urutan+1,
			book.ISBN,
			book.Penulis,
			book.Tahun,
			book.Judul,
			book.Gambar,
			book.Stok,
		)
	}
}

func AddBook() {
	inputanUser := bufio.NewReader(os.Stdin)
	isbn := ""
	penulis := ""
	//tahun := 0
	judul := ""
	gambar := ""
	//page_number := 0
	

	draftBook := [] model.Book{}

	for{
		fmt.Print("Silahkan Masukan ISBN Buku : ")
		isbn, _ = inputanUser.ReadString('\n')
		isbn = strings.TrimSpace(isbn)
		isbn = strings.Replace(
			isbn,
			"\n",
			"",
			1)
		isbn = strings.Replace(
			isbn,
			"\r",
			"",
			1)

		fmt.Print("Silahkan Masukan Penulis : ")
		penulis, _ = inputanUser.ReadString('\n')
		penulis = strings.TrimSpace(penulis)
		penulis = strings.Replace(
			penulis,
			"\n",
			"",
			1)
		penulis = strings.Replace(
			penulis,
			"\r",
			"",
			1)

		fmt.Print("Silahkan Masukan Tahun Terbit : ")
		tahunInput, _ := inputanUser.ReadString('\n')
		tahunInput = strings.TrimSpace(tahunInput)
		tahunInput = strings.Replace(tahunInput, "\n", "", 1)
		tahunInput = strings.Replace(tahunInput, "\r", "", 1)
		
		// Convert tahunInput to an integer
		tahun, err := strconv.Atoi(tahunInput)
		if err != nil {
			fmt.Println("Error converting input to integer:", err)
			// Handle the error appropriately, such as asking the user to input again
		}

		// Memeriksa apakah kode buku sudah digunakan sebelumnya
		fmt.Print("Silahkan Masukan Judul : ")
		judul, _ = inputanUser.ReadString('\n')
		judul = strings.TrimSpace(judul)
		judul = strings.Replace(
			judul,
			"\n",
			"",
			1)
		judul = strings.Replace(
			judul,
			"\r",
			"",
			1)

		
		fmt.Print("Silahkan Masukan Link Gambar : ")
		gambar, _ = inputanUser.ReadString('\n')
		gambar = strings.TrimSpace(gambar)
		gambar = strings.Replace(
			gambar,
			"\n",
			"",
			1)
		gambar = strings.Replace(
			gambar,
			"\r",
			"",
			1)

		fmt.Print("Silahkan Masukan Jumlah Stok : ")
		stokInput, _ := inputanUser.ReadString('\n')
		stokInput = strings.TrimSpace(stokInput)
		stokInput = strings.Replace(stokInput, "\n", "", 1)
		stokInput = strings.Replace(stokInput, "\r", "", 1)
		
		// Convert pageInput to an integer
		stok, err := strconv.Atoi(stokInput)
		if err != nil {
			fmt.Println("Error converting input to integer:", err)
			// Handle the error appropriately
		}

		fmt.Println("")

		var pilihanMenuPesanan = 0
		fmt.Println("Ketik 1 atau tombol lain untuk tambah pesanan, ketik 0 untuk keluar")
		stringInput, _ := inputanUser.ReadString('\n')
		stringInput = strings.TrimSpace(stringInput)
		pilihanMenuPesanan, _ = strconv.Atoi(stringInput)
		
		draftBook = append(draftBook, model.Book{
			ISBN:  isbn,
			Penulis: penulis,
			Tahun: uint(tahun),
			Judul: judul,
			Gambar: gambar,
			Stok : uint(stok),
		})

		if pilihanMenuPesanan == 0 {
			break
		}
	}

	for _, book := range draftBook {
		err := book.Create(config.Mysql.DB)
		if err != nil{
			log.Fatal(err)
		}	
	} 

	fmt.Println("Berhasil Menambah Buku!")
}

func EditBook() {
	inputanUser := bufio.NewReader(os.Stdin)
	isbn := ""
	penulis := ""
	judul := ""
	gambar := ""

	fmt.Print("Silahkan Masukan ID Buku Yang Mau Diedit : ")
	idInput, _ := inputanUser.ReadString('\n')
	idInput = strings.TrimSpace(idInput)
	idInput = strings.Replace(idInput, "\n", "", 1)
	idInput = strings.Replace(idInput, "\r", "", 1)
	// Convert tahunInput to an integer
	id, err := strconv.Atoi(idInput)
	if err != nil {
		fmt.Println("Error converting input to integer:", err)
		// Handle the error appropriately, such as asking the user to input again
	}


	fmt.Print("Silahkan Masukan ISBN Buku : ")
		isbn, _ = inputanUser.ReadString('\n')
		isbn = strings.TrimSpace(isbn)
		isbn = strings.Replace(
			isbn,
			"\n",
			"",
			1)
		isbn = strings.Replace(
			isbn,
			"\r",
			"",
			1)

	fmt.Print("Silahkan Masukan Penulis : ")
		penulis, _ = inputanUser.ReadString('\n')
		penulis = strings.TrimSpace(penulis)
		penulis = strings.Replace(
			penulis,
			"\n",
			"",
			1)
		penulis = strings.Replace(
			penulis,
			"\r",
			"",
			1)

	fmt.Print("Silahkan Masukan Tahun Terbit : ")
	tahunInput, _ := inputanUser.ReadString('\n')
	tahunInput = strings.TrimSpace(tahunInput)
	tahunInput = strings.Replace(tahunInput, "\n", "", 1)
	tahunInput = strings.Replace(tahunInput, "\r", "", 1)
	
	// Convert tahunInput to an integer
	tahun, err := strconv.Atoi(tahunInput)
	if err != nil {
		fmt.Println("Error converting input to integer:", err)
		// Handle the error appropriately, such as asking the user to input again
	}

	// Memeriksa apakah kode buku sudah digunakan sebelumnya
	fmt.Print("Silahkan Masukan Judul : ")
	judul, _ = inputanUser.ReadString('\n')
	judul = strings.TrimSpace(judul)
	judul = strings.Replace(
		judul,
		"\n",
		"",
		1)
	judul = strings.Replace(
		judul,
		"\r",
		"",
		1)

	
	fmt.Print("Silahkan Masukan Link Gambar : ")
	gambar, _ = inputanUser.ReadString('\n')
	gambar = strings.TrimSpace(gambar)
	gambar = strings.Replace(
		gambar,
		"\n",
		"",
		1)
	gambar = strings.Replace(
		gambar,
		"\r",
		"",
		1)

	fmt.Print("Silahkan Masukan Jumlah Stok : ")
	stokInput, _ := inputanUser.ReadString('\n')
	stokInput = strings.TrimSpace(stokInput)
	stokInput = strings.Replace(stokInput, "\n", "", 1)
	stokInput = strings.Replace(stokInput, "\r", "", 1)
	
	// Convert pageInput to an integer
	stok, err := strconv.Atoi(stokInput)
	if err != nil {
		fmt.Println("Error converting input to integer:", err)
		// Handle the error appropriately
	}

	fmt.Println("")
	
	book := model.Book{
		ID: uint(id),
		ISBN:  isbn,
		Penulis: penulis,
		Tahun: uint(tahun),
		Judul: judul,
		Gambar: gambar,
		Stok : uint(stok),
	}


	err = book.UpdateOne(config.Mysql.DB)
	if err != nil{
		log.Fatal(err)
	}	

	fmt.Printf("Berhasil Mengedit Buku Dengan id %d! \n",id)

	Menu()
}