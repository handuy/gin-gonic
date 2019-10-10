## Upload file binary

1. Form upload có dạng
```html
<form action="/upload-video" method="POST" enctype="multipart/form-data">
    <input type="file" name="video-file" accept="video/*">
    <input type="submit">
</form>
```
Chú ý giá trị của thuộc tính **name** của thẻ input là **video-file**. Chúng ta sẽ dùng 
giá trị này để lấy dữ liệu file upload

2. Hàm xử lý POST request gửi đến **/upload-video**
```go
func UploadVideo(c *gin.Context) {
    // Giá trị thuộc tính name của thẻ input là video-file
	file, err := c.FormFile("video-file")
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	fileType := file.Header["Content-Type"][0]

	if GetAllowFormat(fileType, allowedMediaType) == "" {
		c.JSON(http.StatusBadRequest, "Sai định dạng")
		return
    }
    
    // Lưu file vào thư mục upload với tên giống như tên file trong form upload
	err = c.SaveUploadedFile(file, path.Join("./upload", file.Filename))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
}
```