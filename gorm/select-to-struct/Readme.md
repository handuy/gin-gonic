## SELECT dữ liệu từ bảng và đổ vào struct

1. Setup

Xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/connect-mysql-db

2. Tạo struct hứng dữ liệu đổ về

```go
type Employee struct {
	EmpNo      int
	FirstName  string
	LastName   string
}
```
Tên các trường cần khớp với tên các cột tương ứng:
- Cột tên là emp_no --> Tên trường là EmpNo
- Cột tên là first_name --> Tên trường là FirstName
- Cột tên là last_name --> Tên trường là LastName

3. SELECT từ bảng và đổ vào struct

```go
var firstEmployee Employee
db.Table("employees").Select("emp_no, first_name, last_name").Where("emp_no = ?", 10001).Scan(&firstEmployee)
```