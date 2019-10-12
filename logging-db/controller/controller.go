package controller

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Controller struct {
	DB *gorm.DB
}

func(control *Controller) HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func(control *Controller) UploadVideo(c *gin.Context) {
	file, err := c.FormFile("video-file")
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		control.LogError(c, err.Error())
		return
	}

	fileType := file.Header["Content-Type"][0]

	if GetAllowFormat(fileType, allowedMediaType) == "" {
		c.JSON(http.StatusBadRequest, "Định dạng không hợp lệ. Yêu cầu: mp4")
		control.LogError(c, "Định dạng không hợp lệ. Yêu cầu: mp4")
		return
	}

	err = c.SaveUploadedFile(file, path.Join("./upload", file.Filename))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		control.LogError(c, err.Error())
		return
	}
}
