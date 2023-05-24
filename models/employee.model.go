package models

import (
	"net/http"

	"echo_rest/db"
)

type Employee struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Phone_Number string `json:"phone_number"`
}

func FetchAllEmployee() (Response, error) {
	var obj Employee
	var arrobj []Employee
	var res Response

	con := db.CreatCon()
	sqlStatement := "select * from employee"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Address, &obj.Phone_Number)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Messages = "Success!"
	res.Data = arrobj

	return res, nil
}

func PostEmployee(name string, address string, phone_number string) (Response, error) {
	var res Response

	con := db.CreatCon()
	sqlStatement := "INSERT INTO employee (name, address, phone_number) VALUES($1, $2, $3)"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, address, phone_number)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Messages = "Success!"
	res.Data = map[string]interface{}{"lastID": result}

	return res, nil
}

func PutEmployee(id string, name string, address string, phone_number string) (Response, error) {
	var res Response

	con := db.CreatCon()
	sqlStatement := "UPDATE employee SET name=$1, address=$2, phone_number=$3 WHERE id=$4"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, address, phone_number, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Messages = "Success!"
	res.Data = map[string]interface{}{
		"name": name,
		"address": address,
		"phone_number": phone_number,
		"rowsAffected": rowsAffected,
	}

	return res, nil
}

func DelEmployee(id string) (Response, error) {
	var res Response

	con := db.CreatCon()
	sqlStatement := "DELETE FROM employee WHERE id=$1"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Messages = "Success!"
	res.Data = map[string]interface{}{
		"rowsAffected": rowsAffected,
	}

	return res, nil
}
