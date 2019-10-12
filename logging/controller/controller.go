package controller

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func UploadVideo(c *gin.Context) {
	file, err := c.FormFile("video-file")
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		LogError(c, err.Error())
		return
	}

	fileType := file.Header["Content-Type"][0]

	if GetAllowFormat(fileType, allowedMediaType) == "" {
		c.JSON(http.StatusBadRequest, "Định dạng không hợp lệ. Yêu cầu: mp4")
		LogError(c, "Định dạng không hợp lệ. Yêu cầu: mp4")
		return
	}

	err = c.SaveUploadedFile(file, path.Join("./upload", file.Filename))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		LogError(c, err.Error())
		return
	}
}
