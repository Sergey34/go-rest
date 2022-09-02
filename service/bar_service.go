package service

import "go-rest/model"

type BarService struct {
}

func NewBarService() BarService {
	return BarService{}
}

func (srv BarService) Bar(body model.Bar) (model.Bar, error) {
	body.BarBar = body.BarBar + "_updated"
	return body, nil
}
