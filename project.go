package asana

import (
//	"fmt"
)

type Project struct {
	Id          int64
	Archived    bool
	Created_at  string
	Modified_at string
	Name        string
	Notes       string
}

func GetProject(ac *asanaclient, project_id int64) Project {
  var project Project
  ac.GetResponse(GetProjectPath(project_id), &project)
  return project
}

func GetProjectsInWorkspace(ac *asanaclient, workspace_id int64) []Project {
	var projects []Project
	ac.GetResponse(GetProjectsInWorkspacePath(workspace_id), &projects)
	return projects
}
