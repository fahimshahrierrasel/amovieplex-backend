package auth

import (
	"amovieplex-backend/src/api/helpers"
	"amovieplex-backend/src/common"
	"amovieplex-backend/src/common/errors"
	"amovieplex-backend/src/data/db"
	"amovieplex-backend/src/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type User = models.User

type SignUpRequestBody struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password" binding:"required"`
	Role        string `json:"role" binding:"required"`
}

type LoginRequestBody struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password" binding:"required"`
}

func register(ctx *gin.Context) {
	data := common.JSON{}
	var requestBody SignUpRequestBody

	if err := ctx.BindJSON(&requestBody); err != nil {
		log.Printf("Request Body: %v, Error: %v", requestBody, err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.MakeResponse(data, true, errors.ErrorCodeMessage(errors.ERRReqBody)))
		return
	}

	hashedPassword, err := helpers.MakePasswordHash(requestBody.Password)
	if err != nil {
		log.Printf("Error Making Password Hash: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.MakeResponse(data, true, errors.ERRAuthRegHash.ErrorCode()))
		return
	}

	if !common.IsValidRole(requestBody.Role) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.MakeResponse(data, true, errors.ERRAuthRegInvRole.ErrorCode()))
		return
	}

	newUser := User{
		Email:       requestBody.Email,
		PhoneNumber: requestBody.PhoneNumber,
		Password:    hashedPassword,
		Role:        requestBody.Role,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   time.Time{},
	}

	ok := db.CreateUser(ctx, newUser)
	if ok {
		data["status"] = "User Successfully Created"
	} else {
		data["status"] = "Sorry!!, User Not Created Unwanted Behaviour"
	}
	ctx.JSON(http.StatusOK, helpers.MakeResponse(data, !ok, ""))
}

func login(ctx *gin.Context) {
	data := common.JSON{}
	var requestBody LoginRequestBody
	if err := ctx.BindJSON(&requestBody); err != nil {
		log.Printf("Request Body: %v, Error: %v", requestBody, err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.MakeResponse(data, true, errors.ErrorCodeMessage(errors.ERRReqBody)))
		return
	}
	user, err := db.GetUser(ctx, requestBody.Email, requestBody.PhoneNumber)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.MakeResponse(data, true, errors.ErrorCodeMessage(errors.ERRAuthLogin)))
		return
	}

	if !helpers.CheckPasswordHash(requestBody.Password, user.Password) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.MakeResponse(data, true, errors.ErrorCodeMessage(errors.ERRAuthLogin)))
		return
	}

	userData := common.JSON{}
	userData["_id"] = user.ID
	userData["email"] = user.Email
	userData["phone_number"] = user.PhoneNumber
	userData["role"] = user.Role

	token, err := helpers.GenerateToken(userData)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.MakeResponse(data, true, errors.ERRAuthLoginTokGen.ErrorCode()))
		return
	}
	data["token"] = token

	ctx.JSON(http.StatusOK, helpers.MakeResponse(data, false, ""))
}
