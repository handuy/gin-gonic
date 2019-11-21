## Tạo transaction và rollback khi xảy ra lỗi

1. Setup

Xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/connect-mysql-read-config

2. Khởi tạo transaction

```go
tx := db.Begin()
```

3. Sau khi khởi tạo transaction, thay vì dùng db thì ta sẽ dùng tx để CRUD:

- Nếu xảy ra lỗi --> tx.Rollback() để cancel toàn bộ transaction
- Tất cả các CRUD operation thành công --> tx.Commit() để lưu thay đổi vào disk

```go
tx := db.Begin()
if err := tx.Error; err != nil {
	return
}

// Khi code API thì thông tin về user mới sẽ lấy từ file JSON hoặc form data của client gửi lên
var newUser = User{
	Name:      "Java",
	Email:     "java@oracle.com",
	Age:       20,
	IsActive:  true,
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

errorUpdate := tx.Exec(`
	UPDATEEEEE departments
	SET dept_name = 'Quality Test'
	WHERE dept_no = ?
`, "d006").Error
if errorUpdate != nil {
	tx.Rollback()
	return
}

tx.Commit()
```