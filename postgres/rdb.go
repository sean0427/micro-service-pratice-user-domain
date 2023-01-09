package rdbreposity

import (
	"context"

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
	// TODO
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
