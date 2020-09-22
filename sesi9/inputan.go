package main
 
import (
    "fmt"
    "log"
    "html/template"
    "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.NotFound(w, r)
		return
	}
	var t, _ = template.ParseFiles("template/view.html")
	t.Execute(w, nil)
	return
}

func register(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
		http.ServeFile(w, r, "template/view.html")
		return
	}
 
    switch r.Method {
    case "GET":     
         http.ServeFile(w, r, "template/view.html")
    case "POST":
        
        name := r.FormValue("name")
		password := r.FormValue("password")
		vpassword := r.FormValue("vpassword")
		email := r.FormValue("email")
		gender := r.FormValue("gender")
		address := r.FormValue("alamat")
		agama := r.FormValue("agama")
		
        fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Password = %s\n", password)
		fmt.Fprintf(w, "Verifikasi Password = %s\n", vpassword)
		fmt.Fprintf(w, "Email = %s\n", email)
		fmt.Fprintf(w, "Gender = %s\n", gender)
		fmt.Fprintf(w, "Alamat = %s\n", address)
		fmt.Fprintf(w, "Agama = %s\n", agama)
    default:
        fmt.Fprintf(w, " System Error")
    }
}
 
func main() {
    http.HandleFunc("/", index)
    http.HandleFunc("/register", register)
    fmt.Printf("Server Running\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}