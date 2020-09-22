package main

import (
	"fmt"
	"log"
	"net/http"
)
const index=`<html> 
<head>
<title> Dashboard- Golang </title>
</head>
<body>
<h1>Dashboard</h1>
</body>
</html>`
func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"it's Works")

}

func dashboard(w http.ResponseWriter, r *http.Request){
	templ, err := template.New("dashboard").Parse(templateHtml)
	if err !:= nil{
		fmt.Println("Filed to add template: ", err)	
	}
	err = templ.Execute(w, nil)
	if err != nil {
		fmt.Println("failed to execute html:", err)
	}
}

func main(){
	http.HandleFunc("/", index)
	http.HandleFunc("/dashboard", dashboard)
	fmt.Println("Web Server Running")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("Error running service: ", err)
	}
}


