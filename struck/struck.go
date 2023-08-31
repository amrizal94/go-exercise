package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/go-playground/locales/en_US"
)

func main() {
	l := en_US.New()
	byson := motor{
		cc:    150,
		merek: "Yamaha Byson",
		banDepan: ban{
			lebar: 80,
		},
	}
	var myMotor []motor
	byson.banBelakang.lebar = 90
	byson.banBelakang.aspectRatio = 100
	myMotor = append(myMotor, byson)
	beat := motor{
		cc:    110,
		merek: "Honda Beat",
	}
	beat.banDepan.lebar = 70
	beat.banBelakang.lebar = 70
	myMotor = append(myMotor, beat)
	for _, m := range myMotor {
		pajak(&m)
		fmt.Printf("motor %s kenak pajak sebesar %v\n", m.merek, m.pajak)
		switch {
		case cekBan(m):
			fmt.Printf("ban depan %v cocok dengan ban belakang yg lebih besar\n", m.merek)
		default:
			fmt.Printf("ganti ban belakang %v dengan lebih besar\n", m.merek)
		}
	}
	hpku := hp{
		merek: "apple",
		speck: speck{
			processor: "M1",
			gpu:       "7-core",
		},
		harga: 1000000,
	}

	fmt.Println(hpku)
	fmt.Printf("->Detail HPku<-\nmerek: %v\nprocessor: %v\ngpu: %v\nharga: Rp. %v\n", hpku.merek, hpku.processor, hpku.gpu, l.FmtNumber(float64(hpku.harga), 2))
	myIpad := struct {
		version string
		harga   int
	}{
		version: "Ipad Air 5",
		harga:   5000000,
	}
	fmt.Println(myIpad)
	s := fmt.Sprintf("%.2f", 11211.12112)
	log.Println(s)
	f := strconv.Quote(`"Fran & Freddie's Diner	☺"`)
	log.Println(f)
	var f64 = 78.12

	// Percent
	fmt.Println(l.FmtPercent(f64, 2))
	// Number
	fmt.Println(l.FmtNumber(float64(hpku.harga), 2))

}

type motor struct {
	merek       string
	cc          int
	pajak       int
	banDepan    ban
	banBelakang ban
}

type ban struct {
	lebar       byte
	aspectRatio byte
}

var pajak = func(kendaraan *motor) {
	switch {
	case kendaraan.cc < 150:
		kendaraan.pajak = 1000
	case kendaraan.cc <= 200:
		kendaraan.pajak = 2000
	default:
		kendaraan.pajak = 5000
	}

}

var cekBan = func(k motor) (isProposional bool) {
	if k.banDepan.lebar < k.banBelakang.lebar {
		isProposional = true
	}
	return
}

type speck struct {
	processor string
	gpu       string
}

type hp struct {
	speck
	merek string
	harga int
}
