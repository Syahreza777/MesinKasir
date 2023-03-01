package storage

import "fmt"

type Tranksaksi struct {
	Barang Product
	AmbilBarang,
	SubTotal,
	Jumlah,
	Bayar,
	UangKurang,
	Kembalian,
	TotalHarga int
	Simpan []int
}

func (t *Tranksaksi) PilihBarang(barang []Product) {
	fmt.Print("Pilih barang yang anda inginkan: ")
	fmt.Scanf("%d\n", &t.AmbilBarang)

	t.Barang = barang[t.AmbilBarang-1]
}

func (t *Tranksaksi) JumlahBarang() {
	fmt.Print("Berapa jumlah barang yang anda beli? : ")
	fmt.Scanf("%d\n", &t.Jumlah)
}

func (t *Tranksaksi) HitungTotal() int {
	t.SubTotal = t.Barang.Harga * t.Jumlah
	fmt.Println("Sub total nya adalah Rp.", t.SubTotal)
	return t.SubTotal
}

func (t *Tranksaksi) SimpanSubTotal(num int) {
	t.Simpan = append(t.Simpan, num)
	fmt.Println()
}

func (t *Tranksaksi) TambahSemuaSubTotal() int {
	t.TotalHarga = 0
	for _, num := range t.Simpan {
		t.TotalHarga += num
	}
	fmt.Println("Total semua barang yang sudah dibeli Rp.", t.TotalHarga)
	return t.TotalHarga
}

func (t *Tranksaksi) MasukkanUang() {
	fmt.Print("Masukkan uang anda: ")
	fmt.Scanf("%d\n", &t.Bayar)

	for t.Bayar < t.TotalHarga {
		t.UangKurang = t.TotalHarga - t.Bayar
		fmt.Println("Maaf Uang Anda Kurang Rp.", t.UangKurang)
		fmt.Println()
		fmt.Print("Masukkan Uang Yang Dibayar : ")
		fmt.Scanf("%d\n", &t.Bayar)
	}

	if t.Bayar > t.TotalHarga {
		t.Kembalian = t.Bayar - t.TotalHarga
		fmt.Println("Uang anda lebih, ini kembaliannya: Rp.", t.Kembalian)
	}

	fmt.Println("Pembayaran Lunas")
	fmt.Println()
}
