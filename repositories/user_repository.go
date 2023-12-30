package repository

import "creative-write/model/domain"

type UserRepositoryInterfaces interface {
	Find(email string) (domain.User, error)
	Save(*domain.User) error
}
