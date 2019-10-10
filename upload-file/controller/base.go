package controller

var allowedMediaType = []string{"video/mp4"}

func GetAllowFormat(format string, allowFormat []string) string {
	var allow string
	for _, item := range allowFormat {
		if format == item {
			allow = item
		}
	}
	return allow
}