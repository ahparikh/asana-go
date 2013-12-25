package main

import (
   "asana"
   "fmt"
)

func main() {
  ac := asana.NewAsanaClient("3xIOLUZ.SGUUZi3bz6FtkuTiKszhDAzC")
  user := asana.GetUserInfo(ac, "me")
  //user := asana.GetAllUsers(ac)
  fmt.Printf("%+v\n", user)
}
