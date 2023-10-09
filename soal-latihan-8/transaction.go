package main

import (
	"fmt"
	"strings"

	"github.com/rodaine/table"
)

var transaction = func(hargaBarang *int, inputKodeBarang *string, products *Products, w chan string) {
	var namaBarang string
	for _, v := range products.Items {
		if v.KodeBarang == strings.ToUpper(*inputKodeBarang) {
			*hargaBarang = v.HargaSatuanRP
			namaBarang = v.NamaBarang
		}
	}
	tbl := table.New("Nama Barang", "Harga Satuan")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	tbl.AddRow(namaBarang, fmt.Sprintf("Rp %v", l.FmtNumber(float64(*hargaBarang), 2)))
	tbl.Print()
	fmt.Println()
	w <- ""
}
