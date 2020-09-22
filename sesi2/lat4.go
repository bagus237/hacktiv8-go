package main

import (
	"fmt"
	"os"
)

type Profile struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func mainanan () {

	argsRaw := os.Args
	args := argsRaw[1:]
	generateProfile(args)
}

func generateProfile(data []string) {
	profile := Profile{
		Nama:      data[0],
		Alamat:    data[1],
		Pekerjaan: data[2],
		Alasan:    data[3],
	}
	fmt.Println("Nama: " + profile.Nama)
	fmt.Println("Alamat: " + profile.Alamat)
	fmt.Println("Pekerjaan: " + profile.Pekerjaan)
	fmt.Println("Alasan: " + profile.Alasan)

}
