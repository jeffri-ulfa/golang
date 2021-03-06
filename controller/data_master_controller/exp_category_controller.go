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

func ReturnAllExpCategory(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response

	db := db.Connect()
	_con := model1.ModelExp_init{DB: db}
	ExcuteData, err := _con.ReturnAllExp()
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

func SearchDataExpCategory(w http.ResponseWriter, r *http.Request) {
	var _response initialize.ResponseSearch
	db := db.Connect()
	type Name struct {
		Keyword  string `json:"keyword"`
		Page     int    `json:"page"`
		Show_data int    `json:"show_data"`
	}
	var Keyword Name
	json.NewDecoder(r.Body).Decode(&Keyword)
	_con := model1.ModelExp_init{DB: db}
	result, err, totalDataSearch := _con.SearchExpCategoryModels(Keyword.Keyword, Keyword.Page, Keyword.Show_data)
	if err != nil {
		log.Println(err.Error())
	}

	if r.Method == "POST" {
		if result == nil {
			var _response initialize.ResponseDataNull
			_response.Status = http.StatusBadRequest
			_response.Message = "Exp Category Data doesn't exists."
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


func ReturnAllExpCategoryPagination(w http.ResponseWriter, r *http.Request) {
	var exp initialize.ExpCategory
	var arrExpCategory []initialize.ExpCategory
	var _response initialize.Response

	db := db.Connect()
	defer db.Close()
	code := mux.Vars(r)

	totalDataPerPage, _ := strconv.Atoi(code["perPage"])
	page, _ := strconv.Atoi(code["page"])

	var totalData int
	err := db.QueryRow("SELECT COUNT(*) FROM exp_category").Scan(&totalData)

	totalPage := int(math.Ceil(float64(totalData) / float64(totalDataPerPage)))

	if page <= 0 {
		page = 1
	}

	firstIndex := (totalDataPerPage * page) - totalDataPerPage

	query := fmt.Sprintf("select id_exp, exp_category, created_date, created_time, code_category, content, rule_code from exp_category limit %d,%d", firstIndex, totalDataPerPage)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&exp.Id_exp, &exp.Exp_category, &exp.Created_date, &exp.Created_time, &exp.Code_category, &exp.Content, &exp.Rule_code); err != nil {

			log.Fatal(err.Error())

		} else {
			arrExpCategory = append(arrExpCategory, exp)
		}
	}
	if r.Method == "GET" {
		if arrExpCategory != nil {
			_response.Status = http.StatusOK
			_response.Message = "Success"
			_response.TotalPage = totalPage
			_response.CurrentPage = page
			_response.Data = arrExpCategory
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

func GetExpCategory(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response
	db := db.Connect()

	_id := r.URL.Query().Get("id_exp")

	_con := model1.ModelExp_init{DB: db}
	ExcuteData, err := _con.GetDataExp(_id)
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

func CreateExpCategory(w http.ResponseWriter, r *http.Request) {
	var init_insert initialize.ExpCategory
	var _response initialize.Response
	json.NewDecoder(r.Body).Decode(&init_insert)
	db := db.Connect()

	_con := model1.ModelExp_init{DB: db}
	ExcuteData, _ := _con.InsertDataExp(&init_insert)

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

func UpdateExpCategory(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response
	var init_insert initialize.ExpCategory
	json.NewDecoder(r.Body).Decode(&init_insert)
	db := db.Connect()

	_con := model1.ModelExp_init{DB: db}
	ExcuteData, _ := _con.UpdateDataExp(&init_insert)

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

func DeleteExpCategory(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response
	db := db.Connect()
	params := mux.Vars(r)
	delete := params["id_exp"]
	stmt, err := db.Exec("DELETE FROM exp_category WHERE id_exp = ?", delete)
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
