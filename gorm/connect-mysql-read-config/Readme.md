## Connect MySQL server

1. Khởi tạo MySQL server bằng Docker:
```docker
docker run --name learn-mysql -v /Users/duy/mysql-data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123 -p 8080:3306 -d mysql:latest
```

2. Go get package
```go
go get -u github.com/jinzhu/gorm
go get -u github.com/go-sql-driver/mysql
```

3. Tạo file config.local.json lưu thông tin kết nối database
```json
{
    "database": {
       "user": "root",
       "password": "123",
       "database": "employees",
       "address": "localhost:8080"
    }    
}
```

4. Tạo package model chứa hàm đọc file config.local.json và lưu vào struct

5. Connect MySQL database

```go
func ConnectDb(user string, password string, database string, address string) (*gorm.DB) {
	connectionInfo := fmt.Sprintf(`%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local`, user, password, address, database)

	db, err := gorm.Open("mysql", connectionInfo)
	if err != nil {
		panic(err)
	}
	return db
}
```