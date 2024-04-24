package main

import (
	"coba-soal-challenge-1/challenge1"
	"fmt"
)

func main() {
	// var barangBelanjaan string
	// var ListBarang []string

	// for i := 0; i < 5; i++ {
	// 	fmt.Scan(&barangBelanjaan)
	// 	ListBarang = append(ListBarang, barangBelanjaan)
	// }

	// fmt.Println(ListBarang)

	///Challenge
	listBelanjaanCustomer := []string{"fanta", "cokelat", "parfum", "minyak goreng", "mie instan"}
	var strukBelanja []string
	var tempKasir challenge1.MenuKasir

	tempKasir = challenge1.Belanja{
		ListBarang: &strukBelanja,
	}

	for _, v := range listBelanjaanCustomer {
		tempKasir.TambahBarang(v)
	}

	fmt.Println(strukBelanja)

}
