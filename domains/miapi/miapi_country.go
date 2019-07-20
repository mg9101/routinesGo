package miapi

import (
	"../../utils/apierrors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Country struct {
	Id 					string `json:"id"`
	Name 				string `json:"name"`
	Locale				string `json:"locale"`
	CurrencyID 			string `json:"currency_id"`
	DecimalSeparator 	string `json:"decimal_separator"`
	ThousandsSeparator 	string `json:"thousands_separator"`
	TimeZone 			string `json:"time_zone"`
	GeoInformation 		struct{
		Location 	struct{
			Latitude 	float64 `json:"latitude"`
			Longitude 	float64 `json:"longitude"`
		} `json:"location"`
	} `json:"geo_information"`
	States 				[]struct{
		Id 				string `json:"id"`
		Name 			string `json:"name"`
	} `json:"states"`

}

//const urlCountry = "https://api.mercadolibre.com/countries/"
const urlCountry = "http://localhost:8081/countries/"

func (country *Country) Get(c chan *apierrors.ApiError) {
	if country.Id == "" {
		c <- &apierrors.ApiError{
			Message:"No se encontró el país",
			Status:http.StatusBadRequest,
		}
		return
	}

	url := fmt.Sprintf("%s%s", urlCountry, country.Id)
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

	if err := json.Unmarshal(data, &country); err != nil {
		c <- &apierrors.ApiError{
			Message:err.Error(),
			Status:http.StatusInternalServerError,
		}
		return
	}
	c <- nil
}

