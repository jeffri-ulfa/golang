package data_master_controller

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"../../db"
	"../../initialize"
	model1 "../../model1/data_master_model"
	"../../response"
	"github.com/gorilla/mux"
)

func ReturnAllPartTimeUnder18Salary(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response

	db := db.Connect()
	_con := model1.ModelUnder_init{DB: db}
	ExcuteData, err := _con.ReturnAllDataUnder18()
	if err != nil {
		log.Println(err.Error())
	}

	if r.Method == "GET" {
		if ExcuteData == nil {
			_response.Status = http.StatusBadRequest
			_response.Message = "Sorry Your Input Missing Body Bad Request"
			_response.Data = "Null"
			response.ResponseJson(w, _response.Status, _response)
		} else {
			_response.Status = http.StatusOK
			_response.Message = "Success"
			_response.Data = ExcuteData
			response.ResponseJson(w, _response.Status, _response)
		}
	} else {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Sorry Your Method Missing Not Allowed"
		_response.Data = "Null"
		response.ResponseJson(w, _response.Status, _response)
	}

}

func SeacrhDataPartTimeUnder18Salary(w http.ResponseWriter, r *http.Request) {
	var _response initialize.ResponseSearch
	db := db.Connect()
	type Name struct {
		Keyword  string `json:"keyword"`
		Page     int    `json:"page"`
		Show_data int    `json:"show_data"`
	}
	var Keyword Name
	json.NewDecoder(r.Body).Decode(&Keyword)
	_con := model1.ModelUnder_init{DB: db}
	result, err, totalDataSearch := _con.SearchPartTimeUnder18SalaryModels(Keyword.Keyword, Keyword.Page, Keyword.Show_data)
	if err != nil {
		log.Println(err.Error())
	}

	if r.Method == "POST" {
		if result == nil {
			var _response initialize.ResponseDataNull
			_response.Status = http.StatusBadRequest
			_response.Message = "Part Time Under 18 Salary Data doesn't exists."
			response.ResponseJson(w, _response.Status, _response)
		} else {
			_response.Status = http.StatusOK
			_response.Message = "Success"
			_response.TotalData = totalDataSearch
			_response.TotalPage = (totalDataSearch / Keyword.Show_data) + 1
			_response.Data = result
			response.ResponseJson(w, _response.Status, _response)
		}
	} else {
		var _response initialize.ResponseDataNull
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Sorry Your Method Missing Not Allowed"
		response.ResponseJson(w, _response.Status, _response)
	}
}


func ReturnAllPartTimeUnder18SalaryPagination(w http.ResponseWriter, r *http.Request) {
	var partTimeUnderSalary initialize.PartTimeUnder18Salary
	var arrPartTimeUnder18Salary []initialize.PartTimeUnder18Salary
	var _response initialize.Response

	db := db.Connect()
	defer db.Close()
	code := mux.Vars(r)

	totalDataPerPage, _ := strconv.Atoi(code["perPage"])
	page, _ := strconv.Atoi(code["page"])

	var totalData int
	err := db.QueryRow("SELECT COUNT(*) FROM part_time_under_18_salary").Scan(&totalData)

	totalPage := int(math.Ceil(float64(totalData) / float64(totalDataPerPage)))
	if page <= 0 {
		page = 1
	}

	firstIndex := (totalDataPerPage * page) - totalDataPerPage

	query := fmt.Sprintf("select id_part_time_under_18_salary,id_code_store,salary from part_time_under_18_salary limit %d,%d", firstIndex, totalDataPerPage)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&partTimeUnderSalary.Id_part_time_under_18_salary, &partTimeUnderSalary.Id_code_store, &partTimeUnderSalary.Salary); err != nil {

			log.Fatal(err.Error())

		} else {
			arrPartTimeUnder18Salary = append(arrPartTimeUnder18Salary, partTimeUnderSalary)
		}
	}
	if r.Method == "GET" {
		if arrPartTimeUnder18Salary != nil {
			_response.Status = http.StatusOK
			_response.Message = "Success"
			_response.TotalPage = totalPage
			_response.CurrentPage = page
			_response.Data = arrPartTimeUnder18Salary
			response.ResponseJson(w, _response.Status, _response)
		} else if page > totalPage {
			_response.Status = http.StatusBadRequest
			_response.Message = "Sorry Your Input Missing Body Bad Request"
			_response.TotalPage = totalPage
			_response.CurrentPage = page
			_response.Data = "Null"
			response.ResponseJson(w, _response.Status, _response)
		}
	} else {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Sorry Your Method Missing Not Allowed"
		_response.TotalPage = totalPage
		_response.CurrentPage = page
		_response.Data = "Null"
		response.ResponseJson(w, _response.Status, _response)
	}
}

