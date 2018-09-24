package model

import (
	"database/sql"
	"time"
)

type UserInfoResult struct {
	UserID      int            `db:"user_id"`
	UserName    string         `db:"user_name"`
	Password    string         `db:"passwd"`
	PhoneNumber sql.NullString `db:"phone_number"`
	Email       sql.NullString `db:"email"`
	Address     sql.NullString `db:"address"`
	UpdateDate  time.Time      `db:"update_date"`
}

type UserInfoResultJson struct {
	UserID      int       `json:"userId"`
	UserName    string    `json:"userName"`
	Password    string    `json:"password"`
	PhoneNumber string    `json:"phoneNumber"`
	Email       string    `json:"email"`
	Address     string    `json:"address"`
	UpdateDate  time.Time `json:"updateDate"`
}

func FormatUserInfo(userInfoResult []UserInfoResult) []UserInfoResultJson {
	userInfoList := []UserInfoResultJson{}

	for _, userInfo := range userInfoResult {
		json := UserInfoResultJson{}
		json.UserID = userInfo.UserID
		json.UserName = userInfo.UserName
		json.Password = userInfo.Password
		if userInfo.PhoneNumber.Valid {
			json.PhoneNumber = userInfo.PhoneNumber.String
		}
		if userInfo.Email.Valid {
			json.Email = userInfo.Email.String
		}
		if userInfo.Address.Valid {
			json.Address = userInfo.Address.String
		}
		json.UpdateDate = userInfo.UpdateDate

		userInfoList = append(userInfoList, json)
	}

	return userInfoList
}
