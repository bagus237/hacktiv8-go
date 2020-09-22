package main

import ( 
		"net/http"
		"fmt"
		"text/template"
		"strings"
		"log"
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
		 r.ParseForm()
		 fmt.Println("Nama -->>", r.Form["name"])
		 fmt.Println("Password -->>", r.Form["password"])
		 fmt.Println("Verifikasin Password -->>", r.Form["vpassword"])
		 fmt.Println("Email-->>", r.Form["email"]) 
		 fmt.Println("Jenis Kelamin-->>", r.Form["gender"]) 
		 fmt.Println("Alamat-->>", r.Form["alamat"])
		 fmt.Println("Agama -->>", r.Form["agama"]) 
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
