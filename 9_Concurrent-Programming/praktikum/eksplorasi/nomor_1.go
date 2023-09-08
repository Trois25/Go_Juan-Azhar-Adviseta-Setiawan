package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

// Struct untuk menyimpan setiap produk
type Products struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

func print(dataChannel <-chan Products, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Products Data")
	for product := range dataChannel {
		fmt.Println("===")
		fmt.Println("Title : ", product.Title)
		fmt.Println("Price : ", product.Price)
		fmt.Println("Category : ", product.Category)
		fmt.Println("===")
	}
}

func main() {
	// Membuat channel untuk mengirim data produk
	productsChannel := make(chan Products)

	// Membuat WaitGroup untuk menunggu goroutine selesai
	var wg sync.WaitGroup

	// Cek API menggunakan metode GET
	response, err := http.Get("https://fakestoreapi.com/products")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	defer response.Body.Close()

	// Mengambil data
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Membuat objek slice
	var products []Products

	// Mengurai JSON ke dalam slice
	json.Unmarshal(responseData, &products)

	// Menjalankan goroutine untuk mengirim data produk ke channel
	for _, product := range products {
		wg.Add(1)
		go func(p Products) {
			defer wg.Done()
			productsChannel <- p
		}(product)
	}

	// Menjalankan goroutine untuk mencetak produk dari channel
	go print(productsChannel, &wg)

	// Menunggu semua goroutine selesai
	wg.Wait()

	// Menutup channel setelah semua produk selesai dikirim
	close(productsChannel)
}
