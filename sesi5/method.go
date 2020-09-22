package main

import "fmt"
type Interior struct {
	Kursi int
	Stir  bool
}

type Mobil struct {
	Merk  string
	Warna string
	Ban   int
	Kaca  int
	Interior
}

func (inter Interior) PrintJumlahKursi() {
	fmt.Println(inter.Kursi)
}

func main() {
	p := Mobil{
		Interior: Interior{
			Kursi: 4,
			Stir:  true,
		},
	}
	p.Kursi = 2
	p.Stir = false

	p.PrintJumlahKursi()
}
