package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/rodaine/table"
)

func listItems(products *Products, w chan string) {
	byteValue, err := os.ReadFile("data.json")
	if err != nil {
		log.Println(err)
	}
	tbl := table.New("KB", "Nama Barang", "Harga Satuan")

	json.Unmarshal(byteValue, &products)

	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for _, v := range products.Items {
		tbl.AddRow(v.KodeBarang, v.NamaBarang, fmt.Sprintf("Rp %v", l.FmtNumber(float64(v.HargaSatuanRP), 2)))
	}
	tbl.Print()
	fmt.Println()
	w <- ""
}
