package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error)
	StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error)
	StoreManyCategory(ctx context.Context, categories []entity.Category) error
	GetCategoryByID(ctx context.Context, id int) (entity.Category, error)
	UpdateCategory(ctx context.Context, category *entity.Category) error
	DeleteCategory(ctx context.Context, id int) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error) {
	var result []entity.Category
	r.db.WithContext(ctx)

	err := r.db.Where("user_id = ?", id).Find(&result).Error

	if err != nil {
		return []entity.Category{}, err
	}

	return result, nil
}

func (r *categoryRepository) StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error) {
	r.db.WithContext(ctx)

	if err := r.db.Create(&category).Error; err != nil {
		return 0, err
	}

	return category.ID, nil
}

func (r *categoryRepository) StoreManyCategory(ctx context.Context, categories []entity.Category) error {
	r.db.WithContext(ctx)

	if err := r.db.Create(&categories).Error; err != nil {
		return err
	}

	return nil
}

func (r *categoryRepository) GetCategoryByID(ctx context.Context, id int) (entity.Category, error) {
	r.db.WithContext(ctx)

	var result entity.Category

	err := r.db.Where("id = ?", id).First(&result).Error

	if err != nil {
		return entity.Category{}, err
	}

	return result, nil
}

func (r *categoryRepository) UpdateCategory(ctx context.Context, category *entity.Category) error {
	r.db.WithContext(ctx)

	if err := r.db.Model(&entity.Category{}).Where("id = ?", category.ID).Updates(&category).Error; err != nil {
		return err
	}

	return nil
}

func (r *categoryRepository) DeleteCategory(ctx context.Context, id int) error {
	r.db.WithContext(ctx)

	if err := r.db.Where("id = ?", id).Delete(&entity.Category{}).Error; err != nil {
		return err
	}

	return nil
}
