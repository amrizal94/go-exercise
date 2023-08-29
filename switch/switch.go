package main

import (
	"fmt"
	"math/rand"
)

var seminggu = map[int]string{
	1: "Senin",
	2: "Selasa",
	3: "Rabu",
	4: "Kamis",
	5: `Jum'at`,
	6: "Sabtu",
	7: "Minggu",
}

const (
	Senin  = 1
	Selasa = 2
	Rabu   = 3
	Kamis  = 4
	Jumat  = 5
	Sabtu  = 6
	Minggu = 7
)

func main() {
	hari := rand.Intn(len(seminggu)) + 1

	switch hari {
	case Senin, Rabu:
		fmt.Println("Latihan Badminton pada hari", seminggu[hari])
	case Selasa, Kamis:
		fmt.Println("Latihan Sepak Bola pada hari", seminggu[hari])
	case Jumat, Sabtu:
		fmt.Println("Latihan Futsal pada hari", seminggu[hari])
	case Minggu:
		fmt.Println("Latihan Renang pada hari", seminggu[hari])
	default:
		fmt.Println("ini Hari apa?", seminggu[hari])
	}

	switch  {
	case hari <= Selasa:
		fmt.Printf("Pada hari %v kita belajar nodeJs\n", seminggu[hari])
	case hari <= Kamis:
		fmt.Printf("Sekarang hari %v waktunya kita belajar database \n", seminggu[hari])
	default:
		fmt.Println("lanjut nanti untuk hari", seminggu[hari])

	}

}
