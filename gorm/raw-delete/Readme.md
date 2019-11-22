## UPDATE dữ liệu bằng raw query

1. Setup

Xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/connect-mysql-read-config

2. DELETE dữ liệu bằng raw query

```go
errDelete := db.Exec(`
    DELETE FROM users
    WHERE id = ?
`, 2).Error
// Khi code API thì chỗ này trả về HTTP status code 500
if errDelete != nil {
    log.Println(errDelete)
    return
}
```