package controller

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"sesi12/models"
	
)


var db *gorm.DB
//Koneksi ke MySQL
func init() {
	var err error
	db, err =
		gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Gagal Conect Ke Database")
	}
	db.AutoMigrate(&models.Feedback{})
}

func CretedFeedback(c *gin.Context) {
	var std models.Feedbackedit
	var model models.Feedback
	c.Bind(&std)
	validasi := ValidatorCreated(std)
	model = TransferEPToModel(std)
	if validasi != "" {
		c.JSON(http.StatusOK, gin.H{"message": http.StatusOK, "result": validasi})
	} else {
		db.Create(&model)
		c.JSON(http.StatusOK, gin.H{"message": http.StatusOK, "result": model})
	}
}

func TransferEPToModel(FE models.Feedbackedit) models.Feedback {
	var model models.Feedback
	model = models.Feedback{
		Name:        FE.Name,
		Pesan:      FE.Pesan,
		
	}
	return model
}

func transferModelToEP(model models.Feedback) models.Feedbackedit {
	var FE models.Feedbackedit
	FE = models.Feedbackedit{
		ID:          model.ID,
		Name:        model.Name,
		Pesan: 	 model.Pesan,	
	}
	return FE
}

// Untuk Validasi data not null
func ValidatorCreated(FE models.Feedbackedit) string {

	var kosong string = " Is Empty"

	if FE.Name == "" {
		return "name" + kosong
	}

	if FE.Pesan == "" {
		return "pesan" + kosong
	}

	return ""
}


func fetchAllPerson(c *gin.Context) {
	var model [] models.Feedback
	var EP [] models.Feedbackedit

	db.Find(&model)

	if len(model) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": http.StatusNotFound, "result": "Data Tidak Ada"})
	}

	for _, item := range model {
		EP = append(EP, transferModelToEP(item))
	}
	c.JSON(http.StatusOK, gin.H{"message": http.StatusOK, "result": EP})
}