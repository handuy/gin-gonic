## UPDATE dữ liệu bằng raw query

1. Setup

Xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/connect-mysql-read-config

2. UPDATE dữ liệu bằng raw query

```go
rowAffected := db.Exec(`
    UPDATE employees
    SET first_name = 'Lionel', last_name = 'Messi', gender = 'M'
    WHERE emp_no = ?
`, 10001).RowsAffected

log.Println("row affected", rowAffected)
```