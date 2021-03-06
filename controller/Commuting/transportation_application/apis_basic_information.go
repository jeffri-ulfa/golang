package transportation_application

import (
	"../../../db"
	"../../../initialize"
	"../../../initialize/Commuting"
	models_enter_the_information "../../../models/Commuting/transportation_application"
	_Response "../../../response"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func ReturnCreateCommutingBasicInformation(w http.ResponseWriter, r *http.Request) {

	var init_insert Commuting.InsertBasicInformation
	var _response initialize.ResponseMaster
	json.NewDecoder(r.Body).Decode(&init_insert)
	param := mux.Vars(r)
	employee_number := param["employee_number"]
	db := db.Connect()
	if r.Method != "POST" {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Status Method Not Allowed"
		_response.Data = nil
		_Response.ResponseJson(w, _response.Status, _response)
	} else {
		_model := models_enter_the_information.Models_init_basic_information{DB: db}
		resultData, err := _model.Model_InsertBasicInformation(&init_insert,employee_number)
		defer db.Close()
		if err == "Success Response" {
			_response.Status = http.StatusOK
			_response.Message = err
			_response.Data = resultData
			_Response.ResponseJson(w, _response.Status, _response)
		} else {
			_response.Status = http.StatusBadRequest
			_response.Message = err
			_response.Data = nil
			_Response.ResponseJson(w, _response.Status, _response)
		}
	}
}

func ReturnGetByCommutingBasicInformation(w http.ResponseWriter, r *http.Request) {

	var _response initialize.ResponseMaster

	storeNumber := r.FormValue("store_number")
	employeeNumber := r.FormValue("employee_number")
	db := db.Connect()
	if r.Method != "POST" {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Status Method Not Allowed"
		_response.Data = nil
		_Response.ResponseJson(w, _response.Status, _response)
	} else {
		_model := models_enter_the_information.Models_init_basic_information{DB: db}
		ResultData, err := _model.Model_GetByIdCommutingBasicInformation(storeNumber, employeeNumber)
		defer db.Close()
		if err != nil {
			_response.Status = http.StatusBadRequest
			_response.Message = "Missing Body Request"
			_response.Data = nil
			_Response.ResponseJson(w, _response.Status, _response)
		} else {
			_response.Status = http.StatusOK
			_response.Message = "Success Response"
			_response.Data = ResultData
			_Response.ResponseJson(w, _response.Status, _response)
		}
	}

}
