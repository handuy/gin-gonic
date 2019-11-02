package controller

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gin-gonic/api-long-flutter/model"

	jwt_lib "github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

type userJSON struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (gc *Controller) Register(c *gin.Context) {
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	password := c.PostForm("password")

	if name == "" {
		c.JSON(http.StatusBadRequest, model.ErrorMesssage{
			Message: "Vui lòng nhập tên",
		})
		return
	}

	if phone == "" {
		c.JSON(http.StatusBadRequest, model.ErrorMesssage{
			Message: "Vui lòng nhập số điện thoại",
		})
		return
	}

	if len(password) < 4 {
		c.JSON(http.StatusBadRequest, model.ErrorMesssage{
			Message: "Mật khẩu phải có tối thiểu 4 kí tự",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMesssage{
			Message: "Lỗi xử lý server",
		})
	}

	passwordHash := string(hash)
	userID := xid.New().String()

	var newUser = model.User{
		ID:       userID,
		Name:     name,
		Phone:    phone,
		Avatar:   "",
		Address:  "",
		Password: passwordHash,
		Role:     1,
	}
	gc.DB.Create(newUser)

	// Tạo token với Header lưu thông tin chung:
	// Loại token: JWT
	// Thuật toán mã hoá: HS256
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))

	// Truyền dữ liệu vào phần Claim của token
	// Dữ liệu có kiểu map[string]interface{} mô phỏng một cấu trúc dạng JSON
	token.Claims = jwt_lib.MapClaims{
		"userId": userID,
		"Role":   1,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	}

	// Tạo Signature cho token
	// Signature = HS256(Header, Claim, mysupersecretpassword)
	// Sử dụng mysupersecretpassword như một input đầu vào
	// để thuật toán HS256 tạo ra chuỗi signature
	tokenString, err := token.SignedString([]byte(model.SecretKey))
	if err != nil {
		c.JSON(500, gin.H{"message": "Không tạo được token token"})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})

	var userInfo = model.UserInfo{
		UserType:    newUser.Role,
		UserProfile: newUser,
		Token:       tokenString,
	}

	var rsp = model.SignupLoginResponse{
		ResponseTime: time.Now().String(),
		Code:         0,
		Message:      "Đăng nhập thành công",
		Data:         userInfo,
	}

	c.JSON(http.StatusOK, rsp)
	return
}

func (gc *Controller) RegisterJSON(c *gin.Context) {
	var registerInfo userJSON
	err := c.BindJSON(&registerInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorMesssage{
			Message: "Form đăng kí không hợp lệ",
		})
		return
	}

	log.Println("------------", registerInfo)

	if registerInfo.Name == "" {
		c.JSON(http.StatusBadRequest, model.ErrorMesssage{
			Message: "Vui lòng nhập tên",
		})
		return
	}

	// TODO: Kiểm tra xem số điện thoại đã tồn tại trong hệ thống chưa
	if registerInfo.Phone == "" {
		c.JSON(http.StatusBadRequest, model.ErrorMesssage{
			Message: "Vui lòng nhập số điện thoại",
		})
		return
	}

	if len(registerInfo.Password) < 4 {
		c.JSON(http.StatusBadRequest, model.ErrorMesssage{
			Message: "Mật khẩu phải có tối thiểu 4 kí tự",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(registerInfo.Password), bcrypt.MinCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMesssage{
			Message: "Lỗi server",
		})
		return
	}
	passwordHash := string(hash)

	userID := xid.New().String()

	var newUser = model.User{
		ID:       userID,
		Name:     registerInfo.Name,
		Phone:    registerInfo.Phone,
		Avatar:   "",
		Address:  "",
		Password: passwordHash,
		Role:     1,
	}

	errInsertDb := gc.DB.Table("user").Create(newUser).Error
	if errInsertDb != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMesssage{
			Message: "Lỗi server",
		})
		return
	}

	// Tạo token với Header lưu thông tin chung:
	// Loại token: JWT
	// Thuật toán mã hoá: HS256
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))

	// Truyền dữ liệu vào phần Claim của token
	// Dữ liệu có kiểu map[string]interface{} mô phỏng một cấu trúc dạng JSON
	token.Claims = jwt_lib.MapClaims{
		"userId": userID,
		"Role":   1,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	}

	// Tạo Signature cho token
	// Signature = HS256(Header, Claim, mysupersecretpassword)
	// Sử dụng mysupersecretpassword như một input đầu vào
	// để thuật toán HS256 tạo ra chuỗi signature
	tokenString, err := token.SignedString([]byte(model.SecretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMesssage{
			Message: "Lỗi server",
		})
		return
	}

	var userInfo = model.UserInfo{
		UserType:    newUser.Role,
		UserProfile: newUser,
		Token:       tokenString,
	}

	var rsp = model.SignupLoginResponse{
		ResponseTime: time.Now().String(),
		Code:         0,
		Message:      "Đăng kí thành công",
		Data:         userInfo,
	}

	c.JSON(http.StatusOK, rsp)
	return
}

