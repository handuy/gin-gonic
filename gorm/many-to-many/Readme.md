## Tạo foreign key constraint cho cột trong bảng

1. Setup

Xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/connect-mysql-read-config

2. Định nghĩa struct tạo bảng, trong đó StudentClass đóng vai trò tạo bảng trung gian student_classes

```go
type Student struct {
	Id   string `gorm:"primary_key"`
	Name string `gorm:"type:varchar(100)"`
}

type Class struct {
	Id   string `gorm:"primary_key"`
	Name string
}

type StudentClass struct {
	StudentId string `gorm:"primary_key"`
	ClassId   string `gorm:"primary_key"`
}
```

3. Tạo bảng và tạo foreign key cho bảng trung gian

```go
var student Student
var class Class
var studentClass StudentClass

errCreateStudent := db.AutoMigrate(student).Error
if errCreateStudent != nil {
	log.Println(errCreateStudent)
	return
}

errCreateClass := db.AutoMigrate(class).Error
if errCreateClass != nil {
	log.Println(errCreateClass)
	return
}

errCreateStudentClass := db.AutoMigrate(studentClass).
			AddForeignKey("student_id", "students(id)", "RESTRICT", "RESTRICT").
			AddForeignKey("class_id", "classes(id)", "RESTRICT", "RESTRICT").Error
if errCreateStudentClass != nil {
	log.Println(errCreateStudentClass)
	return
}
```

### Issue liên quan đến tạo Foreign Key constraint bằng tag

https://github.com/jinzhu/gorm/issues/450