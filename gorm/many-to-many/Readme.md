## Tạo foreign key constraint cho cột trong bảng

1. Setup

Xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/connect-mysql-read-config

2. Định nghĩa struct tạo bảng

```go
type User struct {
	Id          string `gorm:"primary_key"`
	Name        string `gorm:"type:varchar(100)"`
	Age         int
}

type CreditCard struct {
	Number    string `gorm:"primary_key"`
	ExpiredAt time.Time
	UserId    string
}
```

3. Thêm foregin key constraint cho cột UserId ở bảng CreditCard, link đến primary key Id ở bảng User

```go
db.AutoMigrate(&User{})
db.AutoMigrate(&CreditCard{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
```

### Lưu ý

Đúng ra là chỉ cần đánh tag ForeignKey vào struct định nghĩa bảng là có thể tạo được foreign key constraint, 
tuy nhiên chưa tìm ra cách