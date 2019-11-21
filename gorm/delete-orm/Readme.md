## DELETE dữ liệu thông qua ID của struct

1. Setup

Xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/connect-mysql-read-config

2. Tạo struct để lấy thông tin item cần delete
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

3. Xác định ID của post cần xoá

```go
var postInfo Post
	
// WARNING !!!!!!!
// Khi code API thì ID của post sẽ lấy từ đường dẫn hoặc form-data/JSON
// Cần kiểm tra kĩ, không để xảy ra trường hợp postInfo.ID = ""
// Nếu không gorm sẽ xoá sạch dữ liệu trong bảng --> DISASTER !!!!!
postInfo.ID = "123abc"
```

4. Xoá dữ liệu

```go
errDeletePost := db.Delete(&postInfo).Error
// Khi code API thì chỗ này trả về HTTP status code 500
if errDeletePost != nil {
	log.Println(errDeletePost)
	return
}
```

5. Kiểm tra xem đã DELETE thành công chưa

```go
var postInfoAfter Post
// Khi code API thì ID của user sẽ lấy từ đường dẫn (GET request) hoặc form-data/JSON (POST)
errGetPost := db.Where("id = ?", "123abc").Find(&postInfoAfter).Error
// Khi code API thì chỗ này trả về HTTP status code 500
if errGetPost != nil {
	log.Println(errGetPost)
	return
}
log.Println("After update", postInfoAfter)
```