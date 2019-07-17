package miapi


import (
	"../../utils/apierrors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Category struct {
	Id 					string `json:"id"`
	Name 				string `json:"name"`
}

const urlCategory = "https://api.mercadolibre.com/categories/"

func (category *Category) Get() *apierrors.ApiError {
	if category.Id == "" {
		return &apierrors.ApiError{
			Message:"No se encontró la categoría",
			Status:http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%s", urlCategory, category.Id)
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

	if err := json.Unmarshal(data, &category); err != nil {
		return &apierrors.ApiError{
			Message:err.Error(),
			Status:http.StatusInternalServerError,
		}
	}

	return nil
}

