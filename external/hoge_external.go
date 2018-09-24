package external

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

type hogeExternal struct{}

type HogeExternalIF interface {
	FindHoge() string
}

func HogeExternal() HogeExternalIF {
	return &hogeExternal{}
}

func (self *hogeExternal) FindHoge() string {
	resp, body, errs := gorequest.New().Get("http://localhost:1323/v1/MsUser").End()

	if errs != nil {
		panic(errs)
	}

	fmt.Println(resp.StatusCode)
	fmt.Println(body)

	return body
}
