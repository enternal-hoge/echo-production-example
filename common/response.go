package common

import (
	"net/http"
)

type ResponseJson struct {
	Result  int         `json:"result"`
	Message interface{} `json:"message"`
}

func FormatResult(list interface{}, postFlg bool, err error) (int, ResponseJson) {

	hs := 500
	if postFlg {
		hs = http.StatusCreated
	} else {
		hs = http.StatusOK
	}

	rj := ResponseJson{}
	if err != nil {
		hs = http.StatusInternalServerError
		rj.Result = http.StatusInternalServerError
		rj.Message = err
	} else {
		rj.Result = hs
		rj.Message = list
	}

	return hs, rj
}
