package main

import (
	"./controllers/miapi"
	"github.com/gin-gonic/gin"
)

const (
	port = ":8080"
)

var (
	router = gin.Default()
)

func main() {
	router.GET("/miapi/user/:userID", miapi.GetUser)
	router.GET("/miapi/site/:siteID", miapi.GetSite)
	router.GET("/miapi/sites", miapi.GetSites)
	router.GET("/miapi/category/:categoryID", miapi.GetCategory)
	router.GET("/miapi/country/:countryID", miapi.GetCountry)
	router.GET("/miapi/result/:userID", miapi.GetResult)

	router.Run(port)
}
