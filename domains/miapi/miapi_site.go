package miapi

import (
	"../../utils/apierrors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Site struct {
	Id 					string `json:"id"`
	Name 				string `json:"name"`
	CountryID			string `json:"country_id"`
	DefaultCurrentID	string `json:"default_currency_id"`
}

type Sites []Site

//const urlSite = "https://api.mercadolibre.com/sites/"
const urlSite = "http://localhost:8081/sites/"

func (site *Site) Get(c chan *apierrors.ApiError) {
	if site.Id == "" {
		c <- &apierrors.ApiError{
			Message:"No se encontró el usuario",
			Status:http.StatusBadRequest,
		}
		return
	}

	url := fmt.Sprintf("%s%s", urlSite, site.Id)
	res, err := http.Get(url)
	if err != nil {
		c <- &apierrors.ApiError{
			Message:err.Error(),
			Status:http.StatusInternalServerError,
		}
		return
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c <- &apierrors.ApiError{
			Message:err.Error(),
			Status:http.StatusInternalServerError,
		}
		return
	}

	if err := json.Unmarshal(data, &site); err != nil {
		c <- &apierrors.ApiError{
			Message:err.Error(),
			Status:http.StatusInternalServerError,
		}
		return
	}
	c <- nil
}

func (sites *Sites) GetAll() *apierrors.ApiError {

	url := fmt.Sprintf("%s", urlSite)
	res, err := http.Get(url)
	if err != nil {
		return &apierrors.ApiError{
			Message:err.Error(),
			Status:http.StatusInternalServerError,
		}
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &apierrors.ApiError{
			Message:err.Error(),
			Status:http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal(data, &sites); err != nil {
		return &apierrors.ApiError{
			Message:err.Error(),
			Status:http.StatusInternalServerError,
		}
	}

	return nil
}
