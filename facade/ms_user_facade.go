package facade

import (
	"github.com/enternal-hoge/echo-production-example/common"
	"github.com/enternal-hoge/echo-production-example/dao"
	ext "github.com/enternal-hoge/echo-production-example/external"
	"github.com/enternal-hoge/echo-production-example/log"
	"github.com/enternal-hoge/echo-production-example/model"
)

var appLog *log.AppLogger

type msUserFacade struct{}

type MsUserFacadeIF interface {
	Insert(json string) (int, string)
	Update(json string) (int, string)
	Delete(json string) (int, string)
	FindAll() (int, string)
	FindByID(userId string) (int, string)
	FindHoge() string //external service
}

func MsUserFacade() MsUserFacadeIF {
	return &msUserFacade{}
}

func (self *msUserFacade) Insert(json string) (int, string) {
	appLog.Info("### facade.Insert(json string) called. ###")

	err := dao.MsUserDao().Insert(json)
	hs, rs := common.FormatResult(nil, true, err)

	return hs, common.MarshalJson(rs)
}

func (self *msUserFacade) Update(json string) (int, string) {
	appLog.Info("### facade.Update(json string) called. ###")

	err := dao.MsUserDao().Update(json)
	hs, rs := common.FormatResult(nil, true, err)

	return hs, common.MarshalJson(rs)
}

func (self *msUserFacade) Delete(json string) (int, string) {
	appLog.Info("### facade.Delete(json string) called. ###")

	err := dao.MsUserDao().Delete(json)
	hs, rs := common.FormatResult(nil, true, err)

	return hs, common.MarshalJson(rs)
}

func (self *msUserFacade) FindAll() (int, string) {
	appLog.Info("### facade.FindAll() called. ###")

	uir, err := dao.MsUserDao().FindAll()
	uil := model.FormatUserInfo(uir)
	hs, rs := common.FormatResult(uil, true, err)

	return hs, common.MarshalJson(rs)
}

func (self *msUserFacade) FindByID(userId string) (int, string) {
	appLog.Info("### facade.FindByID(userId string) called. ###")

	uir, err := dao.MsUserDao().FindByID(userId)
	uil := model.FormatUserInfo(uir)
	hs, rs := common.FormatResult(uil, true, err)

	return hs, common.MarshalJson(rs)
}

func (self *msUserFacade) FindHoge() string {
	appLog.Info("### facade.FindHoge() called. ###")

	res := ext.HogeExternal().FindHoge()
	//ToDo : JSON　Tranceform
	return res
}
