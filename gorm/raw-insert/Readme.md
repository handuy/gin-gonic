## UPDATE dữ liệu bằng raw query

1. Setup

Xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/connect-mysql-read-config

2. INSERT dữ liệu bằng raw query

```go
errInsert := db.Exec(`
    INSERT INTO users (name, email, age, is_active, average, created_at) VALUES 
    (?, ?, ?, ?, ?, ?)
`, "Azure", "azure@microsoft.com", 11, true, 9.87, time.Now()).Error
// Khi code API thì chỗ này trả về HTTP status code 500
if errInsert != nil {
    log.Println(errInsert)
    return
}
```