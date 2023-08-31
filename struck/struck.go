package main

import "fmt"

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

func main() {
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
}
