package main

import ("fmt")

type Lingkaran struct {
	Diameter float64
}
const (phi = 3.14)
func (l *Lingkaran) SetDiameter(Diameter float64){
	l.Diameter = Diameter
}

func (j *Lingkaran) hitungJariJari() float64{
	r := j.Diameter/2
	return r
}

func (s *Lingkaran) hitungLuas() float64{
	luas := phi * (s.hitungJariJari() * s.hitungJariJari())
	return luas
}

func (k *Lingkaran) hitungKeliling() float64{
	keliling := 2 * phi * k.hitungJariJari()
	return keliling
}

func main() {
	o := Lingkaran{}
	o.SetDiameter(7)

	fmt.Println("Diameter:", o.Diameter)
	fmt.Println("Jari-jari:", o.hitungJariJari())
	fmt.Println("Luas:", o.hitungLuas())
	fmt.Println("Keliling:",o.hitungKeliling())
}