## SELECT dữ liệu từ bảng và đổ vào struct

1. Setup

Xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/connect-mysql-db

2. Tạo struct hứng dữ liệu đổ về

```go
type Employee struct {
	EmpNo      int
	BeginName  string
	FinalName   string
	Sex     string
}
```

3. SELECT từ bảng và đổ vào struct

```go
var firstEmployee Employee
db.Table("employees").Select("emp_no, first_name AS begin_name, last_name AS final_name, gender AS sex").Where("emp_no = ?", 10001).Scan(&firstEmployee)
```
Để đảm bảo nguyên tắc tên trường trong struct khớp với tên cột trong bảng, sử dụng ALIAS cho tên bảng:
- first_name AS begin_name --> khớp với trường BeginName
- last_name AS final_name --> khớp với trường FinalName
- gender AS sex --> khớp với trường Sex
- Chú ý tên trường cần viết hoa chữ cái đầu