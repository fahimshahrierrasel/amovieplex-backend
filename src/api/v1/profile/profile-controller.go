package profile

import (
	"amovieplex-backend/src/api/helpers"
	"amovieplex-backend/src/common"
	"amovieplex-backend/src/common/errors"
	"amovieplex-backend/src/data/db"
	"amovieplex-backend/src/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"
)

type Profile = models.Profile

type NewRequestBody struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	DOB       string `json:"dob"`
	User      string `json:"user"  binding:"required"`
}

func create(ctx *gin.Context) {
	data := common.JSON{}
	var requestBody NewRequestBody
	if err := ctx.BindJSON(&requestBody); err != nil {
		log.Printf("Request Body: %v", requestBody)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.MakeResponse(data, true, "request body is not correct"))
		return
	}

	userID, err := primitive.ObjectIDFromHex(requestBody.User)
	if err != nil {
		log.Printf("Error Making ObjectID from String: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.MakeResponse(data, true, errors.ErrorCodeMessage(errors.ERRReqBody)))
		return
	}

	var dob = time.Time{}
	if requestBody.DOB != "" {
		dob, err = time.Parse(common.TimeFormatLayout, requestBody.DOB)
		if err != nil {
			log.Printf("Error Parsing time from String: %v", err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				helpers.MakeResponse(data, true, errors.ErrorCodeMessage(errors.ERRReqBody)))
			return
		}
	}

	newProfile := Profile{
		FirstName: requestBody.FirstName,
		LastName:  requestBody.LastName,
		DOB:       dob,
		User:      userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Time{},
	}

	profileID, err := db.CreateProfile(ctx, newProfile)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.MakeResponse(data, true, err.Error()))
		return
	}
	data["_id"] = profileID
	ctx.JSON(http.StatusOK,
		helpers.MakeResponse(data, false, ""))
}

func getProfileByUserID(ctx *gin.Context) {
	userID := ctx.Param("user_id")

	profile, err := db.GetProfile(ctx, userID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound,
			helpers.MakeResponse(common.JSON{}, true, errors.ErrorCodeMessage(errors.ERRNotFound)))
		return
	}

	serializer := ProfileSerializer{profile}

	ctx.JSON(http.StatusOK,
		helpers.MakeResponse(serializer.Response(), false, ""))
}
