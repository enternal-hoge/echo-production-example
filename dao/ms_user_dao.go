package dao

import (
	"echo-production-example/common"
	"echo-production-example/log"
	"echo-production-example/model"

	"github.com/tidwall/gjson"
)

var appLog *log.AppLogger

type msUserDao struct{}

type MsUserDaoIF interface {
	Insert(json string) error
	Update(json string) error
	Delete(json string) error
	FindAll() ([]model.UserInfoResult, error)
	FindByID(userId string) ([]model.UserInfoResult, error)
}

func MsUserDao() MsUserDaoIF {
	return &msUserDao{}
}

func (self *msUserDao) Insert(json string) error {
	appLog.Info("### dao.Insert(json string) called. ###")

	userID := gjson.Get(json, "userId").Int()
	userName := gjson.Get(json, "userName").String()
	password := gjson.Get(json, "password").String()
	phoneNumber := gjson.Get(json, "phoneNumber").String()
	email := gjson.Get(json, "email").String()
	address := gjson.Get(json, "address").String()

	updateDate := common.NowUTC()

	db := common.SqlxConnect()
	query := `INSERT INTO
							user_info
							(
								user_id,user_name, passwd, phone_number, email, address, update_date
							) VALUES (
								?, ?, ?, ?, ?, ?, ?
							)`
	_, err := db.Exec(query, userID, userName, password, phoneNumber, email, address, updateDate)

	return err
}

func (self *msUserDao) Update(json string) error {
	appLog.Info("### dao.Update(json string) called. ###")

	userID := gjson.Get(json, "userId").Int()
	userName := gjson.Get(json, "userName").String()
	password := gjson.Get(json, "password").String()
	phoneNumber := gjson.Get(json, "phoneNumber").String()
	email := gjson.Get(json, "email").String()
	address := gjson.Get(json, "address").String()

	updateDate := common.NowUTC()

	db := common.SqlxConnect()
	q := `UPDATE
					user_info
				SET
					user_name = ?,
					passwd = ?,
					phone_number = ?,
					email = ?,
					address = ?,
					update_date = ?
				WHERE
				user_id = ?`
	_, err := db.Exec(q, userName, password, phoneNumber, email, address, updateDate, userID)

	return err
}

func (self *msUserDao) Delete(json string) error {
	appLog.Info("### dao.Delete(json string) called. ###")

	userID := gjson.Get(json, "userId").Int()
	db := common.SqlxConnect()

	q := `DELETE
				FROM
					user_info
				WHERE
					user_id = ?
				`
	_, err := db.Exec(q, userID)

	return err

}

func (self *msUserDao) FindAll() ([]model.UserInfoResult, error) {
	appLog.Info("### dao.FindAll() called. ###")

	uir := []model.UserInfoResult{}
	db := common.SqlxConnect()

	q := `SELECT
					user_id,
					user_name,
					passwd,
					phone_number,
					email,
					address,
					update_date
				FROM
					user_info
				`
	err := db.Select(&uir, q)

	return uir, err
}

func (self *msUserDao) FindByID(userId string) ([]model.UserInfoResult, error) {
	appLog.Info("### dao.FindByID(json string) called. ###")

	uir := []model.UserInfoResult{}
	db := common.SqlxConnect()

	q := `SELECT
					user_id,
					user_name,
					passwd,
					phone_number,
					email,
					address,
					update_date
				FROM
					user_info
				WHERE
					user_id = ?`
	err := db.Select(&uir, q, common.String2Int(userId))

	return uir, err
}
