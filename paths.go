package asana

import (
	"fmt"
)

const (
	BASE_URL     string = "https://app.asana.com/api/"
	VERSION      string = "1.0"
	USER_SUBPATH string = "/users"
)

/* User Paths */
func GetUserPath(user string) string {
	base_path := getBasePath()

	if len(user) > 0 {
		return base_path + USER_SUBPATH + "/" + user
	}

	return base_path + USER_SUBPATH
}

/* Workspace Paths */
func GetWorkspacePath(workspace_id int64) string {
	base_path := getBasePath()
	if workspace_id > 0 {
		return fmt.Sprintf("%s/workspaces/%d", base_path, workspace_id)
	}
	return fmt.Sprintf("%s/workspaces/", base_path)
}

func GetUserWorkspacePath(workspace_id int64) string {
	if workspace_id > 0 {
		return fmt.Sprintf("%s/users", GetWorkspacePath(workspace_id))
	}
	return ""
}

func GetProjectsInWorkspacePath(workspace_id int64) string {
	if workspace_id > 0 {
		return fmt.Sprintf("%s/projects", GetWorkspacePath(workspace_id))
	}
	return ""
}

/* Project Paths */
func GetProjectPath(project_id int64) string {
	base_path := getBasePath()
	if project_id > 0 {
		return fmt.Sprintf("%s/projects/%d", base_path, project_id)
	}
	return fmt.Sprintf("%s/projects/", base_path)
}

func GetTasksInProjectPath(project_id int64) string {
	if project_id > 0 {
		return fmt.Sprintf("%s/tasks", GetProjectPath(project_id))
	}
	return ""
}

/* Task Paths */
func GetTaskPath(task_id int64) string {
	base_path := getBasePath()
	if task_id > 0 {
		return fmt.Sprintf("%s/tasks/%d", base_path, task_id)
	}
	return fmt.Sprintf("%s/tasks/", base_path)
}


// Private
func getBasePath() string {
	return BASE_URL + VERSION
}
