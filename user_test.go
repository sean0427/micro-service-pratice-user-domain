package user

import (
	"context"
	"strconv"
	"testing"

	"github.com/sean0427/micro-service-pratice-user-domain/api_model"
	"github.com/sean0427/micro-service-pratice-user-domain/model"
)

type mockRepo struct{}

func (r *mockRepo) Get(ctx context.Context, params *api_model.GetUsersParams) ([]*model.User, error) {
	return []*model.User{{ID: 1234}}, nil
}

func (r *mockRepo) GetByID(ctx context.Context, id int64) (*model.User, error) {
	return &model.User{
		ID:   1234,
		Name: strconv.Itoa(int(id))}, nil
}

func (r *mockRepo) Create(ctx context.Context, params *api_model.CreateUserParams) (int64, error) {
	id, err := strconv.Atoi(params.Name)
	return int64(id), err
}

func (r *mockRepo) Update(ctx context.Context, id int64, prarams *api_model.UpdateUserParams) (*model.User, error) {
	return &model.User{
		ID:   prarams.ID,
		Name: prarams.Name,
	}, nil
}

func (r *mockRepo) Delete(ctx context.Context, id int64) error {
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
		const testID = 1
		item, err := testService.GetByID(context.Background(), testID)
		if err != nil {
			t.Error(err)
		}

		if item.Name != strconv.Itoa(testID) {
			t.Errorf("Get user by name is not equal")
		}

		if item.ID == 0 {
			t.Errorf("Returned user id should not be zero")
		}
	})
}

func TestUserService_Create(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		// workaound name as id
		const testId int64 = 123
		user := &api_model.CreateUserParams{
			Name: strconv.Itoa(int(testId)),
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
		const testId = 1234
		testUser := &api_model.UpdateUserParams{
			ID:   1234,
			Name: "1234",
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
		const testId = 1234

		err := testService.Delete(context.Background(), testId)
		if err != nil {
			t.Error(err)
		}
	})

}
