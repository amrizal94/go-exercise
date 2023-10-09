package main

import (
	"fmt"

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

var headerFmt = color.New(color.FgGreen, color.Underline).SprintfFunc()
var columnFmt = color.New(color.FgYellow).SprintfFunc()
var l = en_US.New()
var products Products
var inputKodeBarang, namaBarang string
var hargaBarang int

func main() {

	var namaPembeli string
	var totalBayar, jumlahBeli, totalPembelian, discount, potonganHarga int
	wait := make(chan string)

	fmt.Print("Nama Pembeli: ")
	fmt.Scanln(&namaPembeli)
	fmt.Println()
	go listItems(wait)
	<-wait
	fmt.Print("Kode Barang: ")
	fmt.Scanln(&inputKodeBarang)
	fmt.Println()
	go transaction(wait)
	<-wait
	fmt.Print("Jumlah Beli: ")
	fmt.Scanln(&jumlahBeli)
	totalPembelian = hargaBarang * jumlahBeli

	if totalPembelian >= 7_500_000 {
		discount = 5
	}
	potonganHarga = totalPembelian / 100 * discount
	totalBayar = totalPembelian - potonganHarga

	tbl := table.New("", "", "")
	tbl.WithFirstColumnFormatter(columnFmt)
	tbl.AddRow("Total Pembelian", ":", fmt.Sprintf("Rp %v", l.FmtNumber(float64(totalPembelian), 2)))
	tbl.AddRow("Mendapat potongan", ":", fmt.Sprint(discount, "% ")+fmt.Sprintf("(-%v)", fmt.Sprintf("Rp %v", l.FmtNumber(float64(potonganHarga), 2))))
	tbl.AddRow("Total Bayar", ":", fmt.Sprintf("Rp %v", l.FmtNumber(float64(totalBayar), 2)))
	tbl.Print()
}
