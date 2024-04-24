package kasir

import "fmt"

type Belanja struct {
	ListBarang *[]string
}

type MenuKasir interface {
	TambahBarang(namaBarang string)
}

func (c Belanja) TambahBarang(namaBarang string) {
	*c.ListBarang = append(*c.ListBarang, namaBarang)
}

func (c Belanja) HapusBarang() {
	fmt.Println(*c.ListBarang)
}

type MenuKasir2 interface {
	HapusBarang()
}
