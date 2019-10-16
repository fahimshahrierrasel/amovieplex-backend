package movie

import (
	"amovieplex-backend/src/api/helpers"
	"amovieplex-backend/src/data/db"
	"amovieplex-backend/src/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"
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

var (
	TimeFormatLayout = "2006-01-02"
)

func create(ctx *gin.Context) {
	data := map[string]interface{}{}
	var requestBody RequestBody
	if err := ctx.BindJSON(&requestBody); err != nil {
		log.Printf("Request Body: %v", requestBody)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.MakeResponse(data, true, "request body is not correct"))
		return
	}
	var genres []primitive.ObjectID
	for _, genreID := range requestBody.Genres {
		primitiveGenreID, _ := primitive.ObjectIDFromHex(genreID)
		genres = append(genres, primitiveGenreID)
	}

	primitiveRatingID, _ := primitive.ObjectIDFromHex(requestBody.Rating)

	releaseTime, err := time.Parse(TimeFormatLayout, requestBody.ReleaseDate)
	if err != nil {
		log.Printf("Time Formatting Error: %v", err)
	}

	newMovie := Movie{Title: requestBody.Title, Plot: requestBody.Plot, Director: requestBody.Director,
		Starring: requestBody.Starring, ReleaseDate: releaseTime, RunningTime: requestBody.RunningTime,
		Genres: genres, Rating: primitiveRatingID, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}}

	ok := db.CreateMovie(ctx, newMovie)
	if ok {
		data["status"] = "Movie Successfully Created"
	} else {
		data["status"] = "Sorry!!, Movie Not Created Unwanted Behaviour"
	}
	ctx.JSON(http.StatusOK, helpers.MakeResponse(data, !ok, ""))
}

func getAll(ctx *gin.Context) {
	result := db.GetAllMovie(ctx)
	serializer := Serializer{result}
	ctx.JSON(http.StatusOK, helpers.MakeResponse(serializer.Response(), false, ""))
}

func softDelete(ctx *gin.Context) {
	movieID := ctx.Param("movie_id")
	data := map[string]interface{}{}
	ok := db.SoftDeleteMovie(ctx, movieID)
	if ok {
		data["status"] = "Movie Successfully Soft Deleted"
	} else {
		data["status"] = "Sorry!!, Movie Not Soft Deleted Unwanted Behaviour"
	}
	ctx.JSON(http.StatusOK, helpers.MakeResponse(data, !ok, ""))
}

func permanentDelete(ctx *gin.Context) {
	movieID := ctx.Param("movie_id")
	data := map[string]interface{}{}
	ok := db.DeleteMovie(ctx, movieID)
	if ok {
		data["status"] = "Movie Successfully Deleted"
	} else {
		data["status"] = "Sorry!!, Movie Not Deleted Unwanted Behaviour"
	}
	ctx.JSON(http.StatusOK, helpers.MakeResponse(data, !ok, ""))
}
