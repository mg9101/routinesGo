package miapi

import (
	"../../utils/apierrors"
	"sync"
)

type Result struct {
	User *User
	Site *Site
	Country *Country
}

func (result *Result) Get(id int64) *apierrors.ApiError {
	var wg sync.WaitGroup
	user := &User{
		ID: id,
	}
	result.User = user
	result.User.Get()
	result.Country = &Country{Id:result.User.CountryID}
	result.Site = &Site{Id:result.User.SiteID}
	c := make(chan *apierrors.ApiError)
	defer close(c)
	wg.Add(1)
	go result.Country.Get(c)

	wg.Add(1)
	go result.Site.Get(c)


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
		return  err
	}
	wg.Wait()
	return nil
}