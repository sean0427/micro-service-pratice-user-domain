package model

type GetUsersParams struct {
	Name *string
}

func StringToPointer(s string) *string {
	return &s
}
