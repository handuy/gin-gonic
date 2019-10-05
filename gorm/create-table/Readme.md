## Tạo bảng mới

1. Setup

Xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/connect-mysql-db

2. Tạo struct định nghĩa bảng

```go
type Company struct {
	Id        int `gorm:"primary_key"`
	Name      string `gorm:"type:varchar(100)"`
	Address   string `gorm:"type:varchar(100)"`
	IsGlobal  bool
	CreatedAt time.Time
	OtherInfo model.ManagerInfo   `sql:"TYPE:json"`
}
```
trong đó model.ManagerInfo là một struct khác được tạo ở package model
```go
type ManagerInfo struct {
	Name     string
	Age      int
	HireDate time.Time
}
```
struct này sẽ implement 2 method là Value() và Scan()

3. Tạo bảng

```go
var company Company
db.CreateTable(company)
```
Bảng company sẽ có trường other_info có kiểu dữ liệu JSON
