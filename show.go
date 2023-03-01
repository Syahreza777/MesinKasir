package storage

import "fmt"

type Product struct {
	Nama  string
	Harga int
}

type SimpanProduct struct {
	Products []Product
}

func (s *SimpanProduct) TampilkanBarang() {
	fmt.Println()
	fmt.Println("Selamat datang di toko saya bang!!")
	fmt.Println("==================================")
	fmt.Println("Daftar Barang: ")

	for i, item := range s.Products {
		fmt.Println(i+1., item.Nama, ": ", "Rp.", item.Harga)
	}

	fmt.Println("==================================")
}
