package controller

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"final_project/structs"

	bcrypt "golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/go-sessions"
	
)

var db *sql.DB
var err error
var tpl *template.Template

func init() {
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1)/final_go")
	CheckErr(err)
	err = db.Ping()
	CheckErr(err)

	tpl = template.Must(template.ParseGlob("views/*"))
}


func CheckErr(err error) bool {
	if err != nil {
		log.Fatalln(err)
		return false
	}
	return true
}

func show(username string) structs.User {
	var users = structs.User{}
	err = db.QueryRow(`
		SELECT id, 
		username, 
		first_name, 
		last_name, 
		password 
		FROM users WHERE username=?
		`, username).
		Scan(
			&users.ID,
			&users.Username,
			&users.FirstName,
			&users.LastName,
			&users.Password,
		)
	return users
}

func Index(w http.ResponseWriter, r *http.Request) {
	_, err = template.ParseFiles("views/index.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rows, err := db.Query(
		`SELECT * FROM artikel WHERE status = "publish";
		`)
	CheckErr(err)

	artikels := make([]structs.Artikel, 0)
	for rows.Next() {
		art := structs.Artikel{}
		rows.Scan(&art.ID, &art.Name, &art.Isi, &art.Status)
		artikels = append(artikels, art)
	}
	log.Println(artikels)
	tpl.ExecuteTemplate(w, "index.html", artikels)

}

func Home(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	if len(session.GetString("username")) == 0 {
		session.Clear()
		sessions.Destroy(w, r)
		http.Redirect(w, r, "/login", 301)
	}
	var t, err = template.ParseFiles("views/home.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	t.Execute(w, r)
	return

}


func Contact(w http.ResponseWriter, r *http.Request) {

	session := sessions.Start(w, r)
	if len(session.GetString("username")) == 0 {
		session.Clear()
		sessions.Destroy(w, r)
		http.Redirect(w, r, "/login", 301)
	}

	_, err = template.ParseFiles("views/kontaklogin.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rows, err := db.Query(
		`SELECT * FROM pesan;
		`)
	CheckErr(err)

	ct := make([]structs.Pesan, 0)
	for rows.Next() {
		psn := structs.Pesan{}
		rows.Scan(&psn.ID, &psn.Nama, &psn.Email, &psn.Pesan)
		ct = append(ct, psn)
	}
	log.Println(ct)
	tpl.ExecuteTemplate(w, "kontaklogin.html", ct)

}


func Article(w http.ResponseWriter, r *http.Request) {

	session := sessions.Start(w, r)
	if len(session.GetString("username")) == 0 {
		session.Clear()
		sessions.Destroy(w, r)
		http.Redirect(w, r, "/login", 301)
	}

	_, err = template.ParseFiles("views/article.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rows, err := db.Query(
		`SELECT * FROM artikel;
		`)
	CheckErr(err)

	artikels := make([]structs.Artikel, 0)
	for rows.Next() {
		art := structs.Artikel{}
		rows.Scan(&art.ID, &art.Name, &art.Isi, &art.Status)
		artikels = append(artikels, art)
	}
	log.Println(artikels)
	tpl.ExecuteTemplate(w, "article.html", artikels)

}

func About(w http.ResponseWriter, r *http.Request) {

	var t, err = template.ParseFiles("views/about.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	t.Execute(w, r)
	return

}

func Kontak(w http.ResponseWriter, r *http.Request) {
	var t, err = template.ParseFiles("views/kontak.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	t.Execute(w, r)
	return

}

func Aboutlogin(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	if len(session.GetString("username")) == 0 {
		session.Clear()
		sessions.Destroy(w, r)
		http.Redirect(w, r, "/login", 301)
	}

	var t, err = template.ParseFiles("views/aboutlogin.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	t.Execute(w, r)
	return

}


func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "views/register.html")
		return
	}

	username := r.FormValue("email")
	first_name := r.FormValue("first_name")
	last_name := r.FormValue("last_name")
	password := r.FormValue("password")

	users := show(username)

	if (structs.User{}) == users {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		if len(hashedPassword) != 0 && CheckErr(err) {
			stmt, err := db.Prepare("INSERT INTO users SET username=?, password=?, first_name=?, last_name=?")
			if err == nil {
				_, err := stmt.Exec(&username, &hashedPassword, &first_name, &last_name)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
		}
	} else {
		http.Redirect(w, r, "/register", 302)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "views/login.html")
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	users := show(username)

	//deskripsi dan compare password
	var passwordTes = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))

	if passwordTes == nil {
		//login success
		session := sessions.Start(w, r)
		session.Set("username", users.Username)
		session.Set("name", users.FirstName)
		http.Redirect(w, r, "/home", 302)
	} else {
		//login failed
		http.Redirect(w, r, "/", 302)

	}

}
func Logout(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	session.Clear()
	sessions.Destroy(w, r)
	http.Redirect(w, r, "/", 302)
}

