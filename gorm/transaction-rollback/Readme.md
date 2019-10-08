## Tạo transaction và rollback khi xảy ra lỗi

1. Setup

Xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/connect-mysql-read-config

2. Khởi tạo transaction

```go
tx := db.Begin()
```

3. Sau khi khởi tạo transaction, thay vì dùng db thì ta sẽ dùng tx để CRUD:

- Nếu xảy ra lỗi --> tx.Rollback() để cancel transaction
- Tất cả các CRUD operation thành công --> tx.Commit() để lưu thay đổi vào disk

```go
var department = Department{
	DeptNo:   "d012",
	DeptName: "Technical",
}

tx := db.Begin()

errorInsert := tx.Exec(`
	INSERT INTO departments VALUES (?, ?)`, department.DeptNo, department.DeptName).Error
if errorInsert != nil {
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