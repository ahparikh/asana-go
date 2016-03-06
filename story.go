package asana

type Story struct {
	Created_at string
	Created_by User
	Text       string
	Target     Task
	Source     string
	Type       string
}

func getCommentInfo(ac *asanaclient, storyId int64)    {}
func getCommentsForTask(ac *asanaclient, taskId int64) {}
