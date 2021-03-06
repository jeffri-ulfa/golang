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

func ReturnAllStoreInformation(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response

	db := db.Connect()
	_con := model1.Models_init{DB: db}
	ExcuteData, err := _con.ReturnAllStoreInformationModel()
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

func SortStoreInformation(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response
	db := db.Connect()
	ASC := "asc"
	DESC := "desc"

	Sort := r.FormValue("sort")
	Col := r.FormValue("kolom")
                   
	_con := model1.Models_init{DB: db}
	if Sort == ASC {
		ExcuteData, err := _con.SortASCStoreInformationModel(Sort, Col); 
		if err != nil {
			log.Println(err.Error())
		}
		if r.Method == "POST" {
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
	} else if Sort == DESC {
		ExcuteData, err := _con.SortDESCStoreInformationModel(Sort, Col);
		if err != nil {
			log.Println(err.Error())
		}

		if r.Method == "POST" {
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
	} else {
		_response.Status = http.StatusBadRequest
		_response.Message = "Sorry Your Input Missing Body Bad Request"
		_response.Data = []string{}
		response.ResponseJson(w, _response.Status, _response)
	}
}

func SearchDataStoreInformation(w http.ResponseWriter, r *http.Request) {
	var _response initialize.ResponseSearch
	db := db.Connect()
	type Name struct {
		Keyword  string `json:"keyword"`
		Page     int    `json:"page"`
		Show_data int    `json:"show_data"`
	}
	var Keyword Name
	json.NewDecoder(r.Body).Decode(&Keyword)
	_con := model1.Models_init{DB: db}
	result, err, totalDataSearch := _con.SearchStoreInformationModels(Keyword.Keyword, Keyword.Page, Keyword.Show_data)
	if err != nil {
		log.Println(err.Error())
	}

	if r.Method == "POST" {
		if result == nil {
			var _response initialize.ResponseDataNull
			_response.Status = http.StatusBadRequest
			_response.Message = "Store Information Data doesn't exists."
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


func ReturnAllFilterInformation(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response
	_id := r.FormValue("id_code_store")
	db := db.Connect()
	_con := model1.Models_init{DB: db}
	ExcuteData, err := _con.ReturnFilterStoreInformationModel(_id)
	if err != nil {
		log.Println(err.Error())
	}

	if r.Method == "POST" {
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

func ReturnAllStoreInformationPagination(w http.ResponseWriter, r *http.Request) {
	var storeInformation initialize.StoreInformation
	var arrStoreInformation []initialize.StoreInformation
	var _response initialize.Response

	db := db.Connect()
	defer db.Close()
	code := mux.Vars(r)

	totalDataPerPage, _ := strconv.Atoi(code["perPage"])
	page, _ := strconv.Atoi(code["page"])

	var totalData int
	err := db.QueryRow("SELECT COUNT(*) FROM store_information").Scan(&totalData)

	totalPage := int(math.Ceil(float64(totalData) / float64(totalDataPerPage)))

	if page <= 0 {
		page = 1
	}

	firstIndex := (totalDataPerPage * page) - totalDataPerPage

	query := fmt.Sprintf("select id_code_store,code_store,store_name from store_information limit %d,%d", firstIndex, totalDataPerPage)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&storeInformation.Id_code_store, &storeInformation.Code_store, &storeInformation.Store_name); err != nil {
			log.Fatal(err.Error())
		} else {
			arrStoreInformation = append(arrStoreInformation, storeInformation)
		}
	}
	if r.Method == "GET" {
		if arrStoreInformation != nil {
			_response.Status = http.StatusOK
			_response.Message = "Success"
			_response.TotalPage = totalPage
			_response.CurrentPage = page
			_response.Data = arrStoreInformation
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

func GetStoreInformation(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response
	db := db.Connect()
	_id := r.URL.Query().Get("id_code_store")

	_con := model1.Models_init{DB: db}
	ExcuteData, err := _con.GetIdStoreInformation(_id)
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

func CreateStoreInformation(w http.ResponseWriter, r *http.Request) {
	var init_insert initialize.StoreInformation
	var _response initialize.Response
	json.NewDecoder(r.Body).Decode(&init_insert)
	db := db.Connect()

	_con := model1.Models_init{DB: db}
	ExcuteData, _ := _con.InsertStoreInformation(&init_insert)

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

func UpdateStoreInformation(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response
	var init_insert initialize.StoreInformation
	json.NewDecoder(r.Body).Decode(&init_insert)
	db := db.Connect()

	_con := model1.Models_init{DB: db}
	ExcuteData, _ := _con.UpdateStoreInformation(&init_insert)

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

func DeleteStoreInformation(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response
	db := db.Connect()
	params := mux.Vars(r)
	delete := params["id_code_store"]
	stmt, err := db.Exec("DELETE FROM store_information WHERE id_code_store = ?", delete)
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
