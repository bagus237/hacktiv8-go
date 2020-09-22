package main

import ( 
		"net/http"
		"fmt"
		"text/template"
		"strings"
		"log"
		"sesi10/models"
		"io/ioutil"
		"encoding/json"
)

type MyMux struct{}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request){
	switch r.URL.Path {
	case "/":
		index(w, r)
	case "/register":
		register(w, r)
	default:
		http.NotFound(w, r)
	}
}

func index(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
	http.NotFound(w, r)
	return
	}
	var t, _ = template.ParseFiles("template/view.html")
	t.Execute(w, nil)
	return
}


func register(w http.ResponseWriter, r *http.Request){
	switch strings.ToUpper(r.Method){
	case "GET":
		var t, _ = template.ParseFiles("template/view.html")
		t.Execute(w, nil)
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil{
			fmt.Println("gagal:", err)
			return
		}
		regisReq := models.RegisterRequest{}
		if err := json.Unmarshal(body, &regisReq); err != nil{
			fmt.Println("false to parse data", err)
			return
		}
		 r.ParseForm()
		 fmt.Println("JSON NAMA:", regisReq.Nama)
		 fmt.Println("JSON Password:", regisReq.Password)
		 fmt.Println("JSON Verivikasi Password:", regisReq.Vpassword)
		 fmt.Println("JSON Email:", regisReq.Email)
		 fmt.Println("JSON JK:", regisReq.Gender)
		 fmt.Println("JSON Alamat:", regisReq.Alamat)
		 fmt.Println("JSON Agama:", regisReq.Agama)
		 http.Redirect(w, r, "/", http.StatusSeeOther) 
	default:
		http.NotFound(w, r)
		return
	}
}

func main(){
	fmt.Println("Server Run In Port :9000")
	configHTTP := &MyMux{}
	err := http.ListenAndServe(":9000", configHTTP)
	if err != nil {
		log.Fatal("Error running service: ", err)
	}
	
}
