package api_model

type GetUsersParams struct {
	Name *string
}

type CreateUserParams struct {
	Name     string
	Email    string
	Password string
}

type UpdateUserParams struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

func StringToPointer(s string) *string {
	return &s
}
