package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/go-pdf/fpdf"
)

type MyLibrary struct{
	book_kode   string 
	title       string
	writer      string
	publisher   string 
	page_number int    
	year        int    
}

var ListOfBook [] MyLibrary 
var firstInit = false

func main() {
	// ListOfBook = append(ListOfBook, MyLibrary{
	// 	book_kode:    "1001",
	// 	title:        "Conan Movie",
	// 	writer:       "Aoyama Gosho",
	// 	publisher:    "Toei Animation",
	// 	page_number:  99,
	// 	year:         2019,
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
		DeleteBook()
	case 5:
		os.Exit(0)
	}

	Menu()
}

func DeleteBook() {
	fmt.Println("=================================")
	fmt.Println("Hapus Buku")
	fmt.Println("=================================")
	ViewBook()

	fmt.Println("=================================")
	var urutanBook int
	fmt.Print("Masukan Urutan Pesanan : ")
	_, err := fmt.Scanln(&urutanBook)
	if err != nil {
		fmt.Println("Terjadi error:", err)
	}

	if (urutanBook-1) < 0 ||
		(urutanBook-1) > len(ListOfBook) {
		fmt.Println("Urutan Buku Tidak Sesuai")
		DeleteBook()
		return
	}

	err = os.Remove(fmt.Sprintf("json/book-%s.json", ListOfBook[urutanBook-1].book_kode))
	if err != nil {
		fmt.Println("Terjadi error:", err)
	}

	fmt.Println("Pesanan Berhasil Dihapus!")
}

func EditBook() {
	// Tampilkan daftar buku untuk dipilih
	ViewBook()

	fmt.Println("=================================")
	fmt.Println("Edit Buku")
	fmt.Println("=================================")

	// Meminta input dari pengguna untuk memilih buku yang akan diedit
	var urutanBook int
	fmt.Print("Masukkan Urutan Buku yang Ingin Diedit: ")
	_, err := fmt.Scanln(&urutanBook)
	if err != nil {
		fmt.Println("Terjadi error:", err)
		return
	}

	// Validasi urutan buku
	if urutanBook < 1 || urutanBook > len(ListOfBook) {
		fmt.Println("Urutan buku tidak valid.")
		return
	}

	// Meminta input baru dari pengguna
	inputanUser := bufio.NewReader(os.Stdin)

	fmt.Println("Masukkan informasi buku yang diperbarui:")

	// Meminta input baru dari pengguna
	fmt.Print("Judul: ")
	newTitle, _ := inputanUser.ReadString('\n')
	newTitle = strings.TrimSpace(newTitle)

	fmt.Print("Penulis: ")
	newWriter, _ := inputanUser.ReadString('\n')
	newWriter = strings.TrimSpace(newWriter)

	fmt.Print("Penerbit: ")
	newPublisher, _ := inputanUser.ReadString('\n')
	newPublisher = strings.TrimSpace(newPublisher)

	fmt.Print("Jumlah Halaman: ")
	newPageInput, _ := inputanUser.ReadString('\n')
	newPageInput = strings.TrimSpace(newPageInput)
	newPageInput = strings.Replace(newPageInput, "\n", "", 1)
	newPageInput = strings.Replace(newPageInput, "\r", "", 1)

	newPageNumber, err := strconv.Atoi(newPageInput)
	if err != nil {
		fmt.Println("Error converting input to integer:", err)
		return
	}

	fmt.Print("Tahun Terbit: ")
	newYearInput, _ := inputanUser.ReadString('\n')
	newYearInput = strings.TrimSpace(newYearInput)
	newYearInput = strings.Replace(newYearInput, "\n", "", 1)
	newYearInput = strings.Replace(newYearInput, "\r", "", 1)

	newYear, err := strconv.Atoi(newYearInput)
	if err != nil {
		fmt.Println("Error converting input to integer:", err)
		return
	}

	// Update informasi buku dalam struktur data
	ListOfBook[urutanBook-1].title = newTitle
	ListOfBook[urutanBook-1].writer = newWriter
	ListOfBook[urutanBook-1].publisher = newPublisher
	ListOfBook[urutanBook-1].page_number = newPageNumber
	ListOfBook[urutanBook-1].year = newYear

	// Simpan kembali ke file JSON
	saveEditedBook(ListOfBook[urutanBook-1])

	fmt.Println("Buku berhasil diperbarui.")
}	

