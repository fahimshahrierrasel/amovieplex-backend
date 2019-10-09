package rating

import (
	"amovieplex-backend/src/api/helpers"
	"amovieplex-backend/src/data/db"
	"amovieplex-backend/src/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Rating = models.Rating

type RequestBody struct {
	Name     string `json:"name" binding:"required"`
	AgeLimit int    `json:"age_limit" binding:"required"`
}

func create(ctx *gin.Context) {
	data := map[string]interface{}{}
	var requestBody RequestBody
	if err := ctx.BindJSON(&requestBody); err != nil {
		fmt.Println(requestBody)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.MakeResponse(data, true, "request body is not correct"))
		return
	}
	newRating := Rating{Name: requestBody.Name, AgeLimit: requestBody.AgeLimit,
		CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}}
	ok := db.CreateRating(ctx, newRating)
	if ok {
		data["status"] = "Rating Successfully Created"
	}else{
		data["status"] = "Sorry!!, Rating Not Created Unwanted Behaviour"
	}
	ctx.JSON(http.StatusOK, helpers.MakeResponse(data, !ok, ""))
}

func getAll(ctx *gin.Context) {
	data := map[string]interface{}{}
	db.GetAllRating(ctx)
	ctx.JSON(http.StatusOK, helpers.MakeResponse(data, true, ""))
}
