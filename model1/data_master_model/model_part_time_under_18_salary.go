package model1

import (
	"log"
"strconv"
	"../../db"
	"../../initialize"
	"../../models"
)

type ModelUnder_init models.DB_init

func (model1 ModelUnder_init) ReturnAllDataUnder18() (arrAll []initialize.PartTimeUnder18Salary, err error) {
	var all initialize.PartTimeUnder18Salary
	db := db.Connect()

	rows, err := db.Query("SELECT id_part_time_under_18_salary, id_code_store, salary FROM part_time_under_18_salary")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&all.Id_part_time_under_18_salary, &all.Id_code_store, &all.Salary); err != nil {
			log.Println(err.Error())
		} else {
			arrAll = append(arrAll, all)
		}
	}
	return arrAll, nil
}

func (model1 ModelUnder_init) SearchPartTimeUnder18SalaryModels(Keyword string, page int ,limit int) (arrSearch []initialize.PartTimeUnder18Salary, err error, CheckData int) {
	var Search initialize.PartTimeUnder18Salary
	db := db.Connect()
	querylimit := ``
	if strconv.Itoa(page) == "" && strconv.Itoa(limit) == ""{
		querylimit = ``
	}else {
		pageacheck := strconv.Itoa((page-1)*limit)
		showadata := strconv.Itoa(limit)
		querylimit = ` LIMIT `+pageacheck+`,`+showadata
	}
	db.QueryRow("SELECT count(*) FROM part_time_under_18_salary WHERE CONCAT_WS('', id_code_store, salary) LIKE ?", "%" + Keyword + "%").Scan(&CheckData)
	queryT := `SELECT id_part_time_under_18_salary, id_code_store, salary FROM part_time_under_18_salary WHERE CONCAT_WS('', id_code_store, salary) LIKE ?` +querylimit

	rows, err := db.Query(queryT, "%" + Keyword + "%")

	if err != nil {
		log.Print(err)
	}

	defer db.Close()
	for rows.Next() {
		if err := rows.Scan(&Search.Id_part_time_under_18_salary, &Search.Id_code_store, &Search.Salary); err != nil {
			log.Fatal(err.Error())
		} else {
			arrSearch = append(arrSearch, Search)
		}
	}

	return arrSearch, nil, CheckData
}

func (model1 ModelUnder_init) GetAllDataPartTimeUnder(Id_part_time_under_18_salary string) (arrGet []initialize.PartTimeUnder18Salary, err error) {
	var get initialize.PartTimeUnder18Salary
	db := db.Connect()

	result, err := db.Query("SELECT id_part_time_under_18_salary, id_code_store, salary FROM part_time_under_18_salary WHERE id_part_time_under_18_salary = ?", Id_part_time_under_18_salary)
	if err != nil {
		log.Println(err.Error())
	}
	defer result.Close()
	for result.Next() {

		err := result.Scan(&get.Id_part_time_under_18_salary, &get.Id_code_store, &get.Salary)
		if err != nil {
			log.Println(err.Error())
		} else {
			arrGet = append(arrGet, get)
		}
	}

	return arrGet, nil
}

func (model1 ModelUnder_init) InsertDataPartTimeUnder(insert *initialize.PartTimeUnder18Salary) (arrInsert []initialize.PartTimeUnder18Salary, err error) {
	db := db.Connect()
	stmt, err := db.Prepare("INSERT INTO part_time_under_18_salary (id_code_store, salary) VALUES(?,?)")
	if err != nil {
		log.Println(err.Error())
	}
	defer db.Close()

	stmt.Exec(insert.Id_code_store, insert.Salary)
	if err != nil {
		log.Println(err.Error())
	}

	Excute := initialize.PartTimeUnder18Salary{
		Id_code_store: insert.Id_code_store,
		Salary:        insert.Salary,
	}

	arrInsert = append(arrInsert, Excute)

	return arrInsert, nil
}

func (model1 ModelUnder_init) UpdateDataPartTimeUnder(update *initialize.PartTimeUnder18Salary) (arrUpdate []initialize.PartTimeUnder18Salary, err error) {
	db := db.Connect()

	stmt, err := db.Prepare("UPDATE part_time_under_18_salary SET id_code_store = ?, salary = ? WHERE id_part_time_under_18_salary = ?")
	if err != nil {
		log.Println(err.Error())
	}

	stmt.Exec(update.Id_code_store, update.Salary, update.Id_part_time_under_18_salary)
	if err != nil {
		log.Println(err.Error())
	}

	Excute := initialize.PartTimeUnder18Salary{
		Id_part_time_under_18_salary: update.Id_part_time_under_18_salary,
		Id_code_store:                update.Id_code_store,
		Salary:                       update.Salary,
	}

	arrUpdate = append(arrUpdate, Excute)

	return arrUpdate, nil

}

func (model1 ModelUnder_init) DeleteDataPartTimeUnder(delete *initialize.PartTimeUnder18Salary) (arrDelete []initialize.PartTimeUnder18Salary, err error) {
	db := db.Connect()
	stmt, err := db.Prepare("DELETE FROM part_time_under_18_salary WHERE id_part_time_under_18_salary = ?")
	if err != nil {
		log.Println(err.Error())
	}

	stmt.Exec(delete.Id_part_time_under_18_salary)
	if err != nil {
		log.Println(err.Error())
	}

	Excute := initialize.PartTimeUnder18Salary{
		Id_part_time_under_18_salary: delete.Id_part_time_under_18_salary,
	}

	arrDelete = append(arrDelete, Excute)

	return arrDelete, nil
}
