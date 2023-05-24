package models

import (
	"database/sql"
	"fmt"

	"echo_rest/db"
	"echo_rest/helpers"
)

type User struct {
	Id       string `json:"id"`
	UserName string `json:"UserName"`
	Password string `json:"Password"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string

	con := db.CreatCon()
	sqlStatement := "SELECT * FROM users WHERE username =$1"
	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.Id, &obj.UserName, &pwd,
	)
	if err == sql.ErrNoRows {
		fmt.Println("Username not Found!")
		return false, err
	}
	if err != nil {
		fmt.Println("Querry Error!")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Password and Hash doesn't match!")
		return false, err
	}

	return true, nil
}
