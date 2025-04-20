package exercises

import (
	"net/http"
	"trainingsplan/api/lib"
	"trainingsplan/api/models"

	"github.com/gin-gonic/gin"
)

// @Summary Get all global exercises
// @Success 200 {array} models.ExerciseResponse
// @Router /exercises [get]
func getExercises(context *gin.Context) {
	var exercises []models.Exercise
	result := lib.Database.Find(&exercises)

	if result.Error != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "An error occured while fetching exercises"})
		return
	}

	context.JSON(http.StatusOK, exercises)
}
