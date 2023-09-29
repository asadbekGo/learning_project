package main

import (
	"app/controller"
	"app/model"
	"fmt"
)

func main() {

	users := controller.UserGenerateData(1000)
	userCon := controller.NewUserController(users)

	// userCon.Create(model.CreateUser{

	// })

	resp := userCon.GetList(model.GetListUserRequest{
		Offset: 1000,
		Limit:  10,
	})

	fmt.Println("Count:", resp.Count)
	for in, user := range resp.Users {
		fmt.Println(in+1, user)
	}
}

// CRUD
// - Create
// - Read
// - Update
// - Delete

// Create(req CreateUser) User
// GetById(req UserPrimaryKey) User
// Update(req UpdateUser) User
// Delete(req UserPrimaryKey)
