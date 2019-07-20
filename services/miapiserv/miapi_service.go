package miapiserv

import (
	"../../domains/miapi"
	"../../utils/apierrors"
	"sync"
)

func GetUserFromAPI(id int64) (*miapi.User, *apierrors.ApiError){
	user := &miapi.User{
		ID: id,
	}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user , nil
}

func GetSiteFromAPI(id string) (*miapi.Site, *apierrors.ApiError){
	site := &miapi.Site	{
		Id: id,
	}
	var wg sync.WaitGroup
	wg.Add(1)
	c := make(chan *apierrors.ApiError)
	site.Get(c)
	var err *apierrors.ApiError
	go func() {
		for e := range c	{
			wg.Done()
			if e != nil {
				err = e
			}
		}
	}()
	if err != nil {
		return nil, err
	}
	wg.Wait()
	return site , nil
}

func GetSitesFromAPI() (*miapi.Sites, *apierrors.ApiError){
	sites := &miapi.Sites{}
	if err := sites.GetAll(); err != nil {
		return nil, err
	}
	return sites , nil
}
func GetCategoryFromAPI(id string) (*miapi.Category, *apierrors.ApiError){
	cat := &miapi.Category	{
		Id: id,
	}
	if err := cat.Get(); err != nil {
		return nil, err
	}
	return cat , nil
}

func GetCountryFromAPI(id string) (*miapi.Country, *apierrors.ApiError){
	country := &miapi.Country	{
		Id: id,
	}
	var wg sync.WaitGroup
	wg.Add(1)
	c := make(chan *apierrors.ApiError)
	country.Get(c)
	var err *apierrors.ApiError
	go func() {
		for e := range c	{
			wg.Done()
			if e != nil {
				err = e
			}
		}
	}()
	if err != nil {
		return nil, err
	}
	wg.Wait()
	return country , nil
}


func GetResultFromAPI(id int64) (*miapi.Result, *apierrors.ApiError){

	result := &miapi.Result{}
	if err := result.Get(id); err != nil {
		return nil, err
	}
	return result , nil
}
