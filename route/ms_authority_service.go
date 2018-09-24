package route

import (
	"net/http"

	"echo-production-example/facade"

	"github.com/labstack/echo"
)

func MsAuthorityActivate(e *echo.Echo) {
	e.GET("/v1/MsAuthority", getMsAuthority)
	e.POST("/v1/MsAuthority", createMsAuthority)
	e.PUT("/v1/MsAuthority", updateMsAuthority)
	e.DELETE("/v1/MsAuthority", deleteMsAuthority)
	e.GET("/v1/MsAuthority/:authority_no", getMsAuthorityByPK)
}

func getMsAuthority(c echo.Context) error {
	facade := facade.MsAuthorityFacade()
	facade.FindAll()
	return c.JSON(http.StatusOK, "{ id: 111, name: 111}")
}

func createMsAuthority(c echo.Context) error {
	facade := facade.MsAuthorityFacade()
	facade.Insert("hoge")
	return c.JSON(http.StatusCreated, "{}")
}

func updateMsAuthority(c echo.Context) error {
	facade := facade.MsAuthorityFacade()
	facade.Update("hoge")
	return c.JSON(http.StatusOK, "{}")
}

func deleteMsAuthority(c echo.Context) error {
	facade := facade.MsAuthorityFacade()
	facade.Delete("hoge")
	return c.JSON(http.StatusOK, "{}")
}

func getMsAuthorityByPK(c echo.Context) error {
	/*
		name := c.Param("authority_no")
		fmt.Println("param : " + name)
		facade.MsAuthority().Update()
	*/
	facade := facade.MsAuthorityFacade()
	facade.FindByID("1", "hoge")
	return c.String(http.StatusOK, "hoge")
}
