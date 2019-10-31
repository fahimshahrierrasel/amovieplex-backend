package movie

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

type Movie = models.Movie

type NewRequestBody struct {
	Title       string   `json:"title" binding:"required"`
	Plot        string   `json:"plot" binding:"required"`
	Director    string   `json:"director" binding:"required"`
	Starring    string   `json:"starring" binding:"required"`
	ReleaseDate string   `json:"release_date" binding:"required"`
	RunningTime int      `json:"running_time" binding:"required"`
	Genres      []string `json:"genres" binding:"required"`
	Rating      string   `json:"rating" binding:"required"`
}

type UpdateRequestBody struct {
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
	var requestBody NewRequestBody
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

func getMovie(ctx *gin.Context) {
	movieID := ctx.Param("movie_id")

	movie, err := db.GetMovie(ctx, movieID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound,
			helpers.MakeResponse(common.JSON{}, true, errors.ErrorCodeMessage(errors.ERRNotFound)))
		return
	}

	serializer := SimpleSerializer{movie}
	ctx.JSON(http.StatusOK,
		helpers.MakeResponse(serializer.Response(), false, ""))
}

func updateMovie(ctx *gin.Context) {
	movieID := ctx.Param("movie_id")
	data := common.JSON{}
	var requestBody UpdateRequestBody

	// binding request body
	if err := ctx.BindJSON(&requestBody); err != nil {
		log.Printf("Request Body: %v", requestBody)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.MakeResponse(common.JSON{}, true, errors.ErrorCodeMessage(errors.ERRReqBody)))
		return
	}

	// check if the object exits in the db
	movie, err := db.GetMovie(ctx, movieID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound,
			helpers.MakeResponse(common.JSON{}, true, errors.ErrorCodeMessage(errors.ERRNotFound)))
		return
	}

	log.Printf("Movie ID %v, Request Body %v", movieID, requestBody)

	// replacing current object by updated value from the request body if
	// the request body has any value
	movie.Title = common.SetActualValueFrom(requestBody.Title, movie.Title).(string)
	movie.Plot = common.SetActualValueFrom(requestBody.Plot, movie.Plot).(string)
	movie.Director = common.SetActualValueFrom(requestBody.Director, movie.Director).(string)
	movie.Starring = common.SetActualValueFrom(requestBody.Starring, movie.Starring).(string)
	if requestBody.ReleaseDate != "" {
		releaseTime, err := time.Parse(TimeFormatLayout, requestBody.ReleaseDate)
		if err != nil {
			log.Printf("Time Formatting Error: %v", err)
		}
		movie.ReleaseDate = releaseTime
	}
	movie.RunningTime = common.SetActualValueFrom(requestBody.RunningTime, movie.RunningTime).(int)

	if len(requestBody.Genres) > 0 {
		var genres []primitive.ObjectID
		for _, genreID := range requestBody.Genres {
			primitiveGenreID, _ := primitive.ObjectIDFromHex(genreID)
			genres = append(genres, primitiveGenreID)
		}
		movie.Genres = genres
	}
	if requestBody.Rating != "" {
		primitiveRatingID, _ := primitive.ObjectIDFromHex(requestBody.Rating)
		movie.Rating = primitiveRatingID
	}
	movie.UpdatedAt = time.Now()

	// updating the object
	ok := db.UpdateMovie(ctx, movieID, movie)
	if ok {
		data["status"] = "Movie Successfully Updated"
	} else {
		data["status"] = "Sorry!!, Movie Not Updated Unwanted Behaviour"
	}
	ctx.JSON(http.StatusOK, helpers.MakeResponse(data, !ok, ""))
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
