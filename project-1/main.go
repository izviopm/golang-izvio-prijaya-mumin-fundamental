package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

func main() {
	books := [][]string{}

	fmt.Println("==================== Console Books Apps ====================")
	fmt.Println("(1/Show All)      Tampilkan Semua Data")
	fmt.Println("(2/Add Data)      Menambahkan Data Baru")
	fmt.Println("(3/Filter Data)   Menampilkan Data Berdasarkan ID")
	fmt.Println("(4/Update Data)   Mengubah Salah Satu Data Berdasarkan ID")
	fmt.Println("(5/Delete Data)   Menghapus Salah Satu Data Berdasarkan ID")
	fmt.Println("(6/Exit)          Keluar")
	fmt.Println("============================================================")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := strings.ToLower(scanner.Text())

		if command == "show all" || command == "1" {
			fmt.Println("Menampilkan Semua Data:")
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"ID", "Judul", "Pengarang", "Penerbit", "Genre"})

			for _, v := range books {
				table.Append(v)
			}

			table.Render()
		} else if command == "add data" || command == "2" {
			fmt.Println("Menambahkan Data Buku Baru")
			fmt.Println("Masukan ID Buku:")
			scanner.Scan()
			bookID := strings.ToUpper(scanner.Text())

			fmt.Println("Masukan Judul Buku:")
			scanner.Scan()
			bookTitle := scanner.Text()

			fmt.Println("Masukan Nama Pengarang:")
			scanner.Scan()
			bookAuthor := scanner.Text()

			fmt.Println("Masukan Nama Penerbit:")
			scanner.Scan()
			bookPublisher := scanner.Text()

			fmt.Println("Masukan Genre Buku:")
			scanner.Scan()
			bookGenre := scanner.Text()

			newData := []string{bookID, bookTitle, bookAuthor, bookPublisher, bookGenre}
			books = append(books, newData)

			fmt.Println("Data Buku baru berhasil ditambahkan.")
		} else if command == "exit" || command == "6" {
			fmt.Println("Program selesai.")
			break
		} else {
			fmt.Println("Perintah tidak valid.")
		}

		time.Sleep(2 * time.Second)
		fmt.Println("\n\n==================== Console Books Apps ====================")
		fmt.Println("(1/Show All)      Tampilkan Semua Data")
		fmt.Println("(2/Add Data)      Menambahkan Data Baru")
		fmt.Println("(3/Filter Data)   Menampilkan Data Berdasarkan ID")
		fmt.Println("(4/Update Data)   Mengubah Salah Satu Data Berdasarkan ID")
		fmt.Println("(5/Delete Data)   Menghapus Salah Satu Data Berdasarkan ID")
		fmt.Println("(6/Exit)          Keluar")
		fmt.Println("============================================================")
	}
}