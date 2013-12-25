package asana

import (
"encoding/json"
"fmt"
)

type User struct {
	Id         int64
	Name       string
	Email      string
	Workspaces []Workspace
}

type DataUser struct {
	Data User
}

type DataUsers struct {
	Data []User
}

func GetUserInfo(ac *asanaclient, user string) (User) {
	body := ac.GetResponse(GetUserPath(user))
	/*
	err := json.Unmarshal(body, &user)

	if err != nil {
		fmt.Println(err)
	}

	return
	*/

	var f DataUser
	err := json.Unmarshal(body, &f)

	if err != nil {
		fmt.Println(err)
	}

	return f.Data
}

func GetAllUsers(ac *asanaclient) ([]User) {
	body := ac.GetResponse(GetUserPath(""))
	var f DataUsers
	err := json.Unmarshal(body, &f)

	if err != nil {
		fmt.Println(err)
	}

	return f.Data
}