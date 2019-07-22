package miapi

import (
	"../../services/miapiserv"
	"../../utils/apierrors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

const (
	paramUserId = "userID"
	paramSiteId = "siteID"
	paramCategoryId= "categoryID"
	paramCountryId= "countryID"
)

var i int

func GetUser(c *gin.Context) {
	userID := c.Param(paramUserId)
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		apiError := apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		}
		c.JSON(apiError.Status, apiError)
		return
	}
	user, apiError := miapiserv.GetUserFromAPI(id)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.JSON(http.StatusOK, user)
}

func GetSite(c *gin.Context) {
	siteID := c.Param(paramSiteId)
	site, apiError := miapiserv.GetSiteFromAPI(siteID)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.JSON(http.StatusOK, site)
}

func GetSites(c *gin.Context) {
	sites, apiError := miapiserv.GetSitesFromAPI()
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.JSON(http.StatusOK, sites)
}

func GetCategory(c *gin.Context) {
	catID := c.Param(paramCategoryId)
	category, apiError := miapiserv.GetCategoryFromAPI(catID)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.JSON(http.StatusOK, category)
}

func GetCountry(c *gin.Context) {
	countryID := c.Param(paramCountryId)
	category, apiError := miapiserv.GetCountryFromAPI(countryID)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.JSON(http.StatusOK, category)
}


func GetResult(c *gin.Context) {
	userID := c.Param(paramUserId)
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		apiError := apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		}
		c.JSON(apiError.Status, apiError)
		return
	}
	result, apiError := miapiserv.GetResultFromAPI(id)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	//limiter := time.Tick(3 * time.Second)
	//<-limiter
	c.JSON(http.StatusOK, result)
}

func Limit(max int) gin.HandlerFunc {
	sema := make(chan struct{}, max)
	return func(c *gin.Context) {
		var called, fulled bool
		defer func() {
			limiter := time.Tick(3 * time.Second)
			<-limiter
			if called == false && fulled == false {
				<-sema
			}
		}()
		select {
			case sema <- struct{}{}:
				c.Next()
				called = true
				<-sema
			default:
				fulled = true
				c.Status(http.StatusBadGateway)
		}
	}
}