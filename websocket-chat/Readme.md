## Tạo websocket chat app đơn giản

1. Mỗi browser/tab khởi tạo một websocket connection đến server qua API GET /ws

2. Ở server: Mỗi khi có request gửi đến /ws, server sẽ:
- Upgrade HTTP --> socket connection
- Tạo một đối tượng **Client**
- Đăng kí **Client** mới tạo với Hub thông qua channel **Register** của Hub

3. Hub là gì ?
- Là một cấu trúc dữ liệu lưu thông tin các Client và message gửi đến
- Mỗi khi có Client mới, nó sẽ được đăng kí với Hub thông qua **Register** channel và được lưu vào trường **Clients**

4. Client sẽ liên tục lắng nghe từ socket connection. Mỗi khi có message mới Client sẽ gửi message này vào channel **Broadcast** của Hub

5. Hub khi nhận thấy channel **Broadcast** nhận được message thì sẽ lưu message này vào trường **Content**, sau đó đẩy message này vào channel **send** của toàn bộ các Client được lưu trong Hub

6. Mỗi Client khi nhận được message ở channel **send** sẽ viết message này vào socket connection để hiển thị cho trình duyệt