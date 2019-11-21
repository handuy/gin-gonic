## UPDATE dữ liệu từ struct vào bảng

1. Setup

Xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/connect-mysql-read-config

2. Tạo struct để lấy thông tin item cần update
Trong trường hợp này, ta lấy luôn struct Post đã dùng để tạo bảng

```go
type Post struct {
	ID        string       `gorm:"primary_key"`
	Name      string       `gorm:"type:varchar(50)"`
	Email     string       `gorm:"type:varchar(100)"`
	Age       int          `gorm:"type:BIGINT"`
	IsActive  bool      
	Average   float32      `gorm:"type:DECIMAL(6,2)"`
	CreatedAt time.Time
}
```

3. SELECT từ bảng và đổ vào struct

```go
// Trước khi update
var postInfo Post
// Khi code API thì ID của user sẽ lấy từ đường dẫn (GET request) hoặc form-data/JSON (POST)
errGetPost := db.Where("id = ?", "123abc").Find(&postInfo).Error
// Khi code API thì chỗ này trả về HTTP status code 500
if errGetPost != nil {
	log.Println(errGetPost)
	return
}
log.Println("Before update", postInfo)
```

4. Cập nhật giá trị của 2 trường Name và Email của biến userInfo vừa lấy dữ liệu về

```go
postInfo.Name = "Kubernetes"
postInfo.Email = "kuber@open.com"
```

5. "Đẩy" biến postInfo vừa được update dữ liệu vào bảng posts

```go
errUpdatePost := db.Save(&postInfo).Error
// Khi code API thì chỗ này trả về HTTP status code 500
if errUpdatePost != nil {
	log.Println(errUpdatePost)
	return
}
```

6. Kiểm tra xem đã UPDATE thành công chưa

```go
var postInfoAfter Post
// Khi code API thì ID của user sẽ lấy từ đường dẫn (GET request) hoặc form-data/JSON (POST)
errGetPost = db.Where("id = ?", "123abc").Find(&postInfoAfter).Error
// Khi code API thì chỗ này trả về HTTP status code 500
if errGetPost != nil {
	log.Println(errGetPost)
	return
}
log.Println("After update", postInfoAfter)
```