package common

import "time"

type Yaml struct {
	Mode  string
	Kind  string
	Mssql mssql
	Mysql mysql
	Echo  echo
}

type mssql struct {
	Server             string
	Port               int
	User               string
	Password           string
	Database           string
	SetMaxOpenConns    int
	SetMaxIdleConns    int
	SetConnMaxLifetime time.Duration
}

type mysql struct {
	ConnectionUri      string
	SetMaxOpenConns    int
	SetMaxIdleConns    int
	SetConnMaxLifetime time.Duration
}

type echo struct {
	Port   string
	Static string
}

type fluent struct {
	Path   string
	Rotate int
}
