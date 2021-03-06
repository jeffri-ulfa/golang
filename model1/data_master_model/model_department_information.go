package model1

import (
	"log"
	"strconv"

	"../../db"
	"../../initialize"
	"../../models"
)

type ModelDept_init models.DB_init

func (model1 ModelDept_init) ReadDataDepartmentInformation() (arrRead []initialize.DepartementInformation, err error) {
	var read initialize.DepartementInformation

	db := db.Connect()

	rows, err := db.Query("SELECT id_department, department_code, department_name, id_code_store FROM department_information")

	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&read.Id_department, &read.Department_code, &read.Department_name, &read.Id_code_store); err != nil {
			log.Println(err.Error())
		} else {
			arrRead = append(arrRead, read)
		}
	}

	return arrRead, nil
}

func (model1 ModelDept_init) GetDataDepartmentInformation(Id_department string) (arraDept []initialize.DepartementInformation, err error) {
	var depart initialize.DepartementInformation

	db := db.Connect()
	result, err := db.Query("SELECT id_department, department_code, department_name, id_code_store FROM department_information WHERE id_department = ?", Id_department)
	if err != nil {
		log.Println(err.Error())
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&depart.Id_department, &depart.Department_code, &depart.Department_name, &depart.Id_code_store)
		if err != nil {
			log.Println(err.Error())
		} else {
			arraDept = append(arraDept, depart)
		}
	}

	return arraDept, nil
}

func (model1 ModelDept_init) SearchDepartmentInformationModels(Keyword string, page int ,limit int) (arrSearch []initialize.DepartementInformation, err error, CheckData int) {
	var Search initialize.DepartementInformation
	db := db.Connect()
	querylimit := ``
	if strconv.Itoa(page) == "" && strconv.Itoa(limit) == ""{
		querylimit = ``
	}else {
		pageacheck := strconv.Itoa((page-1)*limit)
		showadata := strconv.Itoa(limit)
		querylimit = ` LIMIT `+pageacheck+`,`+showadata
	}
	db.QueryRow("SELECT count(*) FROM department_information WHERE CONCAT_WS('',id_department, department_code, department_name, id_code_store) LIKE ?", "%" + Keyword + "%").Scan(&CheckData)
	queryT := `SELECT id_department, department_code, department_name, id_code_store FROM department_information WHERE CONCAT_WS('',id_department, department_code, department_name, id_code_store) LIKE ?` +querylimit

	rows, err := db.Query(queryT, "%" + Keyword + "%")

	if err != nil {
		log.Print(err)
	}

	defer db.Close()
	for rows.Next() {
		if err := rows.Scan(&Search.Id_department, &Search.Department_code, &Search.Department_name, &Search.Id_code_store); err != nil {
			log.Fatal(err.Error())
		} else {
			arrSearch = append(arrSearch, Search)
		}
	}

	return arrSearch, nil, CheckData
}

func (model1 ModelDept_init) InsertDataDepartmentInformation(depart *initialize.DepartementInformation) (arraDept []initialize.DepartementInformation, err error) {
	db := db.Connect()
	stmt, err := db.Prepare("INSERT INTO department_information (department_code,department_name,id_code_store) VALUES (?,?,?)")
	if err != nil {
		log.Println(err.Error())
	}
	defer db.Close()

	_, err = stmt.Exec(depart.Department_code, depart.Department_name, depart.Id_code_store)
	if err != nil {
		log.Println(err.Error())
	}

	Excute := initialize.DepartementInformation{
		Department_code: depart.Department_code,
		Department_name: depart.Department_name,
		Id_code_store:   depart.Id_code_store,
	}

	arraDept = append(arraDept, Excute)

	return arraDept, nil

}

func (Model1 ModelDept_init) UpdateDataDepartmentInformation(update *initialize.DepartementInformation) (arrUpdate []initialize.DepartementInformation, err error) {

	db := db.Connect()

	stmt, err := db.Prepare("UPDATE department_information SET department_code = ?, department_name = ?, id_code_store = ? WHERE id_department = ?")
	if err != nil {
		log.Println(err.Error())
	}

	defer db.Close()

	_, err = stmt.Exec(update.Department_code, update.Department_name, update.Id_code_store, update.Id_department)
	if err != nil {
		log.Println(err.Error())
	}

	Excute := initialize.DepartementInformation{
		Id_department:   update.Id_department,
		Department_code: update.Department_code,
		Department_name: update.Department_name,
		Id_code_store:   update.Id_code_store,
	}

	arrUpdate = append(arrUpdate, Excute)

	return arrUpdate, nil
}
func (model1 ModelDept_init) DeleteDataDepartmentInformation(delete *initialize.DepartementInformation) (arrDelete []initialize.DepartementInformation, err error) {
	db := db.Connect()
	stmt, err := db.Prepare("DELETE FROM department_information WHERE id_department = ?")
	if err != nil {
		log.Println(err.Error())
	}

	_, err = stmt.Exec(delete.Id_department)
	if err != nil {
		log.Println(err.Error())
	}

	Excute := initialize.DepartementInformation{
		Id_department: delete.Id_department,
	}

	arrDelete = append(arrDelete, Excute)

	return arrDelete, nil
}