func saveEditedBook(book MyLibrary) {
    dataJson, err := json.Marshal(book)
    if err != nil {
        fmt.Println("Terjadi error:", err)
        return
    }

    err = os.WriteFile(fmt.Sprintf("D:\\Work\\HackerGolangClass\\library\\json\\book-%s.json", book.book_kode), dataJson, 0644)
    if err != nil {
        fmt.Println("Terjadi error:", err)
        return
    }
}


func ViewBook() {
	//panic("unimplemented")
	fmt.Println("=================================")
	fmt.Println("Daftar Buku")
	fmt.Println("=================================")
	fmt.Println("Memuat data ...")
	ListBook := [] MyLibrary{}

	listJsonBook, err := os.ReadDir("D:\\Work\\HackerGolangClass\\library\\json")
	if err != nil {
		fmt.Println("Terjadi error: ", err)
	}

	wg := sync.WaitGroup{}

	ch := make(chan string)
	chBook := make(chan MyLibrary, len(listJsonBook))

	jumlahPelayan := 5

	for i := 0; i < jumlahPelayan; i++ {
		wg.Add(1)
		go LookBookPerFile(ch, chBook, &wg)
	}

	for _, fileBook := range listJsonBook {
		ch <- fileBook.Name()
	}

	close(ch)

	wg.Wait()

	close(chBook)

	for dataBook := range chBook {
		ListOfBook = append(ListBook, dataBook)
	}

	for urutan, book := range ListBook {
		fmt.Printf("%d. Kode Buku : %s, Judul : %s\n",
			urutan+1,
			book.book_kode,
			book.title,
		)
	}
}

func LookBookPerFile(ch <-chan string, chBook chan MyLibrary, wg *sync.WaitGroup){
	var book MyLibrary
	for book_name := range ch {
		fmt.Println(">>>",book_name)
		dataJSON, err := os.ReadFile(fmt.Sprintf("D:\\Work\\HackerGolangClass\\library\\json\\%s", book_name))
		if err != nil {
			fmt.Println("Terjadi error:", err)
		}

		//fmt.Println("dataJSON",dataJSON)
		fmt.Println("dataJSON", string(dataJSON))

		err = json.Unmarshal(dataJSON, &book)
		if err != nil {
			fmt.Println("Terjadi error:", err)
		}

		fmt.Println("book_kode",book.book_kode)

		chBook <- book
	}
	wg.Done()

}

