## Ghi lỗi ra database và file

1. Kết nối database từ file **config.local.json**. Tham khảo thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/connect-mysql-read-config

2. Tạo bảng error_info lưu thông tin lỗi
```go
type ErrorInfo struct {
	Time  time.Time     // thời gian xảy ra lỗi
	Path  string        // lỗi ở đường dẫn nào
	Method string       // tên method
	Message string      // nội dung lỗi
}

var logInfo controller.ErrorInfo
db.AutoMigrate(logInfo)
```

3. Tạo struct Controller lưu thông tin kết nối database để tái sử dụng cho các route handler
```go
type Controller struct {
	DB *gorm.DB
}

control := &controller.Controller{DB: db}

router.GET("/", control.HomePage)
router.POST("/upload-video", control.UploadVideo)
```

4. Trong từng handler, gọi hàm LogError trong logic check lỗi
```go
func(control *Controller) UploadVideo(c *gin.Context) {
	file, err := c.FormFile("video-file")
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		control.LogError(c, err.Error())
		return
	}

	fileType := file.Header["Content-Type"][0]

	if GetAllowFormat(fileType, allowedMediaType) == "" {
		c.JSON(http.StatusBadRequest, "Định dạng không hợp lệ. Yêu cầu: mp4")
		control.LogError(c, "Định dạng không hợp lệ. Yêu cầu: mp4")
		return
	}

	err = c.SaveUploadedFile(file, path.Join("./upload", file.Filename))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		control.LogError(c, err.Error())
		return
	}
}
```

5. Hàm LogError ghi lỗi vào database, nếu không ghi được vào db thì sẽ ghi ra file
```go
func(control *Controller) LogError(c *gin.Context, message string) {
	var errorInfo ErrorInfo
	errorInfo.Time = time.Now()
	errorInfo.Path = c.FullPath()
	errorInfo.Method = c.Request.Method
	errorInfo.Message = message

	err := control.DB.Create(errorInfo).Error
	if err != nil {
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
}
```