package controllers

import (
	"context"
	"downloadFile/db"
	"downloadFile/model"
	"downloadFile/pdf"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Add ajoute un élément à la base de donnée
func Add(c *gin.Context) {
	collection, cancel := db.ConnectToDB()
	defer cancel()

	//On récupère la data envoyé par notre app react
	data, err := c.GetRawData()
	if err != nil {
		log.Fatalf("Error : %v\n", err)
	}

	//On parse notre data de type []bytes pour qu'elle convienne à notre struct
	var doc model.Facture
	err = json.Unmarshal(data, &doc)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
	}

	//On l'enrengistre dans notre base de donnée
	_, err = collection.InsertOne(context.Background(), doc)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
	}

	pdf.GenPdf(doc)

	//On renvoie un message de succès
	c.JSON(http.StatusOK, gin.H{
		"msg": "./output/" + doc.Username + "-" + doc.Birthday + ".pdf",
	})
}
