package main

import (
	"fmt"
	"math"
)

type shape interface {
	luas() float32
}

type persegiPanjang struct {
	panjang, lebar float32
}

type persegi struct {
	sisi float32
}

type lingkaran struct {
	jari2 float32
}

type segitiga struct {
	alas, tinggi float32
}

func (s segitiga) luas() float32 {
	return s.alas * s.tinggi * 1 / 2
}

func (p persegiPanjang) luas() float32 {
	return p.panjang * p.lebar
}
func (p persegi) luas() float32 {
	return p.sisi * p.sisi
}

func (l lingkaran) luas() float32 {
	return math.Pi * l.jari2 * l.jari2
}

var hitung = func(s shape) {
	fmt.Printf("%v cm²\n", s.luas())
}

func main() {
	hitung(lingkaran{
		jari2: 5,
	})
	hitung(persegiPanjang{
		panjang: 10,
		lebar:   10,
	})

	hitung(persegi{sisi: 7})
	hitung(segitiga{
		alas:   20,
		tinggi: 16,
	})
	var x interface{} = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(x)
	x = "ganti dong"
	fmt.Println(x)
}
