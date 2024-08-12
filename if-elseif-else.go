package main

import "fmt"

func IfElseIfElse() {
	kodeKota := "CGK"

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

	fmt.Println(namaKota)
}