func ArtikelForm(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	if len(session.GetString("username")) == 0 {
		sessions.DestroyAll()
		http.Redirect(w, r, "/login", 301)
	}

	err = tpl.ExecuteTemplate(w, "artikelForm.html", nil)
	CheckErr(err)
}

func CreateArtikel(w http.ResponseWriter, req *http.Request) {
	session := sessions.Start(w, req)
	if len(session.GetString("username")) == 0 {
		session.Clear()
		sessions.Destroy(w, req)
		http.Redirect(w, req, "/login", 301)
	}

	if req.Method == http.MethodPost {
		art := structs.Artikel{}
		art.Name = req.FormValue("name")
		art.Isi = req.FormValue("isi")
		art.Status = req.FormValue("status")
		
		_, err = db.Exec(
			"INSERT INTO artikel (name, isi, status) VALUES (?, ?, ?)",
			art.Name,
			art.Isi,
			art.Status,
		)
		CheckErr(err)
		http.Redirect(w, req, "/artikel", http.StatusSeeOther)
		return
	}
	http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
}

func CreatePesan(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		psn := structs.Pesan{}
		psn.Nama = req.FormValue("nama")
		psn.Email = req.FormValue("email")
		psn.Pesan = req.FormValue("pesan")
		
		_, err = db.Exec(
			"INSERT INTO pesan (nama, email, pesan) VALUES (?, ?, ?)",
			psn.Nama,
			psn.Email,
			psn.Pesan,
		)
		CheckErr(err)
		http.Redirect(w, req, "/contact", http.StatusSeeOther)
		return
	}
	http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
}

func EditArtikel(w http.ResponseWriter, req *http.Request) {
	session := sessions.Start(w, req)
	if len(session.GetString("username")) == 0 {
		session.Clear()
		sessions.Destroy(w, req)
		http.Redirect(w, req, "/login", 301)
	}

	id := req.FormValue("id")
	rows, err := db.Query(
		`SELECT id,
		 	name,
			isi,
			status
		FROM artikel
		WHERE id = ` + id + `;`)
	CheckErr(err)
	art := structs.Artikel{}
	for rows.Next() {
		rows.Scan(&art.ID, &art.Name, &art.Isi, &art.Status)
	}
	tpl.ExecuteTemplate(w, "editArtikel.html", art)
}

func DeleteArtikel(w http.ResponseWriter, req *http.Request) {
	session := sessions.Start(w, req)
	if len(session.GetString("username")) == 0 {
		session.Clear()
		sessions.Destroy(w, req)
		http.Redirect(w, req, "/login", 301)
	}

	id := req.FormValue("id")
	if id == "" {
		http.Error(w, "Please send ID", http.StatusBadRequest)
	}
	_, err := db.Exec("DELETE FROM artikel WHERE id = ?", id)
	CheckErr(err)
	http.Redirect(w, req, "/artikel", http.StatusSeeOther)
}

func DeletePesan(w http.ResponseWriter, req *http.Request) {
	session := sessions.Start(w, req)
	if len(session.GetString("username")) == 0 {
		session.Clear()
		sessions.Destroy(w, req)
		http.Redirect(w, req, "/login", 301)
	}

	id := req.FormValue("id")
	if id == "" {
		http.Error(w, "Please send ID", http.StatusBadRequest)
	}
	_, err := db.Exec("DELETE FROM pesan WHERE id = ?", id)
	CheckErr(err)
	http.Redirect(w, req, "/contactLogin", http.StatusSeeOther)
}


func UpdateArtikel(w http.ResponseWriter, req *http.Request) {
	session := sessions.Start(w, req)
	if len(session.GetString("username")) == 0 {
		session.Clear()
		sessions.Destroy(w, req)
		http.Redirect(w, req, "/login", 301)
	}
	_, err := db.Exec(
		"UPDATE artikel SET name = ?, isi = ?, status = ? WHERE id = ? ",
		req.FormValue("name"),
		req.FormValue("isi"),
		req.FormValue("status"),
		req.FormValue("id"),
	)
	CheckErr(err)
	http.Redirect(w, req, "/artikel", http.StatusSeeOther)
}
