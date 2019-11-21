## SELECT dữ liệu từ bảng và đổ vào slice

1. Setup

Xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/connect-mysql-read-config

2. Tạo struct hứng dữ liệu đổ về
Nếu cần lấy thông tin tất cả các cột thì dùng luôn struct đã dùng để tạo bảng

```go
type User struct {
	ID        int `gorm:"primary_key"`
	Name      string
	Email     string
	Age       int
	IsActive  bool
	Average   float32
	CreatedAt time.Time
}
```

3. SELECT từ bảng và đổ vào slice

```go
var userInfo []User
errGetUser := db.Find(&userInfo).Error
// Khi code API thì chỗ này trả về HTTP status code 500
if errGetUser != nil {
	log.Println(errGetUser)
	return
}

for _,v := range userInfo {
	log.Println("Employee", v)
}
```