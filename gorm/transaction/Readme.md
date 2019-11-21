## Tạo transaction

1. Setup

Xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/connect-mysql-read-config

2. Khởi tạo transaction

```go
tx := db.Begin()
```

3. Sau khi khởi tạo transaction, thay vì dùng db thì ta sẽ dùng tx để CRUD

- Nếu xảy ra lỗi --> tx.Rollback() để cancel toàn bộ transaction
- Tất cả các CRUD operation thành công --> tx.Commit() để lưu thay đổi vào disk

```go
// Khi code API thì thông tin về user mới sẽ lấy từ file JSON hoặc form data của client gửi lên
var newUser = User{
	Name:      "PHP",
	Email:     "php@elephant.com",
	Age:       20,
	IsActive:  false,
	Average:   8.64,
	CreatedAt: time.Now(),
}
errCreateUser := tx.Create(&newUser).Error
// Khi code API thì chỗ này trả về status 500 InternalServerError
if errCreateUser != nil {
	log.Println(errCreateUser)
	tx.Rollback()
	return
}

// Trước khi update
var userInfo User
// Khi code API thì thông tin về user ID 
// và các trường cần update cùng với giá trị tương ứng
// sẽ lấy từ form-data/JSON (POST/PUT)
// trong ví dụ này tạm thời fix cứng
errUpdateUser := tx.Model(&userInfo).Where("id = ?", 3).
	Updates(map[string]interface{}{"name":"AmazonWS", "email":"cloud@aws.com"}).
	Error
// Khi code API thì chỗ này trả về HTTP status code 500
if errUpdateUser != nil {
	log.Println(errUpdateUser)
	tx.Rollback()
	return
}

tx.Commit()
```