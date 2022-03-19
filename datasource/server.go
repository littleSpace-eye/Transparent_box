package datasource

import "github.com/deepmap/oapi-codegen/pkg/types"

type Record struct {
	ID int `sql:"id" json:"id" form:"id"`
	Name string `sql:"name" json:"name" form:"name"`
	Tel string `sql:"tel" json:"tel" form:"tel"`
	E_mail string `sql:"e_mail" json:"e_mail" form:"e_mail"`
	Password string `sql:"password",json:"password" form:"password"`
	Sex bool `sql:"sex" json:"sex" form:"sex"`
	Birth types.Date `sql:"birth" json:"birth" form:"birth"`
}
