package facade

import (
	"echo-production-example/dao"
	"fmt"
)

type msAuthorityFacade struct{}

type MsAuthorityFacadeIF interface {
	Insert(json string) string
	Update(json string) string
	Delete(json string) string
	FindAll() string
	FindByID(id string, name string) string
}

func MsAuthorityFacade() MsAuthorityFacadeIF {
	return &msAuthorityFacade{}
}

func (self *msAuthorityFacade) Insert(json string) string {
	fmt.Println("MsAuthorityFacade Insert invoke()")
	dao.MsAuthorityDao().Insert("hoge")
	return "hoge"
}

func (self *msAuthorityFacade) Update(json string) string {
	fmt.Println("MsAuthorityFacade Update invoke()")
	dao.MsAuthorityDao().Update("hoge")
	return "hoge"
}

func (self *msAuthorityFacade) Delete(json string) string {
	fmt.Println("MsAuthorityFacade Delete invoke()")
	dao.MsAuthorityDao().Delete("hoge")
	return "hoge"
}

func (self *msAuthorityFacade) FindAll() string {
	fmt.Println("MsAuthority FindAll invoke()")
	dao.MsAuthorityDao().FindAll()
	return "hoge"
}

func (self *msAuthorityFacade) FindByID(id string, name string) string {
	fmt.Println("MsAuthorityFacade FindByID invoke()")
	dao.MsAuthorityDao().FindByID("1", "hoge")
	return "hoge"
}
