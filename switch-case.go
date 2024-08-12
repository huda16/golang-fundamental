package main

import "fmt"

func SwitchCase() {
	kodeKota := "CGK"

	var namaKota string
	switch kodeKota {
	case "CGK":
		namaKota = "Jakarta"
	case "BDO":
		namaKota = "Bandung"
	case "DPS":
		namaKota = "Denpasar"
	default:
		namaKota = "Kode kota tidak valid"
	}

	fmt.Println(namaKota)
}
