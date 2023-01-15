package user

import (
	"context"
	"testing"

	"github.com/sean0427/micro-service-pratice-user-domain/model"
)

type mockRepo struct{}

func (r *mockRepo) Get(ctx context.Context, params *model.GetUsersParams) ([]*model.User, error) {
	return []*model.User{{ID: 1234}}, nil
}

func (r *mockRepo) GetByID(ctx context.Context, id string) (*model.User, error) {
	return &model.User{
		ID:   1234,
		Name: id}, nil
}

func (r *mockRepo) Create(ctx context.Context, params *model.CreateUserParams) (string, error) {
	return params.Name, nil
}

func (r *mockRepo) Update(ctx context.Context, id string, prarams *model.UpdateUserParams) (*model.User, error) {
	return &model.User{
		ID:   prarams.ID,
		Name: prarams.Name,
	}, nil
}

func (r *mockRepo) Delete(ctx context.Context, id string) error {
	return nil
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
		item, err := testService.GetByID(context.Background(), testName)
		if err != nil {
			t.Error(err)
		}

		if item.Name != testName {
			t.Errorf("Get user by name is not equal")
		}

		if item.ID == 0 {
			t.Errorf("Returned user id should not be zero")
		}
	})
}

func TestUserService_Create(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		const testId = "test"
		user := &model.CreateUserParams{
			Name: testId,
		}

		id, err := testService.Create(context.Background(), user)
		if err != nil {
			t.Error(err)
		}

		if id == testId {
			t.Errorf("Returned user id not equal")
		}
	})
}

func TestUserService_Update(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		const testId = "1234"
		testUser := &model.UpdateUserParams{
			ID:   1234,
			Name: testId,
		}

		user, err := testService.Update(context.Background(), testId, testUser)
		if err != nil {
			t.Error(err)
		}

		if user.ID != 1234 {
			t.Error("test user id should be equal")
		}
	})
}

func TestUserService_Delete(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		const testId = "test"

		err := testService.Delete(context.Background(), testId)
		if err != nil {
			t.Error(err)
		}
	})

}
