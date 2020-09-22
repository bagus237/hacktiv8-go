package main

import ("fmt")

type BangunDatar interface{
	Luas()
	Keliling()
}

type Lingkaran struct{
	Diameter float64
	JariJari float64
}

func (l Lingkaran) Luas(){
	fmt.Print("Luas")
}

func (l Lingkaran ) Keliling(){
	fmt.Println("keliling")
}

func main(){
	l := Lingkaran{
		Diameter: 14,
		JariJari: 7,
	}

var BangunRuang BangunDatar
BangunRuang = l

var data interface{}
data = 5
fmt.Println(data.(int))
var ruang = BangunRuang.(Lingkaran)
fmt.Println(ruang.JariJari)
fmt.Println(ruang.Diameter)

}