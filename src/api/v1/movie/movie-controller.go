package movie

import (
	"amovieplex-backend/src/api/helpers"
	"amovieplex-backend/src/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

type Movie = models.Movie

type RequestBody struct {
	Title       string   `json:"title"`
	Plot        string   `json:"plot"`
	Director    string   `json:"director"`
	Starring    string   `json:"starring"`
	ReleaseDate string   `json:"release_date"`
	RunningTime int      `json:"running_time"`
	Genres      []string `json:"genres"`
	Rating      string   `json:"rating"`
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
	tempGenres := requestBody.Genres
	newMovie := Movie{Title: requestBody.Title, Plot: requestBody.Plot, Director: requestBody.Director,
		Starring: requestBody.Starring}
}