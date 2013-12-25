package asana

const (
    BASE_URL string = "https://app.asana.com/api/"
    VERSION string = "1.0"
    USER_SUBPATH string = "/users"
)

func GetUserPath(user string) string {
	base_path := getBasePath()

	if len(user) > 0 {
		return base_path + USER_SUBPATH + "/" + user	
	}

	return base_path + USER_SUBPATH
}

// Private
func getBasePath() string {
	return BASE_URL + VERSION
}


