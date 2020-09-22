package models
type (
	Feedback struct {
		ID uint `gorm:"primary_key AUTO_INCREMENT"`
		Name  string `json:"name"`
		Pesan string `json:"pesan"`
	}
	
	Feedbackedit struct {
		ID uint `json:"id"`
		Name string `json:"name"`
		Pesan string `json:"pesan"`
	}

)