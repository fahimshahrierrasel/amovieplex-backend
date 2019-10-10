package genre

import (
	"amovieplex-backend/src/api/helpers"
	"amovieplex-backend/src/data/db"
	"amovieplex-backend/src/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Genre = models.Genre

type RequestBody struct {
	Name string `json:"name" binding:"required"`
}

func create(ctx *gin.Context) {
	data := map[string]interface{}{}
	var requestBody RequestBody
	if err := ctx.BindJSON(&requestBody); err != nil {
		log.Printf("Request Body: %v", requestBody)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.MakeResponse(data, true, "request body is not correct"))
		return
	}
	newGenre := Genre{Name: requestBody.Name, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}}
	ok := db.CreateGenre(ctx, newGenre)
	if ok {
		data["status"] = "Genre Successfully Created"
	} else {
		data["status"] = "Sorry!!, Genre Not Created Unwanted Behaviour"
	}
	ctx.JSON(http.StatusOK, helpers.MakeResponse(data, !ok, ""))
}
