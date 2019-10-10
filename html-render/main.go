package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Club struct {
	Name string
	Year int
	Nation string
}

func getHomePage(c *gin.Context) {
	var clubList = []Club{
		{
			Name: "Man Utd",
			Year: 1879,
			Nation: "England",
		},
		{
			Name: "Juventus",
			Year: 1894,
			Nation: "Italy",
		},
		{
			Name: "Bayern",
			Year: 1901,
			Nation: "Germany",
		},
	}

	var data = gin.H{
		"Name": "Juan Mata",
		"Age": 30,
		"CurrentClub": "Man Utd",
		"ClubList": clubList,
	}

	c.HTML(http.StatusOK, "index.html", data)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*")
	router.Static("/resources", "./resources")
	router.GET("/", getHomePage)
	router.Run(":8085")
}
