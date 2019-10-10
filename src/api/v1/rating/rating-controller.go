package rating

import (
	"amovieplex-backend/src/api/helpers"
	"amovieplex-backend/src/data/db"
	"amovieplex-backend/src/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
		log.Printf("Request Body: %v", requestBody)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.MakeResponse(data, true, "request body is not correct"))
		return
	}
	newRating := Rating{Name: requestBody.Name, AgeLimit: requestBody.AgeLimit,
		CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}}
	ok := db.CreateRating(ctx, newRating)
	if ok {
		data["status"] = "Rating Successfully Created"
	} else {
		data["status"] = "Sorry!!, Rating Not Created Unwanted Behaviour"
	}
	ctx.JSON(http.StatusOK, helpers.MakeResponse(data, !ok, ""))
}

func getAll(ctx *gin.Context) {
	result := db.GetAllRating(ctx)
	serializer := RatingSerializer{result}
	ctx.JSON(http.StatusOK, helpers.MakeResponse(serializer.Response(), false, ""))
}

func softDelete(ctx *gin.Context) {
	ratingID := ctx.Param("rating_id")
	data := map[string]interface{}{}
	ok := db.SoftDeleteRating(ctx, ratingID)
	if ok {
		data["status"] = "Rating Successfully Soft Deleted"
	} else {
		data["status"] = "Sorry!!, Rating Not Soft Deleted Unwanted Behaviour"
	}
	ctx.JSON(http.StatusOK, helpers.MakeResponse(data, !ok, ""))
}

func permanentDelete(ctx *gin.Context) {
	ratingID := ctx.Param("rating_id")
	data := map[string]interface{}{}
	ok := db.DeleteRating(ctx, ratingID)
	if ok {
		data["status"] = "Rating Successfully Deleted"
	} else {
		data["status"] = "Sorry!!, Rating Not Deleted Unwanted Behaviour"
	}
	ctx.JSON(http.StatusOK, helpers.MakeResponse(data, !ok, ""))
}
