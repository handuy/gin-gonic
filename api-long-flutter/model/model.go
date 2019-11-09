package model

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	jsoniter "github.com/json-iterator/go"
)

const SecretKey string = "goldenratio1618"

type AuthorizationHeader struct {
	Token  string `header:"Authorization"`
}

type ErrorMesssage struct {
	Message string `json:"message"`
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	Address  string `json:"address"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}

type UserInfo struct {
	UserType    int    `json:"userType"`
	UserProfile User   `json:"userProfile"`
	Token       string `json:"token"`
}

type SignupLoginResponse struct {
	ResponseTime string   `json:"responseTime"`
	Code         int      `json:"code"`
	Message      string   `json:"message"`
	Data         UserInfo `json:"data"`
}

type Issue struct {
	ID         string     `json:"id"`
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	Address    string     `json:"address"`
	CreatedAt  time.Time  `json:"created_at"`
	Status     int        `json:"status"`
	Media      []string   `json:"media"`
	CreatedBy  string     `json:"created_by"`
}

type IssueGeneralInfo struct {
	ID          string    `json:"id"`
	Status      string    `json:"status"`
	Title       string    `json:"title"`
	Address     string    `json:"add"`
	Time        string    `json:"time"`
	Date        string    `json:"date"`
}

type IssueDetailInfo struct {
	ID          string    `json:"id"`
	Status      string    `json:"status"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Address     string    `json:"add"`
	Time        string    `json:"time"`
	Date        string    `json:"date"`
	Media       string    `json:"media"`
}

type IssuesInfo struct {
	ResultCount string             `json:"resultCount"`
	Result      []IssueGeneralInfo `json:"result"`
}

type ListIssues struct {
	ResponseTime string     `json:"responseTime"`
	Code         int        `json:"code"`
	Message      string     `json:"message"`
	Data         IssuesInfo `json:"data"`
}

type IssueDetailRsp struct {
	ResponseTime string          `json:"responseTime"`
	Code         int             `json:"code"`
	Message      string          `json:"message"`
	Data         IssueDetailInfo `json:"data"`
}

type CreateIssueReq struct {
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	Address    string     `json:"add"`
	Status     int        `json:"status"`
	Media      []string   `json:"media"`
}

type CreateIssueRsp struct {
	ResponseTime string          `json:"responseTime"`
	Code         int             `json:"code"`
	Message      string          `json:"message"`
	Data         Issue           `json:"data"`
} 

type Config struct {
	Database struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
		Address  string `json:"address"`
	} `json:"database"`
}

type File struct {
	Id           int        `json:"id"`
	Name         string     `json:"name"`
	Size         int64        `json:"size"`
	Type         string     `json:"type"`
	UploadedAt   time.Time  `json:"uploaded_at"`
	Url          string     `json:"url"`
	UploadedBy   string     `json:"uploaded_by"`
}

type UploadFileResponse struct {
	ResponseTime string   `json:"responseTime"`
	Code         int      `json:"code"`
	Message      string   `json:"message"`
	Data         string   `json:"data"`
}

func DecodeDataFromJsonFile(f *os.File, data interface{}) error {
	jsonParser := jsoniter.NewDecoder(f)
	err := jsonParser.Decode(&data)
	if err != nil {
		return err
	}

	return nil
}

func SetupConfig() Config {
	var conf Config

	// Đọc file config.dev.json
	configFile, err := os.Open("config.local.json")
	if err != nil {
		// Nếu không có file config.dev.json thì đọc file config.default.json
		configFile, err = os.Open("config.default.json")
		if err != nil {
			panic(err)
		}
		defer configFile.Close()
	}
	defer configFile.Close()

	// Parse dữ liệu JSON và bind vào conf
	err = DecodeDataFromJsonFile(configFile, &conf)
	if err != nil {
		log.Println("Không đọc được file config.")
		panic(err)
	}

	return conf
}

func ConnectDb(user string, password string, database string, address string) *gorm.DB {
	connectionInfo := fmt.Sprintf(`%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local`, user, password, address, database)

	db, err := gorm.Open("mysql", connectionInfo)
	if err != nil {
		panic(err)
	}
	return db
}