func (gc *Controller) Login(c *gin.Context) {
	phone := c.PostForm("phone")
	password := c.PostForm("password")

	// Kiểm tra
	var user model.User
	err := gc.DB.Raw(`
		SELECT *
		FROM user
		WHERE phone = ?
	`, phone).Scan(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.ErrorMesssage{
			Message: "Số điện thoại chưa được đăng kí trong hệ thống",
		})
		// log.Println(err)
		return
	}
	log.Println("---------------", user)

	byteHash := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(byteHash, []byte(password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.ErrorMesssage{
			Message: "Mật khẩu sai",
		})
		return
	}

	// Tạo token với Header lưu thông tin chung:
	// Loại token: JWT
	// Thuật toán mã hoá: HS256
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))

	// Truyền dữ liệu vào phần Claim của token
	// Dữ liệu có kiểu map[string]interface{} mô phỏng một cấu trúc dạng JSON
	token.Claims = jwt_lib.MapClaims{
		"userId": user.ID,
		"Role":   user.Role,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	}

	// Tạo Signature cho token
	// Signature = HS256(Header, Claim, mysupersecretpassword)
	// Sử dụng mysupersecretpassword như một input đầu vào
	// để thuật toán HS256 tạo ra chuỗi signature
	tokenString, err := token.SignedString([]byte(model.SecretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMesssage{
			Message: "Lỗi server",
		})
		return
	}

	var userInfo = model.UserInfo{
		UserType:    user.Role,
		UserProfile: user,
		Token:       tokenString,
	}

	var rsp = model.SignupLoginResponse{
		ResponseTime: time.Now().String(),
		Code:         0,
		Message:      "Đăng nhập thành công",
		Data:         userInfo,
	}

	c.JSON(http.StatusOK, rsp)
	return
}

func (gc *Controller) LoginJSON(c *gin.Context) {
	var loginInfo userJSON
	err := c.BindJSON(&loginInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorMesssage{
			Message: "Thông tin đăng nhập không hợp lệ",
		})
		return
	}

	//
	var user model.User
	err = gc.DB.Raw(`
		SELECT *
		FROM user
		WHERE phone = ?
	`, loginInfo.Phone).Scan(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.ErrorMesssage{
			Message: "Số điện thoại chưa được đăng kí trong hệ thống",
		})
		// log.Println(err)
		return
	}
	log.Println("---------------", user)

	byteHash := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(byteHash, []byte(loginInfo.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.ErrorMesssage{
			Message: "Mật khẩu sai",
		})
		return
	}

	// Tạo token với Header lưu thông tin chung:
	// Loại token: JWT
	// Thuật toán mã hoá: HS256
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))

	// Truyền dữ liệu vào phần Claim của token
	// Dữ liệu có kiểu map[string]interface{} mô phỏng một cấu trúc dạng JSON
	token.Claims = jwt_lib.MapClaims{
		"userId": user.ID,
		"Role":   user.Role,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	}

	// Tạo Signature cho token
	// Signature = HS256(Header, Claim, mysupersecretpassword)
	// Sử dụng mysupersecretpassword như một input đầu vào
	// để thuật toán HS256 tạo ra chuỗi signature
	tokenString, err := token.SignedString([]byte(model.SecretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMesssage{
			Message: "Lỗi server",
		})
		return
	}

	var userInfo = model.UserInfo{
		UserType:    user.Role,
		UserProfile: user,
		Token:       tokenString,
	}

	var rsp = model.SignupLoginResponse{
		ResponseTime: time.Now().String(),
		Code:         0,
		Message:      "Đăng nhập thành công",
		Data:         userInfo,
	}

	c.JSON(http.StatusOK, rsp)
	return
}

