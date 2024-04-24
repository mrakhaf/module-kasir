package main

import (
	"fmt"

	"github.com/mrakhaf/module-kasir/kasir"
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
	var tempKasir kasir.MenuKasir

	tempKasir = kasir.Belanja{
		ListBarang: &strukBelanja,
	}

	for _, v := range listBelanjaanCustomer {
		tempKasir.TambahBarang(v)
	}

	fmt.Println(strukBelanja)

}
