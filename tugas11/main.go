package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"final_project/controller"


	_ "github.com/go-sql-driver/mysql"
	
	
)

var db *sql.DB
var err error
var tpl *template.Template

func init() {
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1)/final_go")
	controller.CheckErr(err)
	err = db.Ping()
	controller.CheckErr(err)

	tpl = template.Must(template.ParseGlob("views/*"))
}

func routes() {
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/home", controller.Home)
	http.HandleFunc("/artikel", controller.Article)
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/artikelForm", controller.ArtikelForm)
	http.HandleFunc("/createPesan", controller.CreatePesan)
	http.HandleFunc("/editArtikel", controller.EditArtikel)
	http.HandleFunc("/deleteArtikel", controller.DeleteArtikel)
	http.HandleFunc("/deletePesan", controller.DeletePesan)
	http.HandleFunc("/updateArtikel", controller.UpdateArtikel)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)
}

func main() {
	
	routes()

	defer db.Close()

	fmt.Println("Server running on port :8000")
	http.ListenAndServe(":8000", nil)
}
