package main

import "fmt"

type BangunDatar interface {
	Luas()
	Keliling()
}

type Lingkaran struct {
	Diameter float32
	JariJari float32
}

func (l Lingkaran) Luas() {
	fmt.Println("Luas")
}
func (l Lingkaran) Keliling() {
	fmt.Println("Keliling")
}

func main() {
	l := Lingkaran{
		Diameter: 14,
		JariJari: 7,
	}

	var bangunRuang BangunDatar
	bangunRuang = l

	var data interface{}
	data = 5
	fmt.Println(data.(int))

	bangunRuang.Luas()
	bangunRuang.Keliling()
	var ruang = bangunRuang.(Lingkaran)
	fmt.Println(ruang.JariJari)
	fmt.Println(ruang.Diameter)
}