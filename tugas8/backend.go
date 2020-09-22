//Bismillahirrohmanirrohim
package main
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"html/template"
	"net/http"
	"math/rand"
)

type weather struct{
	Status weathercondition `json:"Status"`
}

type weathercondition struct {
	Water int `json:"Water"`
	Wind int `json:"Wind"`
}


func Editjson() {

	// declare random value and make json structure
	windspeed := rand.Intn(100)
	waterlevel := rand.Intn(100)
	data := weather{
		Status: weathercondition{Wind: windspeed,
			Water: waterlevel},
	}

	//marshal the file
	file, _ := json.MarshalIndent(data, "", " ")

	//write the file
	_ = ioutil.WriteFile("data.json", file, 0644)
}

func index(w http.ResponseWriter, r *http.Request) {

	
	// run html file
	t, err := template.ParseFiles("home.html")
	if err != nil {
		panic(err)
	}
	err = t.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		panic(err)
	}

	//baca file
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Print(err)
	}
	// json data
	var obj weather

	//unmarshall 
	err = json.Unmarshal(data, &obj)
	if err != nil{
		fmt.Println("error:", err)
	}
	
	Editjson()
	fmt.Fprintln(w, "The Weather is :")
// can access using struct now
	fmt.Fprintln(w, "wind :", obj.Status.Wind, "kmph")
	fmt.Fprintln(w, "water :", obj.Status.Water, "m")

// condition classification
	if obj.Status.Wind <= 6 {
	fmt.Fprintln(w, "Wind :aman")
	}
	if obj.Status.Wind >= 7 && obj.Status.Wind <= 15 {
	fmt.Fprintln(w, "Wind :status siaga")
	}
	if obj.Status.Wind > 15 {
	fmt.Fprintln(w, "Wind :bahaya")
	}

	if obj.Status.Water <= 5 {
	fmt.Fprintln(w, "Water :aman")
	}
	if obj.Status.Water >= 6 && obj.Status.Water <= 8 {
	fmt.Fprintln(w, "Water :status siaga")
	}
	if obj.Status.Water > 8 {
	fmt.Fprintln(w, "Water :bahaya")
	}
}

func main(){
	http.HandleFunc("/", index)
	fmt.Println("Web Server Running")
	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		log.Fatal("Error running service: ", err)
	}
} 

