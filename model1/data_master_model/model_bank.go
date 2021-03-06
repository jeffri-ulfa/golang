package model1

import (
	"log"
	"strconv"
	"../../db"
	"../../initialize"
	"../../models"
)

type ModelBank_init models.DB_init

func (model1 ModelBank_init) ReturnAllDatabank() (arrAll []initialize.Bank, err error) {
	var all initialize.Bank

	db := db.Connect()

	rows, err := db.Query("SELECT id_bank, bank_code, bank_name, branch_code, branch_name, special FROM bank")

	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	for rows.Next() {
		if err := rows.Scan(&all.Id_bank, &all.Bank_code, &all.Bank_name, &all.Branch_code, &all.Branch_name, &all.Special); err != nil {
			log.Println(err.Error())

		} else {
			arrAll = append(arrAll, all)
		}
	}

	return arrAll, nil
}

func (model1 ModelBank_init) SearchDataBankModels(Keyword string, page int ,limit int) (arrSearch []initialize.Bank, err error, CheckData int) {
	var Search initialize.Bank
	db := db.Connect()
	querylimit := ``
	if strconv.Itoa(page) == "" && strconv.Itoa(limit) == ""{
		querylimit = ``
	}else {
		pageacheck := strconv.Itoa((page-1)*limit)
		showadata := strconv.Itoa(limit)
		querylimit = ` LIMIT `+pageacheck+`,`+showadata
	}
	db.QueryRow("SELECT count(*) FROM bank WHERE CONCAT_WS('', bank_code, bank_name, branch_code, branch_name, special) LIKE ?", "%" + Keyword + "%").Scan(&CheckData)
	queryT := `SELECT id_bank, bank_code, bank_name, branch_code, branch_name, special FROM bank WHERE CONCAT_WS('',bank_code, bank_name, branch_code, branch_name, special) LIKE ?` +querylimit

	rows, err := db.Query(queryT, "%" + Keyword + "%")

	if err != nil {
		log.Print(err)
	}

	defer db.Close()
	for rows.Next() {
		if err := rows.Scan(&Search.Id_bank, &Search.Bank_code, &Search.Bank_name, &Search.Branch_code, &Search.Branch_name, &Search.Special); err != nil {
			log.Fatal(err.Error())
		} else {
			arrSearch = append(arrSearch, Search)
		}
	}

	return arrSearch, nil, CheckData
}


func (model1 ModelBank_init) GetDataBank(Id_bank string) (arrGet []initialize.Bank, err error) {
	var all initialize.Bank

	db := db.Connect()

	result, err := db.Query("SELECT id_bank, bank_code, bank_name, branch_code, branch_name, special FROM bank WHERE id_bank = ?", Id_bank)
	if err != nil {
		log.Println(err.Error())
	}
	defer result.Close()
	for result.Next() {

		err := result.Scan(&all.Id_bank, &all.Bank_code, &all.Bank_name, &all.Branch_code, &all.Branch_name, &all.Special)
		if err != nil {
			log.Println(err.Error())
		} else {
			arrGet = append(arrGet, all)
		}
	}

	return arrGet, nil
}

func (model1 ModelBank_init) InsertDataBank(insert *initialize.Bank) (arrInsert []initialize.Bank, err error) {
	db := db.Connect()
	var id_bank int
	Id := db.QueryRow("SELECT MAX(id_bank)+1 FROM bank LIMIT 1").Scan(&id_bank)
	if Id != nil {
		log.Println(err.Error())
	}
	log.Println(id_bank)
	stmt, err := db.Prepare("INSERT INTO bank (id_bank, bank_code, bank_name, branch_code,branch_name,special) VALUES (?,?,?,?,?,?)")
	if err != nil {
		log.Println(err.Error())
	}
	defer db.Close()

	result, err := stmt.Exec(id_bank, insert.Bank_code, insert.Bank_name, insert.Branch_code, insert.Branch_name, insert.Special)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(result)
	Execute := initialize.Bank{
		Id_bank:	 id_bank,
		Bank_code:   insert.Bank_code,
		Bank_name:   insert.Bank_name,
		Branch_code: insert.Branch_code,
		Branch_name: insert.Branch_name,
		Special:     insert.Special,
	}
	log.Println(Execute)
	arrInsert = append(arrInsert, Execute)

	return arrInsert, nil
}

func (model1 ModelBank_init) UdpateDatabank(update *initialize.Bank) (arrUpdate []initialize.Bank, err error) {

	db := db.Connect()

	stmt, err := db.Prepare("UPDATE bank SET bank_code = ?, bank_name = ?, branch_code = ?, branch_name = ? , special = ? WHERE id_bank = ?")
	if err != nil {
		log.Println(err.Error())
	}

	result, err := stmt.Exec(update.Bank_code, update.Bank_name, update.Branch_code, update.Branch_name, update.Special, update.Id_bank)
	log.Println(result)

	Execute := initialize.Bank{
		Id_bank:     update.Id_bank,
		Bank_code:   update.Bank_code,
		Bank_name:   update.Bank_name,
		Branch_code: update.Branch_code,
		Branch_name: update.Branch_name,
		Special:     update.Special,
	}

	arrUpdate = append(arrUpdate, Execute)

	return arrUpdate, nil
}

func (model1 ModelBank_init) DeleteDataBank(delete *initialize.Bank) (arrDelete []initialize.Bank, err error) {
	db := db.Connect()
	stmt, err := db.Prepare("DELETE FROM bank WHERE id_bank = ?")
	if err != nil {
		log.Println(err.Error())
	}

	stmt.Exec(delete.Id_bank)
	if err != nil {
		log.Println(err.Error())
	}

	Execute := initialize.Bank{
		Id_bank: delete.Id_bank,
	}

	arrDelete = append(arrDelete, Execute)

	return arrDelete, nil
}
