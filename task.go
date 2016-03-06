package asana

import (
//	"encoding/json"
//	"fmt"
)

type Task struct {
	Id              int64
	Assignee        User
	Assignee_status string
	Created_at      string
	Completed       bool
	Completed_at    string
	Due_on          string
	Modified_at     string
	Name            string
	Notes           string
	Projects        []Project
}

func getTask(ac *asanaclient, task_id int64) Task {
  var task Task
	ac.GetResponse(GetTasksInProjectPath(task_id), &task)
  return task
}

/*
func getTasksForUser(ac *asanaclient, user_id string) []Task {
  var tasks []Task
	ac.GetResponse(GetTasksInProjectPath(user_id), &tasks)
  return tasks
}
*/

func GetTasksForProject(ac *asanaclient, project_id int64) []Task {
  var tasks []Task
	ac.GetResponse(GetTasksInProjectPath(project_id), &tasks)
  return tasks
}
