package miapi

import (
	"../../services/miapi"
	"../../utils/apierrors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	paramUserId = "userID"
	paramSiteId = "siteID"
	paramCategoryId= "categoryID"
	paramCountryId= "countryID"
)

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
	user, apiError := miapi.GetUserFromAPI(id)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.JSON(http.StatusOK, user)
}

func GetSite(c *gin.Context) {
	siteID := c.Param(paramSiteId)
	site, apiError := miapi.GetSiteFromAPI(siteID)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.JSON(http.StatusOK, site)
}

func GetSites(c *gin.Context) {
	sites, apiError := miapi.GetSitesFromAPI()
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.JSON(http.StatusOK, sites)
}

func GetCategory(c *gin.Context) {
	catID := c.Param(paramCategoryId)
	category, apiError := miapi.GetCategoryFromAPI(catID)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.JSON(http.StatusOK, category)
}

func GetCountry(c *gin.Context) {
	countryID := c.Param(paramCountryId)
	category, apiError := miapi.GetCountryFromAPI(countryID)
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
	result, apiError := miapi.GetResultFromAPI(id)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.JSON(http.StatusOK, result)
}

