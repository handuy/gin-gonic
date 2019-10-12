## Ghi lỗi ra file log

1. Hàm bắt lỗi và ghi ra file log
```go
func LogError(c *gin.Context, message string) {
	var errorInfo ErrorInfo
	errorInfo.Time = time.Now()
	errorInfo.Path = c.FullPath()
	errorInfo.Method = c.Request.Method

	var log = logrus.New()
	log.Out = os.Stdout
	file, err := os.OpenFile("log/logging.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Println(err)
	}

	log.WithFields(logrus.Fields{
		"Path":    errorInfo.Path,
		"Method": errorInfo.Method,
	}).Error(message)
}
```

2. Trong mỗi route, gọi hàm **LogError** khi kiểm tra lỗi, ví dụ như trong **POST /upload-video**
```go
file, err := c.FormFile("video-file")
if err != nil {
	c.JSON(http.StatusBadRequest, err.Error())
	LogError(c, err.Error())
	return
}

fileType := file.Header["Content-Type"][0]

if GetAllowFormat(fileType, allowedMediaType) == "" {
	c.JSON(http.StatusBadRequest, "Định dạng không hợp lệ. Yêu cầu: mp4")
	LogError(c, "Định dạng không hợp lệ. Yêu cầu: mp4")
	return
}

err = c.SaveUploadedFile(file, path.Join("./upload", file.Filename))
if err != nil {
	c.JSON(http.StatusInternalServerError, err.Error())
	LogError(c, err.Error())
	return
}
```