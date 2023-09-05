package main

import("fmt")

type kendaraan struct {
	totalroda int

	kecepatanperjam int
}

type mobil struct {
	kendaraan
}

func (kendaraan *mobil) berjalan() {

	kendaraan.tambahkecepatan(10)

}

func (kendaraan *mobil) tambahkecepatan(kecepatanbaru int) {

	kendaraan.kecepatanperjam = kendaraan.kecepatanperjam + kecepatanbaru
	fmt.Println("Total Kecepatan Kendaraan : ",kendaraan.kecepatanperjam)

}

func main() {

	var mobilcepat mobil

	mobilcepat.berjalan()

	mobilcepat.berjalan()

	mobilcepat.berjalan()

	var mobillamban mobil

	mobillamban.berjalan()

}
