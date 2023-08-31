package main

import "fmt"

type persegi struct {
	lebar   float32
	panjang float32
}

func (p persegi) luas() float32 {
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
}
