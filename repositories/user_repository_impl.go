package repository

import (
	"creative-write/model/domain"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(DB *gorm.DB) UserRepositoryInterfaces {
	return &UserRepositoryImpl{
		DB: DB,
	}
}

func (repo *UserRepositoryImpl) Find(email string) (domain.User, error) {
	user := domain.User{}
	result := repo.DB.Find(&user, "email = ?", strings.ToLower(email))

	fmt.Println("testtt")
	if result.Error != nil {
		return user, fmt.Errorf("user not found")
	}

	return user, nil
}

func (repo *UserRepositoryImpl) Save(user *domain.User) error {
	err := repo.DB.Create(&user).Error

	if err != nil {
		return err
	}

	return nil
}
