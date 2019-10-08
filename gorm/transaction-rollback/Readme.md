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
Để kiểm tra lỗi của từng CRUD opertion, dùng method GetErrors và check length của slice trả về

```go
var department = Department{
	DeptNo:   "d011",
	DeptName: "Design",
}

tx := db.Begin()

errorList := tx.Exec(`
	UPDATE departments
	SET dept_name = 'Quality Test'
	WHERE dept_no = ?
`, "d006").GetErrors()
if len(errorList) != 0 {
	tx.Rollback()
	return
}

// log.Println("errors", errorList, len(errorList))

errorList = tx.Create(department).GetErrors()
if len(errorList) != 0 {
	tx.Rollback()
	return
}

// log.Println("errors", errorList, len(errorList))

tx.Commit()
```