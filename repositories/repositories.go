package repositories

import "projectNIX/repositories/models"

type UserRepositoryInterface interface {
	GetByEmail(email string) models.User
}
