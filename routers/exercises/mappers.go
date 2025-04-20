package exercises

import "trainingsplan/api/models"

func mapEntityToResponse(entity *models.Exercise) models.ExerciseResponse {
	return models.ExerciseResponse{
		Name:        entity.Name,
		Description: entity.Description,
	}
}
