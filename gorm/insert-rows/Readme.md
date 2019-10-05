## INSERT dữ liệu vào bảng

1. Setup

Xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/connect-mysql-db

2. Tạo struct INSERT dữ liệu

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

3. INSERT bản ghi

```go
var managerInfo = model.ManagerInfo{
	Name: "Max Allegri",
	Age: 42,
	HireDate: time.Now(),
}

var company = Company{
	Id: 4,
	Name: "Juventus",
	Address: "Turin",
	IsGlobal: true,
	CreatedAt: time.Now(),
	OtherInfo: managerInfo,
}

db.Create(company)
```