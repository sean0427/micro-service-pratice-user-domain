package postgressql

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/sean0427/micro-service-pratice-user-domain/api_model"
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

func (r *repository) Get(ctx context.Context, params *api_model.GetUsersParams) ([]*model.User, error) {
	var users []*model.User
	tx := r.db.WithContext(ctx)

	if params != nil && (*params).Name != nil {
		tx = tx.Where("name = ?", params.Name)
	}

	result := tx.Find(&users)

	return users, result.Error
}

func (r *repository) GetByID(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	tx := r.db.WithContext(ctx)

	result := tx.Where("id = ?", id).Find(&user)
	return &user, result.Error
}

func (r *repository) Create(ctx context.Context, params *api_model.CreateUserParams) (int64, error) {
	user := model.User{
		Name:     params.Name,
		Email:    params.Email,
		Password: params.Password,
	}

	tx := r.db.WithContext(ctx)
	result := tx.Model(&user).Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}

func (r *repository) Update(ctx context.Context, id int64, params *api_model.UpdateUserParams) (*model.User, error) {
	user := model.User{
		ID:       id,
		Name:     params.Name,
		Email:    params.Email,
		Password: params.Password,
	}
	tx := r.db.WithContext(ctx)

	result := tx.Model(&user).Where("id = ?", params.ID).Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}

	return &user, nil
}

func (r *repository) Delete(ctx context.Context, id int64) error {
	tx := r.db.WithContext(ctx)

	result := tx.Delete(&model.User{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("not found")
	}

	return nil
}

func (r *repository) ExamUserPassword(ctx context.Context, name, password string) (bool, error) {
	tx := r.db.WithContext(ctx)

	// TODO: Check
	var examCount int64 = 0
	tx.Find("name = ?", name).
		Where("password =?", password).
		Count(&examCount)

	if examCount == 1 {
		return true, nil
	}
	return false, nil
}
