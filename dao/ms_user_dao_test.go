package dao

import (
	"echo-production-example/common"
	"echo-production-example/model"
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

const ymlPrefix = "config" //config.yml

var yml model.Yaml

func init() {
	viper.SetConfigName(ymlPrefix)
	viper.AddConfigPath("../")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Errorf("Error reading config file : ")
	}

	err := viper.Unmarshal(&yml)
	if err != nil {
		fmt.Errorf("unable to decode into struct : ")
	}
}
func TestMsUserDao(t *testing.T) {

	t.Run("select : FindAll()", func(t *testing.T) {
		_, err := common.SqlxInit(yml)
		if err != nil {
			panic(err)
		}

		dao := MsUserDao()
		userInfoResult, err := dao.FindAll()
		if err != nil {
			panic(err)
		}

		assert.Equal(t, 1, userInfoResult[0].UserID, "first record column value should be equal")
		assert.Equal(t, "eternal_hoge", userInfoResult[0].UserName, "first record column value should be equal")
		assert.Equal(t, "hoge@123", userInfoResult[0].Password, "first record column value should be equal")
		assert.Equal(t, "03-0000-0000", userInfoResult[0].PhoneNumber.String, "first record column value should be equal")
		assert.Equal(t, true, userInfoResult[0].PhoneNumber.Valid, "first record column value should be equal")
		assert.Equal(t, "eternal-hoge@gmail.com", userInfoResult[0].Email.String, "first record column value should be equal")
		assert.Equal(t, true, userInfoResult[0].Email.Valid, "first record column value should be equal")
		assert.Equal(t, "unknown", userInfoResult[0].Address.String, "first record column value should be equal")
		assert.Equal(t, true, userInfoResult[0].Address.Valid, "first record column value should be equal")

		common.SqlxClose()
	})

	t.Run("insert : Insert(json string)", func(t *testing.T) {
		_, err := common.SqlxInit(yml)
		if err != nil {
			panic(err)
		}

		json := `{ "userId" : 5, "userName" : "page" , "password" : "pagepage@123", "phoneNumber" : "090-0000-0000", "email" : "page@outlook.com"}`
		dao := MsUserDao()
		err = dao.Insert(json)
		if err != nil {
			panic(err)
		}

		assert.Equal(t, nil, err, "first record column value should be equal")

		common.SqlxClose()
	})

	t.Run("Uppate : Updte(json string)", func(t *testing.T) {
		_, err := common.SqlxInit(yml)
		if err != nil {
			panic(err)
		}

		json := `{ "userId" : 5, "userName" : "pagepage" , "password" : "pagepage@123", "phoneNumber" : "090-0000-0000", "email" : "page@outlook.com", "address" : "unknown"}`
		dao := MsUserDao()
		err = dao.Update(json)
		if err != nil {
			panic(err)
		}

		assert.Equal(t, nil, err, "first record column value should be equal")

		common.SqlxClose()
	})

	t.Run("Delete : Delete(json string)", func(t *testing.T) {
		_, err := common.SqlxInit(yml)
		if err != nil {
			panic(err)
		}

		json := `{ "userId" : 5 }`
		dao := MsUserDao()
		err = dao.Delete(json)
		if err != nil {
			panic(err)
		}

		assert.Equal(t, nil, err, "first record column value should be equal")

		common.SqlxClose()
	})

}
