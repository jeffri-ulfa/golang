package list_input_information

import (
	"log"
	"strconv"
	"../../db"
	initialize "../../initialize/permission_to_drive"
	"../../models"
)

type ModelsPermission_init models.DB_init

func (model ModelsPermission_init) ModelPermissionToDrive(totalData int) (arrJoin []initialize.Join, err error) {
	var join initialize.Join
	db := db.Connect()

	result, err := db.Query(`SELECT store_information.id_code_store, store_information.code_store, basic_information.employee_code, basic_information.first_name, 
	basic_information.last_name, commuting_basic_information.driver_license_expiry_date, commuting_basic_information.car_insurance_document_expiry_date, 
	commuting_basic_information.insurance_company, commuting_basic_information.personal_injury, commuting_basic_information.property_damage, commuting_basic_information.status_approve
	FROM store_information INNER JOIN general_information ON store_information.id_code_store = general_information.id_store_code 
	INNER JOIN basic_information ON basic_information.id_basic_information = general_information.id_basic_information
	INNER JOIN commuting_basic_information ON commuting_basic_information.id_general_information = general_information.id_general_information`)
	if err != nil {
		log.Println(err.Error())
	}
	defer db.Close()
	for result.Next() {
		if err := result.Scan(&join.Id_code_store, &join.Code_store, &join.Employee_code,
			 &join.First_name, &join.Last_name, &join.Driver_license_expiry_date, &join.Car_insurance_document_expiry_date, &join.Insurance_company,
			  &join.Personal_injury, &join.Property_damage, &join.Status_approve); err != nil {
			log.Fatal(err.Error())
		} else {
			arrJoin = append(arrJoin, join)
		}
	}

	return arrJoin, nil
}


func (model ModelsPermission_init) ModelPermissionToDriveSearch(Keyword string, page int ,limit int) (arrSearch []initialize.Join, err error, CheckData int) {
	var Search initialize.Join
	db := db.Connect()
	querylimit := ``
	if strconv.Itoa(page) == "" && strconv.Itoa(limit) == ""{
		querylimit = ``
	}else {
		pageacheck := strconv.Itoa((page-1)*limit)
		showadata := strconv.Itoa(limit)
		querylimit = ` LIMIT `+pageacheck+`,`+showadata
	}
	db.QueryRow(`SELECT count(*) FROM general_information INNER JOIN store_information ON general_information.id_store_code = store_information.id_code_store 
	INNER JOIN commuting_basic_information ON commuting_basic_information.id_general_information = general_information.id_general_information 
	INNER JOIN basic_information ON basic_information.id_basic_information = general_information.id_basic_information WHERE CONCAT_WS('', store_information.id_code_store, store_information.code_store,
	commuting_basic_information.driver_license_expiry_date, commuting_basic_information.car_insurance_document_expiry_date, commuting_basic_information.insurance_company,
	commuting_basic_information.personal_injury, commuting_basic_information.property_damage, commuting_basic_information.status_approve, basic_information.employee_code,
	basic_information.first_name, basic_information.last_name) LIKE ?`, `%` + Keyword + `%`).Scan(&CheckData)

	queryT := `SELECT store_information.id_code_store, store_information.code_store, commuting_basic_information.driver_license_expiry_date, commuting_basic_information.car_insurance_document_expiry_date, 
	commuting_basic_information.insurance_company, commuting_basic_information.personal_injury, commuting_basic_information.property_damage, commuting_basic_information.status_approve,
	basic_information.employee_code, basic_information.first_name, basic_information.last_name
	FROM general_information INNER JOIN store_information ON general_information.id_store_code = store_information.id_code_store 
	INNER JOIN commuting_basic_information ON commuting_basic_information.id_general_information = general_information.id_general_information 
	INNER JOIN basic_information ON basic_information.id_basic_information = general_information.id_basic_information WHERE CONCAT_WS('',store_information.code_store, basic_information.employee_code, basic_information.first_name, basic_information.last_name,  commuting_basic_information.driver_license_expiry_date, commuting_basic_information.car_insurance_document_expiry_date, 
	commuting_basic_information.insurance_company, commuting_basic_information.personal_injury, commuting_basic_information.property_damage, commuting_basic_information.status_approve) LIKE ?` +querylimit

	rows, err := db.Query(queryT, "%" + Keyword + "%")

	if err != nil {
		log.Print(err)
	}

	defer db.Close()
	for rows.Next() {
		if err := rows.Scan(&Search.Id_code_store, &Search.Code_store, &Search.Driver_license_expiry_date, &Search.Car_insurance_document_expiry_date, &Search.Insurance_company,
			&Search.Personal_injury, &Search.Property_damage, &Search.Status_approve, &Search.Employee_code,
			&Search.First_name, &Search.Last_name); err != nil {
			log.Fatal(err.Error())
		} else {
			arrSearch = append(arrSearch, Search)
		}
	}

	return arrSearch, nil, CheckData
}

func (Model1 ModelsPermission_init)UpdateDataPermissionToDrive(update *initialize.UpdatePermissionToDrive) (arrUpdate []initialize.UpdatePermissionToDrive, err error) {
	db := db.Connect()
	stmt, err := db.Prepare("UPDATE commuting_basic_information SET permitted_to_drive = ?, status_approve = ? WHERE id_commuting_basic_information = ?")
	if err != nil {
		log.Println(err.Error())
	}
	stmt.Exec(update.Permitted_to_drive, update.Status_approve, update.Id_commuting_basic_information)
	if err != nil {
		log.Println(err)
	}

	Excute := initialize.UpdatePermissionToDrive{
		Id_commuting_basic_information: update.Id_commuting_basic_information,
		Permitted_to_drive:             update.Permitted_to_drive,
		Status_approve:                 update.Status_approve,
	}

	arrUpdate = append(arrUpdate, Excute)

	return arrUpdate, nil
}
