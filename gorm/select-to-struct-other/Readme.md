## SELECT dữ liệu từ bảng và đổ vào struct

1. Setup

Xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/connect-mysql-read-config

2. Tạo struct hứng dữ liệu đổ về

```go
type UserInfo struct {
	UserId      int
	FullName   string
	Email    string
	ActiveOrNot   int
}
```
Tên các trường cần viết theo kiểu PascalCase. Trường hợp struct hứng dữ liệu có tên các trường 
không khớp với tên cột thì khi query cần dùng AS để đổi tên cột:

```go
var userInfo UserInfo
// Sử dụng column alias để cho tên bảng khớp với tên trường của struct hứng dữ liệu
errGetUser := db.Table("users").
			Select("id AS user_id, name AS full_name, email, is_active AS active_or_not").
			Where("id = ?", 2).Scan(&userInfo).Error
// Khi code API thì chỗ này trả về HTTP status code 500
if errGetUser != nil {
	log.Println(errGetUser)
	return
}
```