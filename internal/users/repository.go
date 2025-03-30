package users

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) UpdateUser(user *User) error {
	return r.DB.Save(user).Error
}

func (r *UserRepository) DeleteUser(id uuid.UUID) error {
	return r.DB.Delete(&User{}, "id = ?", id).Error
}
