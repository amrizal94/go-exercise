package main

import (
	"fmt"
	"math"
)

type Shape interface {
	luas() float32
	keliling() float32
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
	alas, tinggi, sisiMiring float64
}

func (s segitiga) luas() float32 {
	return float32(s.alas * s.tinggi * 1 / 2)
}

func (s *segitiga) keliling() float32 {
	s.sisiMiring = math.Sqrt(s.alas*s.alas + s.tinggi*s.tinggi)

	return float32(s.alas + (s.sisiMiring * 2))
}

func (pP persegiPanjang) luas() float32 {
	return pP.panjang * pP.lebar
}

func (pP persegiPanjang) keliling() float32 {
	return 2 * (pP.panjang + pP.lebar)
}

func (p persegi) luas() float32 {
	return p.sisi * p.sisi
}

func (p persegi) keliling() float32 {
	return 4 * p.sisi
}

func (l lingkaran) luas() float32 {
	return math.Pi * l.jari2 * l.jari2
}

func (l lingkaran) keliling() float32 {
	return 2 * math.Pi * l.jari2
}

var hitungLuas = func(s Shape) float32 {
	return s.luas()
}

var hitungKeliling = func(s Shape) float32 {
	return s.keliling()
}

func main() {

	l := lingkaran{jari2: 5}
	fmt.Printf("Diket: jari²: %v cm, maka luas: %v cm² dan keliling lingkaran: %v cm\n", hitungLuas(l), l.luas(), l.keliling())

	pPanjang := persegiPanjang{
		panjang: 17,
		lebar:   7,
	}
	fmt.Printf("Diket: panjang: %v cm dan lebar: %v cm, maka luas: %v cm² dan keliling persegi panjang: %v cm\n", pPanjang.panjang, pPanjang.lebar, pPanjang.luas(), pPanjang.keliling())
	p := persegi{sisi: 7}
	fmt.Printf("Diket: sisi: %v cm, maka luas persegi: %v cm²\n", p.sisi, p.luas())
	s := &segitiga{
		alas:   8,
		tinggi: 6,
	}
	fmt.Printf("Diket: alas: %v cm dan tinggi: %v cm, maka luas segitiga: %v cm²\n", s.alas, s.tinggi, s.luas())
	fmt.Printf("Dicari sisi miring mengunakan pytagoras: %v cm, maka keliling segitiga: %v cm\n", s.sisiMiring, hitungKeliling(s))

	var x interface{} = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(x)
	x = "ganti dong"
	fmt.Println(x)
}
