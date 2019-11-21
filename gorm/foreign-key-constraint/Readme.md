## Tạo foreign key constraint cho cột trong bảng

1. Setup

Xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/connect-mysql-read-config

2. Định nghĩa struct tạo bảng

```go
type User struct {
	ID          string `gorm:"primary_key"`
	Name        string `gorm:"type:varchar(100)"`
	Age         int
}

type CreditCard struct {
	Number    string `gorm:"primary_key"`
	ExpiredAt time.Time
	UserID    string
}
```

3. Thêm foregin key constraint cho cột UserId ở bảng CreditCard, link đến primary key Id ở bảng User

```go
var user User
var creditCard CreditCard

errCreateUser := db.CreateTable(user).Error
if errCreateUser != nil {
	log.Println(errCreateUser)
	return
}

errCreateCard := db.CreateTable(creditCard).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").Error
if errCreateCard != nil {
	log.Println(errCreateCard)
	return
}
```

### Issue liên quan đến tạo Foreign Key constraint bằng tag

https://github.com/jinzhu/gorm/issues/450