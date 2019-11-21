## SELECT dữ liệu từ bảng và đổ vào struct

1. Setup

Xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/connect-mysql-read-config

2. Ví dụ này sử dụng 3 bảng students, classes và student_classes. Chi tiết xem thêm tại: https://github.com/handuy/gin-gonic/tree/master/gorm/many-to-many

Sử dụng phpMyAdmin để INSERT dữ liệu mẫu:

```sql
insert into students values ("1", "Docker"), ("2", "Kubernetes"), ("3", "AWS"), ("4", "GCP")
insert into classes values ("a", "DevOps"), ("b", "Backend"), ("c", "Frontend")
insert into student_classes values ("1", "a"), ("1", "b"), ("2", "b"), ("2", "c"), ("3", "c"), ("4", "a");
```

3. Lấy dữ liệu từ 3 bảng đổ vào struct StudentInfo

```go
type StudentInfo struct {
	StudentId   string
	StudentName string
	ClassName   string
}

var studentInfo []StudentInfo
// Sử dụng column alias để cho tên bảng khớp với tên trường của struct hứng dữ liệu
errGetStudent := db.Table("student_classes").
	Joins("join students on student_classes.student_id = students.id").
	Joins("join classes on student_classes.class_id = classes.id").
	Select("students.id AS student_id, students.name AS student_name, classes.name AS class_name").
	Scan(&studentInfo).Error
// Khi code API thì chỗ này trả về HTTP status code 500
if errGetStudent != nil {
	log.Println(errGetStudent)
	return
}

log.Println("Student info", studentInfo)
```