func GetPartTimeUnder18Salary(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response
	db := db.Connect()

	_id := r.URL.Query().Get("id_part_time_under_18_salary")

	_con := model1.ModelUnder_init{DB: db}
	ExcuteData, err := _con.GetAllDataPartTimeUnder(_id)
	if err != nil {
		log.Println(err.Error())
	}
	if r.Method == "GET" {
		if ExcuteData == nil {
			_response.Status = http.StatusBadRequest
			_response.Message = "Sorry Your Input Missing Body Bad Request"
			_response.Data = "Null"
			response.ResponseJson(w, _response.Status, _response)
		} else {
			_response.Status = http.StatusOK
			_response.Message = "Success"
			_response.Data = ExcuteData
			response.ResponseJson(w, _response.Status, _response)
		}
	} else {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Sorry Your Method Missing Not Allowed"
		_response.Data = "Null"
		response.ResponseJson(w, _response.Status, _response)
	}
}

func CreatePartTimeUnder18Salary(w http.ResponseWriter, r *http.Request) {
	var init_insert initialize.PartTimeUnder18Salary
	var _response initialize.Response
	json.NewDecoder(r.Body).Decode(&init_insert)
	db := db.Connect()

	_con := model1.ModelUnder_init{DB: db}
	ExcuteData, _ := _con.InsertDataPartTimeUnder(&init_insert)

	if r.Method == "POST" {
		if ExcuteData == nil {
			_response.Status = http.StatusBadRequest
			_response.Message = "Sorry Your Input Missing Body Bad Request"
			_response.Data = "Null"
			response.ResponseJson(w, _response.Status, _response)
		} else {
			_response.Status = http.StatusOK
			_response.Message = "Success"
			_response.Data = init_insert
			response.ResponseJson(w, _response.Status, _response)
		}
	} else {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Sorry Your Method Missing Not Allowed"
		_response.Data = "Null"
		response.ResponseJson(w, _response.Status, _response)
	}

}

func UpdatePartTimeUnder18Salary(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response
	var init_insert initialize.PartTimeUnder18Salary
	json.NewDecoder(r.Body).Decode(&init_insert)
	db := db.Connect()

	_con := model1.ModelUnder_init{DB: db}
	ExcuteData, _ := _con.UpdateDataPartTimeUnder(&init_insert)

	if r.Method == "PUT" {
		if ExcuteData == nil {
			_response.Status = http.StatusBadRequest
			_response.Message = "Sorry Your Input Missing Body Bad Request"
			_response.Data = "Null"
			response.ResponseJson(w, _response.Status, _response)
		} else {
			_response.Status = http.StatusOK
			_response.Message = "Success"
			_response.Data = ExcuteData
			response.ResponseJson(w, _response.Status, _response)
		}
	} else {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Sorry Your Method Missing Not Allowed"
		_response.Data = "Null"
		response.ResponseJson(w, _response.Status, _response)
	}
}

func DeletePartTimeUnder18Salary(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response
	db := db.Connect()
	params := mux.Vars(r)
	delete := params["id_part_time_under_18_salary"]
	stmt, err := db.Exec("DELETE FROM part_time_under_18_salary WHERE id_part_time_under_18_salary = ?", delete)
	if err != nil {
		log.Println(err.Error())
	}

	statment, err := stmt.RowsAffected()

	if r.Method == "DELETE" {
		if statment != 1 {
			_response.Status = http.StatusBadRequest
			_response.Message = "Sorry Your Input Missing Body Bad Request"
			_response.Data = nil
			response.ResponseJson(w, _response.Status, _response)
		} else {
			_response.Status = http.StatusOK
			_response.Message = "Success Data has been Deleted with ID"
			_response.Data = delete
			response.ResponseJson(w, _response.Status, _response)
		}
	} else {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Sorry Your Method Missing Not Allowed"
		_response.Data = nil
		response.ResponseJson(w, _response.Status, _response)
	}
}
