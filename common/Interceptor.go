package common

import (
	"echo-production-example/log"

	e "github.com/labstack/echo"
)

var traceLog *log.AppLogger

type TraceLog struct {
	ServiceName string `json:"serviceName"`
	TenantNo    string `json:"tenantNo"`
	LoginUserId string `json:"loginUserId"`
	MachineNo   string `json:"machineNo"`
	IpAddress   string `json:"ipAddress"`
}

func TraceLogInterceptor() e.MiddlewareFunc {
	return func(next e.HandlerFunc) e.HandlerFunc {
		return func(c e.Context) error {
			/* http intercepotor */
			al := log.GetApplogger()

			l := TraceLog{}
			l.ServiceName = c.Request().Header.Get("ServiceName")
			l.TenantNo = c.Request().Header.Get("TenantNo")
			l.LoginUserId = c.Request().Header.Get("LoginUserId")
			l.MachineNo = c.Request().Header.Get("MachineNo")
			l.IpAddress = c.Request().Header.Get("IpAddress")

			al.Trace(MarshalJson(l))

			err := next(c)
			/* after processing route execution */

			return err
		}
	}
}
