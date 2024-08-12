package main

import "fmt"

func ForRange() {
	cityCodes := []string{"CGK", "BDO", "DPS", "XYZ"}

	for _, kodeKota := range cityCodes {
		var namaKota string

		if kodeKota == "CGK" {
			namaKota = "Jakarta"
		} else if kodeKota == "BDO" {
			namaKota = "Bandung"
		} else if kodeKota == "DPS" {
			namaKota = "Denpasar"
		} else {
			namaKota = "Kode kota tidak valid"
		}

		fmt.Printf("Kode Kota: %s, Nama Kota: %s\n", kodeKota, namaKota)
	}
}
