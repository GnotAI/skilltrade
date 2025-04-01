package users

import (
  "errors"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetUserByEmail(email string) (*User, error) {
	var user User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByID(userID uuid.UUID) (*User, error) {
    var user User
    result := r.DB.First(&user, "id = ?", userID)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
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
