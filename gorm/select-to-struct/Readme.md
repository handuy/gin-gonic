## SELECT dữ liệu từ bảng và đổ vào struct

1. Setup

Xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/connect-mysql-read-config

2. Tạo struct hứng dữ liệu đổ về

```go
type UserInfo struct {
	Id      int
	Name   string
	Email    string
	IsActive   int
}
```
Tên các trường cần viết theo kiểu PascalCase, khớp với tên các cột tương ứng:
- Cột tên là id --> Tên trường là ID
- Cột tên là name --> Tên trường là Name
- Cột tên là is_active --> Tên trường là IsActive

3. SELECT từ bảng và đổ vào struct

```go
var userInfo UserInfo
db.Table("users").Select("id, name, email, is_active").Where("id = ?", 2).Scan(&userInfo)
```