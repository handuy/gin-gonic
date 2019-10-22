## Sử dụng redis để lưu session ID

1. Xem trước ví dụ về session-cookie tại:
https://github.com/handuy/gin-gonic/tree/master/session-cookie

2. Tạo một Redis database chạy ở local bằng Docker:
```docker
docker run -d -p 6379:6379 -v $HOME/redis-data:/data --name local-redis redis redis-server --requirepass 123
```

Trong quá trình chạy, nếu gặp lỗi:
```bash
cannot find package "github.com/boj/redistore"
cannot find package "github.com/gomodule/redigo/redis"
```
thì chạy lệnh sau để lấy package về:
```go
go get -u github.com/boj/redistore github.com/gomodule/redigo/redis
```

3. Thay vì lưu thông tin về session ID ở cookie thì ta lưu ở Redis database, còn lại logic xử lý tương tự như khi lưu ở cookie:
- Client gửi request đăng nhập
- Server check thông tin đăng nhập (email, password). Nếu OK thì server set một cặp key-value =  **một chuỗi UUID tự sinh** - struct **Account** lưu thông tin user
```go
type Account struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
```
- Chuỗi UUID ở tạo ra ở bước trên sẽ được trả về client. Client sẽ lưu nó vào trong cookiecookie
- Ở những lần request sau, client đều sẽ gửi cookie chứa UUID lên cho server. Server lấy UUID và kiểm tra trong Redis database. Nếu UUID có tồn tại trong Redis thì client được tính là đã đăng nhập thành công 