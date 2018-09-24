package dao

import (
	"fmt"
)

type msAuthorityDao struct{}

type MsAuthorityDaoIF interface {
	Insert(json string) int
	Update(json string) int
	Delete(json string) int
	FindAll() string
	FindByID(id string, name string) string
}

func MsAuthorityDao() MsAuthorityDaoIF {
	return &msAuthorityDao{}
}

func (self *msAuthorityDao) Insert(json string) int {
	return 1
}

func (self *msAuthorityDao) Update(json string) int {
	fmt.Println("MsAuthorityDao Update invoke()")
	return 1
}

func (self *msAuthorityDao) Delete(json string) int {
	fmt.Println("MsAuthorityDao Delete invoke()")
	return 1
}

func (self *msAuthorityDao) FindAll() string {
	fmt.Println("MsAuthorityDao FindAll invoke()")
	return "hoge"
}

func (self *msAuthorityDao) FindByID(id string, name string) string {
	fmt.Println("MsAuthorityDao FindByID invoke()")
	return "hoge"
}
