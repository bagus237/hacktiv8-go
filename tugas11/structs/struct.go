package structs

type Artikel struct {
	ID   int
	Name string
	Isi  string
	Status string
}

type User struct {
	ID        int
	Username  string
	FirstName string
	LastName  string
	Password  string
}

type Pesan struct {
	ID    int
	Nama  string
	Email string
	Pesan string
}

