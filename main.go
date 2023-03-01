package main

import (
	"fmt"
	"mesin-kasir/storage"
	"strings"
)

func main() {
	show := storage.SimpanProduct{Products: []storage.Product{
		{Nama: "Buku", Harga: 4000},
		{Nama: "Pen", Harga: 2500},
		{Nama: "Setip", Harga: 1500},
	}}

	operasi := true
	Operasi := storage.Tranksaksi{}
	for operasi {
		show.TampilkanBarang()
		Operasi.PilihBarang(show.Products)
		Operasi.JumlahBarang()
		Operasi.SimpanSubTotal(Operasi.HitungTotal())
		var ulangi string

		fmt.Print("Mau nambah barang lagi ga?: ")
		fmt.Scanf("%s\n", &ulangi)
		fmt.Println()
		operasi = strings.ToLower(ulangi) == "ya"

		if operasi == false {
			Operasi.TambahSemuaSubTotal()
			Operasi.MasukkanUang()

			fmt.Print("Ingin lanjut tranksaksi lagi? (ya/n):")
			fmt.Scanf("%s\n", &ulangi)
			operasi = strings.ToLower(ulangi) == "ya"
			for i := range Operasi.Simpan {
				Operasi.Simpan[i] = 0
			}
		}

	}
	fmt.Println("Terima kasih sudah berbelanja disini!!")
	fmt.Println()
}
