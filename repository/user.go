package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	var result entity.User
	r.db.WithContext(ctx)

	err := r.db.Where("id = ?", id).First(&result).Error

	if err == gorm.ErrRecordNotFound {
		return entity.User{}, nil
	} else if err != nil {
		return entity.User{}, err
	}

	return result, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var result entity.User
	r.db.WithContext(ctx)

	err := r.db.Where("email = ?", email).First(&result).Error

	if err == gorm.ErrRecordNotFound {
		return entity.User{}, nil
	} else if err != nil {
		return entity.User{}, err
	}

	return result, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	r.db.WithContext(ctx)

	if err := r.db.Create(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	r.db.WithContext(ctx)

	if err := r.db.Model(&entity.User{}).Where("id = ?", user.ID).Updates(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	r.db.WithContext(ctx)

	if err := r.db.Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
		return err
	}

	return nil
}
