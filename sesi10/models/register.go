package models

type RegisterRequest struct{
	Nama string `json:"nama"`
	Password string `json:"password"`
	Vpassword string `json:"vpassword"`
	Email string `json:"email"`	
	Gender string `json:"gender"`
	Alamat string `json:"alamat"`
	Agama string `json:"agama"`
}