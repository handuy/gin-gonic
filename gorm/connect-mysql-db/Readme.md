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

3. Connect MySQL database
```go
db, err := gorm.Open("mysql", "root:123@(localhost:8080)/employees?charset=utf8&parseTime=True&loc=Local")
```
trong đó:
- user là root
- password là 123: chính là biến môi trường MYSQL_ROOT_PASSWORD=123 khi khởi tạo MySQL docker container
- employees là tên database cần connect đến