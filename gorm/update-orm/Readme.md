## UPDATE dữ liệu từ struct vào bảng

1. Setup

Xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/connect-mysql-read-config

2. Tạo struct để lấy thông tin item cần update
Trong trường hợp này, ta lấy luôn struct User đã dùng để tạo bảng

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

3. SELECT từ bảng và đổ vào struct

```go
// Trước khi update
var userInfo User
// Khi code API thì ID của user sẽ lấy từ đường dẫn (GET request) hoặc form-data/JSON (POST)
errGetUser := db.Where("id = ?", 3).Find(&userInfo).Error
// Khi code API thì chỗ này trả về HTTP status code 500
if errGetUser != nil {
	log.Println(errGetUser)
	return
}
log.Println("Before update", userInfo)
```

4. Cập nhật giá trị của 2 trường Name và Email của biến userInfo vừa lấy dữ liệu về

```go
userInfo.Name = "Kubernetes"
userInfo.Email = "kuber@open.com"
```

5. "Đẩy" biến userInfo vừa được update dữ liệu vào bảng users

```go
errUpdateUser := db.Save(&userInfo).Error
// Khi code API thì chỗ này trả về HTTP status code 500
if errUpdateUser != nil {
	log.Println(errUpdateUser)
	return
}
```

6. Kiểm tra xem đã UPDATE thành công chưa

```go
var userInfoAfter User
// Khi code API thì ID của user sẽ lấy từ đường dẫn (GET request) hoặc form-data/JSON (POST)
errGetUser = db.Where("id = ?", 3).Find(&userInfoAfter).Error
// Khi code API thì chỗ này trả về HTTP status code 500
if errGetUser != nil {
	log.Println(errGetUser)
	return
}
log.Println("After update", userInfoAfter)
```