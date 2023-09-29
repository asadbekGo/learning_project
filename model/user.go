package model

type UserPrimaryKey struct {
	Id string
}

type CreateUser struct {
	FirstName string
	LastName  string
	Age       int
}

type User struct {
	Id        string
	FirstName string
	LastName  string
	Age       int
	CreatedAt string
}

type UpdateUser struct {
	Id        string
	FirstName string
	LastName  string
	Age       int
}

type GetListUserRequest struct {
	Offset int
	Limit  int
}

type GetListUserResponse struct {
	Count int
	Users []User
}
