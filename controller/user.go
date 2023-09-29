package controller

import (
	"app/model"
	"math/rand"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
)

type UserController struct {
	Users []model.User
}

func NewUserController(users []model.User) UserController {
	return UserController{
		Users: users,
	}
}

func (u UserController) GetList(req model.GetListUserRequest) model.GetListUserResponse {

	if req.Limit == 0 {
		req.Limit = 10
	}

	var resp = model.GetListUserResponse{}
	if req.Offset+req.Limit > len(u.Users) {
		return model.GetListUserResponse{}
	}

	for i := req.Offset; i < req.Limit+req.Offset; i++ {
		resp.Users = append(resp.Users, u.Users[i])
	}
	resp.Count = len(u.Users)

	return resp
}

func UserGenerateData(limit int) []model.User {
	var users []model.User

	for i := 0; i < limit; i++ {
		users = append(users, model.User{
			Id:        uuid.New().String(),
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			Age:       rand.Intn(100) + 10,
		})
	}

	return users
}
