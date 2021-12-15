package database

import "projectNIX/repositories/models"

type UserDBRepository struct{}

func (udbr UserDBRepository) GetByEmail(email string) models.User {
	//TODO SELECT email, password_hash, created_at FROM users WHERE email = email
	panic("implement me")
}
