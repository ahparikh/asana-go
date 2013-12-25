package asana

type Task struct {
  Assignee User
  Assignee_status string
  Created_at string
  Completed bool
  Completed_at string
  Due_on string
  Modified_at string
  Name string
  Notes string
  Projects []Project
}

type DataTask struct {
  Data Task
}

func getTaskInfo(ac *asanaclient, task_id string) {}
func getTasksForUser(ac *asanaclient, user_id string) {}
func getTasksForProject(ac *asanaclient, project_id string) {}
