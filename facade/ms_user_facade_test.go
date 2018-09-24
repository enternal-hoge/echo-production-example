package facade

import (
	"github.com/enternal-hoge/echo-production-example/common"
	"github.com/enternal-hoge/echo-production-example/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMsUserFacade(t *testing.T) {

	t.Run("get include all values : format(userInfoResult []model.UserInfoResult)", func(t *testing.T) {

		uir := []model.UserInfoResult{}
		uf1 := model.UserInfoResult{}
		uf1.UserID = 1
		uf1.UserName = "fuga"
		uf1.Password = "fugafuga@123"
		uf1.PhoneNumber.String = "000-0000-0000"
		uf1.PhoneNumber.Valid = true
		uf1.Email.String = "fuga@outlook.jp"
		uf1.Email.Valid = true
		uf1.UpdateDate = common.NowUTC()

		uir = append(uir, uf1)
		uirj := format(uir)

		for _, i := range uirj {
			assert.Equal(t, 1, i.UserID, "value should be equal")
			assert.Equal(t, "fuga", i.UserName, "value should be equal")
			assert.Equal(t, "fugafuga@123", i.Password, "value should be equal")
			assert.Equal(t, "000-0000-0000", i.PhoneNumber, "value should be equal")
			assert.Equal(t, "fuga@outlook.jp", i.Email, "value should be equal")
		}
	})

	t.Run("get part of non values : format(userInfoResult []model.UserInfoResult)", func(t *testing.T) {

		uir := []model.UserInfoResult{}
		uf1 := model.UserInfoResult{}
		uf1.UserID = 1
		uf1.UserName = "fuga"
		uf1.Password = "fugafuga@123"
		uf1.PhoneNumber.String = ""
		uf1.PhoneNumber.Valid = false
		uf1.Email.String = ""
		uf1.Email.Valid = false
		uf1.UpdateDate = common.NowUTC()

		uir = append(uir, uf1)
		uirj := format(uir)

		for _, i := range uirj {
			assert.Equal(t, 1, i.UserID, "value should be equal")
			assert.Equal(t, "fuga", i.UserName, "value should be equal")
			assert.Equal(t, "fugafuga@123", i.Password, "value should be equal")
			assert.Equal(t, "", i.PhoneNumber, "value should be equal")
			assert.Equal(t, "", i.Email, "value should be equal")
		}
	})

}
