package user

import (
	"context"
	"testing"

	"github.com/sean0427/micro-service-pratice-user-domain/model"
)

type mockRepo struct{}

func (r *mockRepo) Get(ctx context.Context, params *model.GetUsersParams) ([]*model.User, error) {
	return []*model.User{{ID: "test"}}, nil
}

func (r *mockRepo) GetByID(ctx context.Context, id string) (*model.User, error) {
	return &model.User{
		ID:   "testfjeia",
		Name: id}, nil
}

func createMockRepo() *mockRepo {
	// TODO
	return &mockRepo{}
}

var testService *UserService

func TestMain(m *testing.M) {
	testService = New(createMockRepo())
}

func TestUserService_Get(t *testing.T) {
	t.Run("Should success get user", func(t *testing.T) {
		list, err := testService.Get(context.TODO(), nil)

		if len(list) == 0 {
			t.Errorf("Get user list is empty")
		}

		if err != nil {
			t.Error(err)
		}
	})
}

func TestUserService_GetByID(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		const testName = "test"
		item, err := testService.GetByID(context.TODO(), testName)
		if err != nil {
			t.Error(err)
		}

		if item.Name != testName {
			t.Errorf("Get user by name is not equal")
		}

		if item.ID == "" {
			t.Errorf("Returned user id should not be empty")
		}
	})
}
