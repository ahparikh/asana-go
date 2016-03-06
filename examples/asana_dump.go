package main

import (
   "asana"
   "fmt"
)

func main() {
//  ac := asana.NewAsanaClient("3xIOLUZ.SGUUZi3bz6FtkuTiKszhDAzC")
  ac := asana.NewAsanaClient("lHpqx3Dd.CChIDcrYewiqWOJqWTJ2S34")


/*
  ub := []byte(`{"id":123456789,"name":"Akash Parikh","email":"ahparikh@gmail.com"}`)
  ubs := []byte(`[{"id":123456789,"name":"Akash Parikh","email":"ahparikh@gmail.com"}],[{"id":987654321,"name":"John","email":"joe@gmail.com"}]`)
  
  var user asana.User
  err := json.Unmarshal(ub, &user)
  if err != nil {
  	fmt.Println(err)
  }

  var users []asana.User
  err = json.Unmarshal(ubs, &users)
  if err != nil {
  	fmt.Println(err)
  }
 */

  user := asana.GetUser(ac, "me")
  _ = user
  fmt.Printf("USER: %+v\n\n", user)
  users := asana.GetAllUsers(ac)
  _ = users
  fmt.Printf("USERS: %+v\n\n", users)
  workspaces := asana.GetAllWorkspaces(ac)
  _ = workspaces
  fmt.Printf("WORKSPACES: %+v\n\n", workspaces)
  workspace_users := asana.GetUsersInWorkspace(ac, 8143298540087)
  _ = workspace_users
  fmt.Printf("USERS_IN_WORKSPACE: %+v\n\n", workspace_users)
  projects := asana.GetProjectsInWorkspace(ac, 8143298540087)
  _ = projects
  fmt.Printf("PROJECTS: %+v\n\n", projects)
  tasks := asana.GetTasksForProject(ac, 8143298540092)
  _ = tasks
  fmt.Printf("tasks: %+v\n\n", tasks)


  /*
  fmt.Printf("%+v\n\n", workspaces)
  fmt.Printf("%+v\n\n", workspace_users)
  fmt.Printf("%+v\n\n", tasks)
  */
}

