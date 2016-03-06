package asana

import (
//	"fmt"
)

//const USER_SUBPATH string = "/users"

type User struct {
	Id         int64
	Name       string
	Email      string
	Workspaces []Workspace
}

func GetUser(ac *asanaclient, user_str string) User {
	var user User
	ac.GetResponse(GetUserPath(user_str), &user)
	return user
}

func GetAllUsers(ac *asanaclient) []User {
	var users []User
	ac.GetResponse(GetUserPath(""), &users)
	return users
}

func GetUsersInWorkspace(ac *asanaclient, workspace_id int64) []User {
	var users []User
	ac.GetResponse(GetUserWorkspacePath(workspace_id), &users)
	return users
}
