## INSERT dữ liệu vào bảng

1. Setup

Xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/connect-mysql-db

2. Struct dùng để INSERT dữ liệu

```go
type User struct {
	ID        int `gorm:"primary_key"`
	Name      string
	Email     string
	Age       int
	IsActive  bool
	Average   float32
	CreatedAt time.Time
}
```

```go
type Post struct {
	ID        string `gorm:"primary_key"`
	Name      string `gorm:"type:varchar(50)"`
	Email     string `gorm:"type:varchar(100)"`
	Age       int    `gorm:"type:BIGINT"`
	IsActive  bool
	Average   float32 `gorm:"type:DECIMAL(6,2)"`
	CreatedAt time.Time
}
```

3. INSERT bản ghi

```go
var newUser = User{
	Name:      "Golang",
	Email:     "golang@goole.com",
	Age:       10,
	IsActive:  true,
	Average:   8.64,
	CreatedAt: time.Now(),
}
errCreateUser := db.Create(&newUser).Error
// Khi code API thì chỗ này trả về status 500 InternalServerError
if errCreateUser != nil {
	log.Println(errCreateUser)
	return
}

var newPost = Post{
	ID: "123abc",
	Name:      "Golang",
	Email:     "golang@goole.com",
	Age:       10,
	IsActive:  true,
	Average:   8.64,
	CreatedAt: time.Now(),
}
errCreatePost := db.Create(&newPost).Error
// Khi code API thì chỗ này trả về status 500 InternalServerError
if errCreatePost != nil {
	log.Println(errCreatePost)
	return
}
```