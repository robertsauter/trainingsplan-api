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
	var exerciseEntities []models.Exercise
	result := lib.Database.Find(&exerciseEntities)

	if result.Error != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "An error occured while fetching exercises"})
		return
	}

	exerciseResponses := make([]models.ExerciseResponse, 0)

	for _, exerciseEntity := range exerciseEntities {
		exerciseResponses = append(exerciseResponses, mapEntityToResponse(&exerciseEntity))
	}

	context.JSON(http.StatusOK, exerciseResponses)
}
