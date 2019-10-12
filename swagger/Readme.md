## Tạo API document bằng Swagger

1. Download và cài đặt:
```go
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
go get -u github.com/alecthomas/template
```

2. Copy file binary **swag** từ thư mục **$GOPATH/bin** sang một trong các thư mục ở **$PATH**
```bash
cp $GOPATH/bin/swag /usr/local/bin
```

3. Viết comment cho từng route, ví dụ với route **POST /signup**
```go
// @Description Đăng kí tài khoản mới
// @Accept multipart/form-data
// @Param phone formData string true "Số điện thoại"
// @Param password formData string true "Mật khẩu tối thiểu 4 kí tự"
// @Success 200 {object} controller.Account
// @Failure 400 {object} controller.ErrorMesssage
// @Failure 500 {object} controller.ErrorMesssage
// @Router /signup [post]
func Register(c *gin.Context) {
	// Code logic ...
}
```
Tham khảo cách viết comment tại: https://github.com/swaggo/swag/tree/master/example

4. Tại thư mục chứa file main.go, gõ lệnh:
```bash
swag init
```
swag sẽ tạo một thư mục **docs** chứa thông tin về các API

5. Trong file main.go
```go
import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	
	_ "gin-gonic/swagger/docs" // đường dẫn đến thư mục docs vừa được tạo ở bước 4
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*.html")

	url := ginSwagger.URL("http://localhost:8085/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.POST("/signup", controller.Register)
	router.POST("/login", controller.Login)

	router.Run(":8085")
}
```

6. Truy cập: http://localhost:8085/swagger/index.html