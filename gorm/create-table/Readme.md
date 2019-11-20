## Tạo bảng mới

1. Setup

Xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/connect-mysql-db

2. Tạo struct User định nghĩa cấu trúc bảng **users**

```go
type User struct {
	ID        int       `gorm:"primary_key"`
	Name      string    
	Email     string    
	Age       int       
	IsActive  bool      
	Average   float32   
	CreatedAt time.Time
}
```

3. Tạo struct Post định nghĩa cấu trúc bảng **posts**

```go
type Post struct {
	ID        int       `gorm:"primary_key"`
	Name      string    `gorm:"type:varchar(50)"`
	Email     string    `gorm:"type:varchar(100)"`
	Age       int       `gorm:"type:BIGINT"`
	IsActive  bool      
	Average   float32   `gorm:"type:DECIMAL(6,2)"`
	CreatedAt time.Time
}
```
Struct Post có thêm các tag để chỉ định rõ kiểu dữ liệu của cột, ví dụ 2 trường Name và Email 
đều là string nhưng khi tạo bảng **posts** thì cột name sẽ có kiểu varchar(50), còn email là 
varchar(100)

4. Tạo bảng

```go
var user User
var post Post

db.CreateTable(user)
db.CreateTable(post)
```
