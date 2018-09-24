package route

import (
	"echo-production-example/common"
	"echo-production-example/facade"
	"echo-production-example/log"

	"net/http"

	"github.com/labstack/echo"
)

var appLog *log.AppLogger

func MsUserActivate(e *echo.Echo) {
	e.GET("/v1/MsUser", getMsUser)
	e.POST("/v1/MsUser", createMsUser)
	e.PUT("/v1/MsUser", updateMsUser)
	e.DELETE("/v1/MsUser", deleteMsUser)
	e.GET("/v1/MsUser/:user_id", getMsUserByPK)
	//external service
	e.GET("/v1/Hoge", getHoge)

}

func getMsUser(c echo.Context) error {
	appLog.Info("### route.getMsUser() called. ###")

	hs, rs := facade.MsUserFacade().FindAll()
	return c.String(hs, rs)

}

func createMsUser(c echo.Context) error {
	appLog.Info("### route.createMsUser() called. ###")

	json, err := common.ByteArray2String(c.Request().Body)
	if err != nil {
		panic(err)
	}

	hs, rs := facade.MsUserFacade().Insert(json)

	return c.String(hs, rs)
}

func updateMsUser(c echo.Context) error {
	appLog.Info("### route.updateMsUser() called. ###")

	json, err := common.ByteArray2String(c.Request().Body)
	if err != nil {
		panic(err)
	}

	hs, rs := facade.MsUserFacade().Update(json)

	return c.String(hs, rs)
}

func deleteMsUser(c echo.Context) error {
	appLog.Info("### route.deleteMsUser() called. ###")

	json, err := common.ByteArray2String(c.Request().Body)
	if err != nil {
		panic(err)
	}

	hs, rs := facade.MsUserFacade().Delete(json)

	return c.String(hs, rs)
}

func getMsUserByPK(c echo.Context) error {
	appLog.Info("### route.getMsUserByPK() called. ###")

	id := c.Param("userId")
	if id == "" {
		panic("error")
	}

	hs, rs := facade.MsUserFacade().FindByID(id)

	return c.String(hs, rs)
}

func getHoge(c echo.Context) error {
	appLog.Info("### route.getHoge() called. ###")

	rs := facade.MsUserFacade().FindHoge()

	return c.String(http.StatusOK, rs)
}
