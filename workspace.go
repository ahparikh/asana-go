package asana

import (
//	"encoding/json"
//	"fmt"
)

type Workspace struct {
	Id   int64
	Name string
}

func GetWorkspace(ac *asanaclient, workspace_id int64) Workspace {
  var workspace Workspace
  ac.GetResponse(GetWorkspacePath(workspace_id), &workspace)
  return workspace
}

func GetAllWorkspaces(ac *asanaclient) []Workspace {
  var workspaces []Workspace
	ac.GetResponse(GetWorkspacePath(0), &workspaces)
  return workspaces
}
