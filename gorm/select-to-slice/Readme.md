## SELECT dữ liệu từ bảng và đổ vào slice

1. Setup

Xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/connect-mysql-db

2. Tạo struct hứng dữ liệu đổ về

```go
type Employee struct {
	EmpNo     int
	BirthDate time.Time
	FirstName string
	LastName  string
}
```

3. SELECT từ bảng và đổ vào slice

```go
var employeeList []Employee
	db.Table("employees").Select("emp_no, birth_date, first_name, last_name").Where("first_name LIKE ?", "Cristinel%").Scan(&employeeList)

	for _, employee := range employeeList {
        log.Println("Employee", employee)
    }
```