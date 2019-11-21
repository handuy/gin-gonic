## SELECT dữ liệu từ bảng và đổ vào struct

1. Setup

Xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/connect-mysql-read-config

2. Ví dụ này sử dụng 2 bảng users và credit_cards. Chi tiết xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/foreign-key-constraint

Sử dụng phpMyAdmin để INSERT dữ liệu mẫu:

```sql
insert into users values ("123", "Golang", 10)
insert into credit_cards VALUES ("abc", "2008-06-01 13:06:01", "123"), ("def", "2009-03-23 10:14:01", "123")
```

3. Lấy dữ liệu từ 2 bảng đổ vào struct UserInfo

```go
type UserInfo struct {
	UserId     string
	UserName   string
	CardNumber string
	ExpiredAt  time.Time
}

var userInfo []UserInfo
// Sử dụng column alias để cho tên bảng khớp với tên trường của struct hứng dữ liệu
errGetUser := db.Table("users").
	Joins("join credit_cards on users.id = credit_cards.user_id").
	Select("users.id AS user_id, users.name AS user_name, credit_cards.number AS card_number, credit_cards.expired_at AS expired_at").
	Scan(&userInfo).Error
// Khi code API thì chỗ này trả về HTTP status code 500
if errGetUser != nil {
	log.Println(errGetUser)
	return
}

log.Println("User info", userInfo)
```