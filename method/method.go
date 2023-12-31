package main

import "fmt"

type persegi struct {
	lebar, panjang float32
}

func (p persegi) luas() float32 {
	p.panjang = 12
	return p.lebar * p.panjang
}

func (p *persegi) pointerLuas() float32 {
	p.panjang = 11
	return p.lebar * p.panjang
}
func main() {
	kotak := persegi{
		lebar:   10.5,
		panjang: 12.3,
	}
	s := fmt.Sprintf(`Diketahui:
lebar = %v
panjang = %v
maka luasnya adalah %.2f`, kotak.lebar, kotak.panjang, kotak.luas())
	fmt.Println(s)

	f := fmt.Sprintf(`Diketahui:
lebar = %v
panjang = %v
maka luasnya adalah %.2f`, kotak.lebar, kotak.panjang, kotak.pointerLuas())
	fmt.Println(f)
}
