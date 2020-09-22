package main
 
import (
    "fmt"
    "log"
    "net/http"
)


func index(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }
 
    switch r.Method {
    case "GET":     
         http.ServeFile(w, r, "view.html")
    case "POST":
        
        name := r.FormValue("name")
		password := r.FormValue("password")
		vpassword := r.FormValue("vpassword")
		email := r.FormValue("email")
		gender := r.FormValue("gender")
		address := r.FormValue("address")
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

func register(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }
 
func main() {
    http.HandleFunc("/", index)
 
    fmt.Printf("Server Running\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}