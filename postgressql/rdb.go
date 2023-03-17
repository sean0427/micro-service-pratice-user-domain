package postgressql

import (
	"context"
	"errors"

	"gorm.io/gorm"

	tool "github.com/sean0427/tool-distributed-system-p/outbox-transaction"

	"github.com/sean0427/micro-service-pratice-user-domain/api_model"
	"github.com/sean0427/micro-service-pratice-user-domain/model"
)

const topicName = "user"

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

	err := tool.TransactionWithOutboxMsg(ctx, r.db, &user, topicName, func(tx *gorm.DB) (int64, error) {
		result := tx.Model(&user).Create(&user)

		// FIXME: workaround
		user.Password = nil
		return user.ID, result.Error
	})

	if err != nil {
		return 0, err
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

	err := tool.TransactionWithOutboxMsg(ctx, r.db, &user, topicName, func(tx *gorm.DB) (int64, error) {
		result := tx.Model(&user).Where("id = ?", params.ID).Save(&user)
		if result.Error != nil {
			return 0, result.Error
		}
		if result.RowsAffected == 0 {
			return 0, errors.New("not found")
		}

		// FIXME: workaround
		user.Password = nil
		return user.ID, nil
	})

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) Delete(ctx context.Context, id int64) error {
	tx := r.db.WithContext(ctx)

	err := tool.TransactionDeleteWithOutboxMsg(ctx, tx, topicName, id, func(tx *gorm.DB) error {
		result := tx.Delete(&model.User{}, "id = ?", id)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("not found")
		}

		return nil
	})

	return err
}

func (r *repository) ExamUserPassword(ctx context.Context, name, password string) (bool, error) {
	tx := r.db.WithContext(ctx)

	var count int64
	result := tx.Model(&model.User{}).
		Where("name =? and password=?", name, password).
		Count(&count)

	if count == 1 {
		return true, nil
	}

	return false, result.Error
}
