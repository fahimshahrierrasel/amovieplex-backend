package rating

import (
	"amovieplex-backend/src/api/helpers"
	"amovieplex-backend/src/common"
	"amovieplex-backend/src/common/errors"
	"amovieplex-backend/src/data/db"
	"amovieplex-backend/src/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Rating = models.Rating

type AddRequestBody struct {
	Name     string `json:"name" binding:"required"`
	AgeLimit int    `json:"age_limit" binding:"required"`
}

type UpdateRequestBody struct {
	Name     string `json:"name"`
	AgeLimit int    `json:"age_limit"`
}

func create(ctx *gin.Context) {
	data := map[string]interface{}{}
	var requestBody AddRequestBody
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
	serializer := Serializer{result}
	ctx.JSON(http.StatusOK, helpers.MakeResponse(serializer.Response(), false, ""))
}

func getRating(ctx *gin.Context) {
	ratingID := ctx.Param("rating_id")

	rating, err := db.GetRating(ctx, ratingID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.MakeResponse(common.JSON{}, true, errors.ErrorCodeMessage(errors.ERRNotFound)))
		return
	}

	serializer := SimpleSerializer{rating}

	ctx.JSON(http.StatusOK,
		helpers.MakeResponse(serializer.Response(), false, ""))
}

func updateRating(ctx *gin.Context) {
	ratingID := ctx.Param("rating_id")
	data := common.JSON{}
	var requestBody UpdateRequestBody

	// checking request body and assigning value
	if err := ctx.BindJSON(&requestBody); err != nil {
		log.Printf("Request Body: %v", requestBody)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.MakeResponse(common.JSON{}, true, errors.ErrorCodeMessage(errors.ERRReqBody)))
		return
	}

	// check if the object exists in the db
	rating, err := db.GetRating(ctx, ratingID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound,
			helpers.MakeResponse(common.JSON{}, true, errors.ErrorCodeMessage(errors.ERRNotFound)))
		return
	}

	// replacing current object by updated value from the request body if
	// the request body has any value
	rating.Name = common.SetActualValueFrom(requestBody.Name, rating.Name).(string)
	rating.AgeLimit = common.SetActualValueFrom(requestBody.AgeLimit, rating.AgeLimit).(int)
	rating.UpdatedAt = time.Now()

	// updating the object
	ok := db.UpdateRating(ctx, ratingID, rating)
	if ok {
		data["status"] = "Rating Successfully Updated"
	} else {
		data["status"] = "Sorry!!, Rating Not Updated Unwanted Behaviour"
	}
	ctx.JSON(http.StatusOK, helpers.MakeResponse(data, !ok, ""))
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
