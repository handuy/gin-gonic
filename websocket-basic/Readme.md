## Tạo websocket connection giữa client và server

1. Tạo một websocket endpoint đóng vai trò chuyển đổi (upgrage) HTTP connection thành Websocket connection
```go
router.GET("/ws", func(c *gin.Context) {
    controller.Wshandler(c.Writer, c.Request)
})
```

2. Code xử lý việc upgrade HTTP connection sang Websocket connection
```go
var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Wshandler(w http.ResponseWriter, r *http.Request) {
    // Biến HTTP connection thành Websocket connection
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
    }
    
    // Chạy vòng lặp vô tận cho đến khi kết nối websocket bị ngắt
	for {
        // Đọc dữ liệu được gửi vào connection
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
        }
        
        // Viết dữ liệu vào connection để trả về cho client
		conn.WriteMessage(t, msg)
	}
}
```