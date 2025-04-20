package exercises

import (
	"github.com/gin-gonic/gin"
)

func CreateRouter(router *gin.Engine) {
	exercisesRouter := router.Group("/exercises")
	exercisesRouter.GET("", getExercises)
}
