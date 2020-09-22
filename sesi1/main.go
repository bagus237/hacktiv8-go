package main
import {
	"fmt"
	"sesi1/config"

}

func main()  {
	person := config.Person{}
	person.Name = "bagus kaguya"
	person.Age = 10
	fmt.Println(person)
	fmt.Println(person.Name)
	fmt.Println(person.Age)
	person.Alamat.Kota = "Bogor"
}