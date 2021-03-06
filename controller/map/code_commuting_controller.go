package controller

import (
	initialize2 "Go_DX_Services/initialize/map"
	"encoding/json"
	"log"
	"net/http"
	"Go_DX_Services/db"
	"Go_DX_Services/initialize"
)

func ReturnAllCodeCommuting(w http.ResponseWriter, r *http.Request) {
	var codeCommunting initialize2.CodeCommuting
	var arrCodeCommuting []initialize2.CodeCommuting
	var response initialize.Response

	db := db.Connect()

	rows, err := db.Query("SELECT * FROM code_commuting")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&codeCommunting.Id_code, &codeCommunting.Code_random, &codeCommunting.Std_deviation, &codeCommunting.Created_time, &codeCommunting.Created_date, &codeCommunting.Status_commuting); err != nil {
			log.Fatal(err.Error())

		} else {
			arrCodeCommuting = append(arrCodeCommuting, codeCommunting)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrCodeCommuting

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
