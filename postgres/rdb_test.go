package rdbreposity

import (
	"context"
	"os"
	"strconv"
	"testing"

	"github.com/sean0427/micro-service-pratice-user-domain/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	testingDB *gorm.DB
)

func TestMain(m *testing.M) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	testingDB = db
	testingDB.AutoMigrate(&model.User{})

	m.Run()
	os.Exit(0)
}

var testGetUsersCases = []struct {
	name       string
	testCount  int
	testParams model.GetUsersParams
	wantCount  int
	wantError  bool
}{
	{
		name:      "zero path - GetUsers",
		testCount: 0,
		wantCount: 0,
		wantError: false,
	},
	{
		name:      "happy path - GetUsers",
		testCount: 10,
		wantCount: 10,
		wantError: false,
	},
	{
		name:      "error path",
		wantCount: 10,
		wantError: true,
	},
	{
		name: "filter path - fullname contains 1",
		testParams: model.GetUsersParams{
			Name: model.StringToPointer("1"),
		},
		testCount: 20,
		wantCount: 3, // 1, 10, 11
		wantError: false,
	},
	{
		name: "filter path - fullname contains test",
		testParams: model.GetUsersParams{
			Name: model.StringToPointer("test"),
		},
		testCount: 20,
		wantCount: 20,
		wantError: false,
	},
}

func Test_reposity_Get(t *testing.T) {
	for _, c := range testGetUsersCases {

		t.Run(c.name, func(t *testing.T) {
			createRandomUserToDB(c.testCount)
			testParams := c.testParams
			repo := repository{db: testingDB}

			prodct, err := repo.Get(context.Background(), &testParams)

			if err != nil && !c.wantError {
				t.Errorf("got error %v", err)
				return
			}
			if len(prodct) != c.wantCount {
				t.Errorf("Expected %d users, got %d", c.wantCount, len(prodct))
			}
		})
	}
}

func createRandomUserToDB(numbers int) {
	for i := 0; i < numbers; i++ {
		user := &model.User{
			ID:   strconv.Itoa(i),
			Name: "test" + strconv.Itoa(i),
		}
		testingDB.Create(user)
	}
}

var testGetUserIdCases = []struct {
	name      string
	id        string
	testCount int
	want      string
	wantError bool
}{
	{
		name:      "happy - get user id",
		id:        "0",
		testCount: 1,
		wantError: false,
	},
	{
		name:      "happy - get user id 2",
		id:        "1",
		testCount: 2,
		wantError: false,
	},
	{
		name:      "happy - get user id 100",
		id:        "99",
		testCount: 100,
		wantError: false,
	},
	{
		name:      "error - not create",
		testCount: 0,
		id:        "1",
		wantError: true,
	},
	{
		name:      "error - random string",
		testCount: 20,
		id:        "gfeawgeawgew",
		wantError: true,
	},
}

func Test_repository_GetByID(t *testing.T) {
	for _, c := range testGetUserIdCases {
		t.Run(c.name, func(t *testing.T) {
			createRandomUserToDB(c.testCount)
			repo := repository{db: testingDB}

			prodct, err := repo.GetByID(context.Background(), c.id)

			if err != nil && !c.wantError {
				t.Errorf("got error %v", err)
				return
			}

			if prodct.ID == c.id {
				t.Errorf("Expected %s, got %s", c.id, prodct.ID)
			}
		})
	}
}

func Test_repository_Create(t *testing.T) {
	tests := []struct {
		name    string
		params  *model.CreateUserParams
		want    string
		wantErr bool
	}{
		{
			name: "happy",
			params: &model.CreateUserParams{
				Name:     "test",
				Email:    "test@test.com",
				Password: "featea",
			},
			want:    "test",
			wantErr: false,
		},
		{
			name: "error",
			params: &model.CreateUserParams{
				Name: "",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{db: testingDB}
			got, err := r.Create(context.Background(), tt.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("repository.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_Delete(t *testing.T) {
	createRandomUserToDB(10)

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "happy",
			id:      "1",
			wantErr: false,
		},
		{
			name:    "error",
			id:      "11",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{db: testingDB}
			if err := r.Delete(context.Background(), tt.id); (err != nil) != tt.wantErr {
				t.Errorf("repository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