func (gc *Controller) ListIssues(c *gin.Context) {
	var headerInfo model.AuthorizationHeader

	if err := c.ShouldBindHeader(&headerInfo); err != nil {
		c.JSON(200, err)
	}

	tokenFromHeader := strings.Replace(headerInfo.Token, "Bearer ", "", -1)

	claims := jwt_lib.MapClaims{}
	tkn, err := jwt_lib.ParseWithClaims(tokenFromHeader, claims, func(token *jwt_lib.Token) (interface{}, error) {
		return []byte(model.SecretKey), nil
	})

	if err != nil {
		if err == jwt_lib.ErrSignatureInvalid {
			log.Println("error 1")
			c.JSON(http.StatusUnauthorized, model.ErrorMesssage{
				Message: "Token không hợp lệ",
			})
			return
		}
		log.Println("error 2", err)
		c.JSON(http.StatusBadRequest, model.ErrorMesssage{
			Message: "Bad request",
		})
		return
	}

	if !tkn.Valid {
		log.Println("error 3")
		c.JSON(http.StatusUnauthorized, model.ErrorMesssage{
			Message: "Token không hợp lệ",
		})
		return
	}

	var userId string
	var roleFromToken int

	for k, v := range claims {
		if k == "userId" {
			userId = v.(string)
		}

		if k == "Role" {
			roleFromToken = int(v.(float64))
		}
	}

	log.Println("--------", userId, roleFromToken)
	if userId == "" {
		c.JSON(http.StatusUnauthorized, model.ErrorMesssage{
			Message: "Token không hợp lệ",
		})
	}

	var issuesInfo []model.IssueGeneralInfo
	errGetIssues := gc.DB.Raw(`
		SELECT id, 
		CASE WHEN status = 0 THEN 'Chưa xử lý' WHEN 1 THEN 'Đang xử lý' ELSE 'Đã xử lý' END AS processed, 
		title, address, DATE_FORMAT(created_at, '%d-%m-%Y') AS date, TIME(created_at) AS time 
		FROM issue
		WHERE created_by = ?
	`, userId).Scan(&issuesInfo).Error
	if errGetIssues != nil {
		log.Println(errGetIssues)
		c.JSON(http.StatusInternalServerError, model.ErrorMesssage{
			Message: "Lỗi server",
		})
		return
	}

	var issueInfo = model.IssuesInfo{
		ResultCount: strconv.Itoa(len(issuesInfo)),
		Result:      issuesInfo,
	}

	var rsp = model.ListIssues{
		ResponseTime: time.Now().String(),
		Code:         0,
		Message:      "Lấy danh sách issue thành công",
		Data:         issueInfo,
	}

	c.JSON(http.StatusOK, rsp)
}


func (gc *Controller) IssueDetail(c *gin.Context) {
	var headerInfo model.AuthorizationHeader

	if err := c.ShouldBindHeader(&headerInfo); err != nil {
		c.JSON(200, err)
	}

	tokenFromHeader := strings.Replace(headerInfo.Token, "Bearer ", "", -1)

	claims := jwt_lib.MapClaims{}
	tkn, err := jwt_lib.ParseWithClaims(tokenFromHeader, claims, func(token *jwt_lib.Token) (interface{}, error) {
		return []byte(model.SecretKey), nil
	})

	if err != nil {
		if err == jwt_lib.ErrSignatureInvalid {
			log.Println("error 1")
			c.JSON(http.StatusUnauthorized, model.ErrorMesssage{
				Message: "Token không hợp lệ",
			})
			return
		}
		log.Println("error 2", err)
		c.JSON(http.StatusBadRequest, model.ErrorMesssage{
			Message: "Bad request",
		})
		return
	}

	if !tkn.Valid {
		log.Println("error 3")
		c.JSON(http.StatusUnauthorized, model.ErrorMesssage{
			Message: "Token không hợp lệ",
		})
		return
	}

	var userId string
	var roleFromToken int

	for k, v := range claims {
		if k == "userId" {
			userId = v.(string)
		}

		if k == "Role" {
			roleFromToken = int(v.(float64))
		}
	}

	log.Println("--------", userId, roleFromToken)
	if userId == "" {
		c.JSON(http.StatusUnauthorized, model.ErrorMesssage{
			Message: "Token không hợp lệ",
		})
	}

	issueId := c.Param("id")

	var issuesInfo model.IssueDetailInfo
	errGetIssues := gc.DB.Raw(`
		SELECT id, 
		CASE WHEN status = 0 THEN 'Chưa xử lý' WHEN 1 THEN 'Đang xử lý' ELSE 'Đã xử lý' END AS status, 
		title, content, address, DATE_FORMAT(created_at, '%d-%m-%Y') AS date, TIME(created_at) AS time, media
		FROM issue
		WHERE created_by = ?
		AND id = ?
	`, userId, issueId).Scan(&issuesInfo).Error
	if errGetIssues != nil {
		log.Println(errGetIssues)
		c.JSON(http.StatusInternalServerError, model.ErrorMesssage{
			Message: "Lỗi server",
		})
		return
	}

	var rsp = model.IssueDetailRsp{
		ResponseTime: time.Now().String(),
		Code:         0,
		Message:      "Lấy issue thành công",
		Data:         issuesInfo,
	}

	c.JSON(http.StatusOK, rsp)
}

