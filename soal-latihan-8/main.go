package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/go-playground/locales/en_US"
	"github.com/rodaine/table"
)

type Products struct {
	Items []Product `json:"products"`
}
type Product struct {
	KodeBarang    string `json:"kode_barang"`
	NamaBarang    string `json:"nama_barang"`
	HargaSatuanRP int    `json:"harga_satuan_rp"`
}

func main() {
	l := en_US.New()
	byteValue, err := os.ReadFile("data.json")
	if err != nil {
		log.Println(err)
	}
	var products Products
	var namaPembeli, kodeBarang string
	var totalBayar, jumlahBeli, totalPembelian, discount, potonganHarga int
	json.Unmarshal(byteValue, &products)
	fmt.Print("Nama Pembeli: ")
	fmt.Scanln(&namaPembeli)

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("KB", "Nama Barang", "Harga Satuan")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, v := range products.Items {
		tbl.AddRow(v.KodeBarang, v.NamaBarang, fmt.Sprintf("Rp %v", l.FmtNumber(float64(v.HargaSatuanRP), 2)))
	}
	tbl.Print()
	fmt.Print("Kode Barang: ")
	fmt.Scanln(&kodeBarang)
	hargaBarang := 0
	namaBarang := ""
	for _, v := range products.Items {
		if v.KodeBarang == strings.ToUpper(kodeBarang) {
			hargaBarang = v.HargaSatuanRP
			namaBarang = v.NamaBarang
		}
	}
	tbl = table.New("Nama Barang", "Harga Satuan")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	tbl.AddRow(namaBarang, fmt.Sprintf("Rp %v", l.FmtNumber(float64(hargaBarang), 2)))
	tbl.Print()
	fmt.Print("Jumlah Beli: ")
	fmt.Scanln(&jumlahBeli)
	totalPembelian = hargaBarang * jumlahBeli
	fmt.Println("Total Pembelian: ", fmt.Sprintf("Rp %v", l.FmtNumber(float64(totalPembelian), 2)))

	if totalPembelian >= 7_500_000 {
		discount = 5
		potonganHarga = totalPembelian / 100 * 5
	}
	totalBayar = totalPembelian - potonganHarga

	fmt.Print("Mendapat potongan ", discount, "% ")
	fmt.Printf("(-%v)\n", fmt.Sprintf("Rp %v", l.FmtNumber(float64(potonganHarga), 2)))
	fmt.Println("Total Bayar: ", fmt.Sprintf("Rp %v", l.FmtNumber(float64(totalBayar), 2)))
}
