## Tạo transaction

1. Setup

Xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/connect-mysql-read-config

2. Khởi tạo transaction

```go
tx := db.Begin()
```

3. Sau khi khởi tạo transaction, thay vì dùng db thì ta sẽ dùng tx để CRUD. Gọi hàm tx.Commit() để lưu 
thay đổi vào disk

```go
var department = Department{
	DeptNo:   "d010",
	DeptName: "Accounting",
}

tx := db.Begin()

tx.Create(department)

tx.Exec(`
	UPDATE departments
	SET dept_name = 'Quality Control'
	WHERE dept_no = ?
`, "d006")

tx.Commit()
```