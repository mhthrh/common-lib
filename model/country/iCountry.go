package country

import (
	cErrModel "github.com/mhthrh/GoNest/model/error"
)

type ICountry interface {
	Load() *cErrModel.XError
	Countries() ([]Country, *cErrModel.XError)
	GetByName(name string) (*Country, *cErrModel.XError)
	GetByCode(code string) (*Country, *cErrModel.XError)
}
type Country struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