func AddBook() {
	ViewBook()

	inputanUser := bufio.NewReader(os.Stdin)
	book_kode := ""
	title := ""
	writer := ""
	publisher := ""
	// page_number := 01
	// year := 0

	draftBook := [] MyLibrary{}

	for{
		fmt.Print("Silahkan Masukan Kode Buku : ")
		book_kode, _ = inputanUser.ReadString('\n')
		book_kode = strings.TrimSpace(book_kode)
		book_kode = strings.Replace(
			book_kode,
			"\n",
			"",
			1)
		book_kode = strings.Replace(
			book_kode,
			"\r",
			"",
			1)

		// Memeriksa apakah kode buku sudah digunakan sebelumnya
        if isBookCodeUsed(book_kode) {
            fmt.Println("Kode buku sudah digunakan. Silahkan masukkan kode buku yang berbeda.")
            continue
        }

		fmt.Print("Silahkan Masukan Judul : ")
		title, _ = inputanUser.ReadString('\n')
		title = strings.TrimSpace(title)
		title = strings.Replace(
			title,
			"\n",
			"",
			1)
		title = strings.Replace(
			title,
			"\r",
			"",
			1)

		fmt.Print("Silahkan Masukan Penulis : ")
		writer, _ = inputanUser.ReadString('\n')
		writer = strings.TrimSpace(writer)
		writer = strings.Replace(
			writer,
			"\n",
			"",
			1)
		writer = strings.Replace(
			writer,
			"\r",
			"",
			1)

		fmt.Print("Silahkan Masukan Publisher : ")
		publisher, _ = inputanUser.ReadString('\n')
		publisher = strings.TrimSpace(publisher)
		publisher = strings.Replace(
			publisher,
			"\n",
			"",
			1)
		publisher = strings.Replace(
			publisher,
			"\r",
			"",
			1)

		fmt.Print("Silahkan Masukan Jumlah Halaman : ")
		pageInput, _ := inputanUser.ReadString('\n')
		pageInput = strings.TrimSpace(pageInput)
		pageInput = strings.Replace(pageInput, "\n", "", 1)
		pageInput = strings.Replace(pageInput, "\r", "", 1)
		
		// Convert pageInput to an integer
		page_number, err := strconv.Atoi(pageInput)
		if err != nil {
			fmt.Println("Error converting input to integer:", err)
			// Handle the error appropriately
		}

		fmt.Print("Silahkan Masukan Tahun Terbit : ")
		yearInput, _ := inputanUser.ReadString('\n')
		yearInput = strings.TrimSpace(yearInput)
		yearInput = strings.Replace(yearInput, "\n", "", 1)
		yearInput = strings.Replace(yearInput, "\r", "", 1)
		
		// Convert yearInput to an integer
		year, err := strconv.Atoi(yearInput)
		if err != nil {
			fmt.Println("Error converting input to integer:", err)
			// Handle the error appropriately, such as asking the user to input again
		}

		// ListOfBook = append(ListOfBook, MyLibrary{
		// 	book_kode:    book_kode,
		// 	title:        title,
		// 	writer:       writer,
		// 	publisher:    publisher,
		// 	page_number:  page_number,
		// 	year:         year,
		// })

		fmt.Println("Berhasil Menambah Buku!")
		fmt.Println("")

		var pilihanMenuPesanan = 0
		fmt.Println("Ketik 1 atau tombol lain untuk tambah pesanan, ketik 0 untuk keluar")
		stringInput, _ := inputanUser.ReadString('\n')
		stringInput = strings.TrimSpace(stringInput)
		pilihanMenuPesanan, _ = strconv.Atoi(stringInput)
		
		draftBook = append(draftBook, MyLibrary{
			book_kode:    book_kode,
			title:        title,
			writer:       writer,
			publisher:    publisher,
			page_number:  page_number,
			year:         year,
		})

		if pilihanMenuPesanan == 0 {
			break
		}
	}

	dataJson, err := json.Marshal(draftBook)
	if err != nil {
		fmt.Println("Terjadi error:", err)
	}
	fmt.Println("Tolong Print Data Json")
	fmt.Println(string(dataJson))

	//2ListOfBook = append(ListOfBook, draftBook...)

	fmt.Println("Menambah Buku...")

	_ = os.Mkdir("D:\\Work\\HackerGolangClass\\library\\json", 0777)

	ch := make(chan MyLibrary)

	wg := sync.WaitGroup{}

	jumlahPelayan := 4

	// Mendaftarkan receiver/pemroses data
	for i := 0; i < jumlahPelayan; i++ {
		wg.Add(1)
		go saveBook(ch, &wg, i)
	}

	// Mengirimkan data ke channel
	for _, book := range draftBook {
		fmt.Println(book)
		ch <- book
	}

	close(ch)

	wg.Wait()

	fmt.Println("Berhasil Menambah Buku!")
}

func saveBook(ch <-chan MyLibrary, wg *sync.WaitGroup, noServ int) {
	//fmt.Printf("heelo")
	//fmt.Println(ch)
	for book := range ch {
		fmt.Println("-book=")
		fmt.Println(book)
		dataJson, err := json.Marshal(book)
		if err != nil {
			fmt.Println("Terjadi error:", err)
		}

		fmt.Println("dataJson");	
		fmt.Println(dataJson)
		fmt.Println("Data JSON:", string(dataJson))

		err = os.WriteFile(fmt.Sprintf("D:\\Work\\HackerGolangClass\\library\\json\\book-%s.json", book.book_kode), dataJson, 0644)
		if err != nil {
			fmt.Println("Terjadi error:", err)
		}

		fmt.Printf("Service No %d Process Book1 ID : %s!\n", noServ, book.book_kode)
	}
	wg.Done()
}

func GeneratePdfPesanan() {
	ViewBook()
	fmt.Println("=================================")
	fmt.Println("Membuat Daftar Buku ...")
	fmt.Println("=================================")
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "", 12)
	pdf.SetLeftMargin(10)
	pdf.SetRightMargin(10)

	for i, book := range ListOfBook {
		bookText := fmt.Sprintf(
			"Book #%d:\nID : %s\nTitle : %s\nPublisher : %d\nWriter : %d\Page Number : %s\n",
			i+1, book.book_kode, book.publisher, book.writer, book.page_number,


		pdf.MultiCell(0, 10, bookText, "0", "L", false)
		pdf.Ln(5)
	}

	err := pdf.OutputFileAndClose(
		fmt.Sprintf("daftar_buku_%s.pdf",
			time.Now().Format("2006-01-02-15-04-05")))

	if err != nil {
		fmt.Println("Terjadi error:", err)
	}
}

func isBookCodeUsed(code string) bool {
    for _, book := range ListOfBook {
        if book.book_kode == code {
            return true
        }
    }
    return false
}