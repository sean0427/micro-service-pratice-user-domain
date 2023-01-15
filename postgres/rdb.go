package rdbreposity

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/sean0427/micro-service-pratice-user-domain/model"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(ctx context.Context, params *model.GetUsersParams) ([]*model.User, error) {
	var users []*model.User
	tx := r.db.WithContext(ctx)

	if params != nil && (*params).Name != nil {
		tx = tx.Where("name = ?", params.Name)
	}

	result := tx.Find(&users)

	return users, result.Error
}

func (r *repository) GetByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	tx := r.db.WithContext(ctx)

	result := tx.Where("id =?", id).Find(&user)
	return &user, result.Error
}

func (r *repository) Create(ctx context.Context, params *model.CreateUserParams) (string, error) {
	var user model.User

	tx := r.db.WithContext(ctx)
	result := tx.Model(&user).Create([]map[string]interface{}{
		{
			"name":     params.Name,
			"email":    params.Email,
			"password": params.Password,
		},
	})
	if result.Error != nil {
		return "", result.Error
	}

	return fmt.Sprintf("%d", user.ID), nil
}

func (r *repository) Update(ctx context.Context, id string, params *model.UpdateUserParams) (*model.User, error) {
	var user model.User
	tx := r.db.WithContext(ctx)

	result := tx.Model(&user).Where("id =?", params.ID).Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}
	return &user, nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	tx := r.db.WithContext(ctx)

	result := tx.Delete(&model.User{}, "id =?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}
