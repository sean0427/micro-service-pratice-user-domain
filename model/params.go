package model

type GetUsersParams struct {
	Name *string
}

type CreateUserParams struct {
	ID       string
	Name     string
	Email    string
	Password string
}

type UpdateUserParams struct {
	ID       string
	Name     string
	Email    string
	Password string
}

func StringToPointer(s string) *string {
	return &s
}