func (gc *Controller) CreateIssue(c *gin.Context) {
	var headerInfo model.AuthorizationHeader

	if err := c.ShouldBindHeader(&headerInfo); err != nil {
		c.JSON(200, err)
	}

	tokenFromHeader := strings.Replace(headerInfo.Token, "Bearer ", "", -1)

	claims := jwt_lib.MapClaims{}
	tkn, err := jwt_lib.ParseWithClaims(tokenFromHeader, claims, func(token *jwt_lib.Token) (interface{}, error) {
		return []byte(model.SecretKey), nil
	})

	if err != nil {
		if err == jwt_lib.ErrSignatureInvalid {
			log.Println("error 1")
			c.JSON(http.StatusUnauthorized, model.ErrorMesssage{
				Message: "Token không hợp lệ",
			})
			return
		}
		log.Println("error 2", err)
		c.JSON(http.StatusBadRequest, model.ErrorMesssage{
			Message: "Bad request",
		})
		return
	}

	if !tkn.Valid {
		log.Println("error 3")
		c.JSON(http.StatusUnauthorized, model.ErrorMesssage{
			Message: "Token không hợp lệ",
		})
		return
	}

	var userId string
	var roleFromToken int

	for k, v := range claims {
		if k == "userId" {
			userId = v.(string)
		}

		if k == "Role" {
			roleFromToken = int(v.(float64))
		}
	}

	log.Println("--------", userId, roleFromToken)
	if userId == "" {
		c.JSON(http.StatusUnauthorized, model.ErrorMesssage{
			Message: "Token không hợp lệ",
		})
	}

	var createIssue model.CreateIssueReq
	err = c.BindJSON(&createIssue)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorMesssage{
			Message: "Yêu cầu tạo issue không hợp lệ",
		})
		return
	}

	var issueInsert = model.Issue{
		ID: xid.New().String(),
		Title: createIssue.Title,
		Content: createIssue.Content,
		Address: createIssue.Address,
		CreatedAt: time.Now(),
		Status: 0,
		Media: createIssue.Media,
		CreatedBy: userId,
	}

	errInsert := gc.DB.Table("issue").Create(&issueInsert).Error
	if errInsert != nil {
		log.Println(errInsert)
		c.JSON(http.StatusBadRequest, model.ErrorMesssage{
			Message: "Yêu cầu tạo issue không hợp lệ",
		})
		return
	}

	var rsp = model.CreateIssueRsp{
		ResponseTime: time.Now().String(),
		Code:         0,
		Message:      "Tạo issue thành công",
		Data:         issueInsert,
	}

	c.JSON(http.StatusOK, rsp)
}

func (gc *Controller) ProfileDetail(c *gin.Context) {
	var headerInfo model.AuthorizationHeader

	if err := c.ShouldBindHeader(&headerInfo); err != nil {
		c.JSON(200, err)
	}

	tokenFromHeader := strings.Replace(headerInfo.Token, "Bearer ", "", -1)

	claims := jwt_lib.MapClaims{}
	tkn, err := jwt_lib.ParseWithClaims(tokenFromHeader, claims, func(token *jwt_lib.Token) (interface{}, error) {
		return []byte(model.SecretKey), nil
	})

	if err != nil {
		if err == jwt_lib.ErrSignatureInvalid {
			log.Println("error 1")
			c.JSON(http.StatusUnauthorized, model.ErrorMesssage{
				Message: "Token không hợp lệ",
			})
			return
		}
		log.Println("error 2", err)
		c.JSON(http.StatusBadRequest, model.ErrorMesssage{
			Message: "Bad request",
		})
		return
	}

	if !tkn.Valid {
		log.Println("error 3")
		c.JSON(http.StatusUnauthorized, model.ErrorMesssage{
			Message: "Token không hợp lệ",
		})
		return
	}

	var userId string
	var roleFromToken int

	for k, v := range claims {
		if k == "userId" {
			userId = v.(string)
		}

		if k == "Role" {
			roleFromToken = int(v.(float64))
		}
	}

	log.Println("--------", userId, roleFromToken)
	if userId == "" {
		c.JSON(http.StatusUnauthorized, model.ErrorMesssage{
			Message: "Token không hợp lệ",
		})
	}

	issueId := c.Param("id")

	var issuesInfo model.IssueDetailInfo
	errGetIssues := gc.DB.Raw(`
		SELECT id, 
		CASE WHEN status = 0 THEN 'Chưa xử lý' WHEN 1 THEN 'Đang xử lý' ELSE 'Đã xử lý' END AS status, 
		title, content, address, DATE_FORMAT(created_at, '%d-%m-%Y') AS date, TIME(created_at) AS time, media
		FROM issue
		WHERE created_by = ?
		AND id = ?
	`, userId, issueId).Scan(&issuesInfo).Error
	if errGetIssues != nil {
		log.Println(errGetIssues)
		c.JSON(http.StatusInternalServerError, model.ErrorMesssage{
			Message: "Lỗi server",
		})
		return
	}

	var rsp = model.IssueDetailRsp{
		ResponseTime: time.Now().String(),
		Code:         0,
		Message:      "Lấy issue thành công",
		Data:         issuesInfo,
	}

	c.JSON(http.StatusOK, rsp)
}

