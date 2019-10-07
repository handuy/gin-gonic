## SELECT dữ liệu từ bảng và đổ vào slice

1. Setup

Xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/connect-mysql-read-config

2. Tạo struct hứng dữ liệu đổ về

```go
type Employee struct {
	Name string
	Age  int
	Gender string
}
```

3. SELECT từ bảng và đổ vào slice, sử dụng raw query

```go
var employeeList []Employee
db.Raw(`
    SELECT CONCAT(first_name, ' ', last_name) AS name, 
    ROUND( DATEDIFF( CURRENT_DATE(), birth_date)/365, 0 ) AS age,
    gender
    FROM employees
    ORDER BY last_name DESC
    LIMIT 20
    OFFSET 10
`).Scan(&employeeList)

for _, employee := range employeeList {
    log.Println("Employee", employee)
}
```