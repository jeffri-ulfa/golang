package enter_the_information

import (
	"log"
	"../../initialize/enter_the_information"
	db2 "../../db"
	"strings"
	models3 "../../models"
)

func CheckDataById(sql string, id string) (CountData int) {
	db := db2.Connect()
	err := db.QueryRow(sql,id).Scan(&CountData)
	if err != nil {
		log.Println(err.Error())
	}
	return CountData
}

func CheckDataByIdNullString(sql string, id models3.NullString) (CountData int) {
	db := db2.Connect()
	err := db.QueryRow(sql,id).Scan(&CountData)
	if err != nil {
		log.Println(err.Error())
	}
	return CountData
}

func CheckDataByIdInt(sql string, id int) (CountData int) {
	db := db2.Connect()
	err := db.QueryRow(sql,id).Scan(&CountData)
	if err != nil {
		log.Println(err.Error())
	}
	return CountData
}

func CheckDataByStoreAndEmployee(sql string, store string,employee string) (CountData int) {
	db := db2.Connect()
	err := db.QueryRow(sql,store,employee).Scan(&CountData)
	if err != nil {
		log.Println(err.Error())
	}
	return CountData
}

func ValidatorInsertBasicInformation(Request *enter_the_information.InsertBasicInformation) (valid bool, message string) {

	if Request.CarInsuranceDocumentExpiryDate == "" {
		return false, "Missing required field in body request → CarInsuranceDocumentExpiryDate = <empty string>"
	}

	if Request.PropertyDamage == "" {
		return false, "Missing required field in body request → PropertyDamage = <empty string>"
	}

	if Request.PersonalInjury == "" {
		return false, "Missing required field in body request → PersonInjury = <empty string>"
	}

	if Request.DriverLicenseExpiryDate == "" {
		return false, "Missing required field in body request → DriverLicenseExpiryDate = <empty string>"
	}

	if Request.InsuranceCompany == "" {
		return false, "Missing required field in body request → InsureCompany = <empty string>"
	}

	if Request.IdGeneralInformation == "" {
		return false, "Missing required field in body request → IdGeneralInformation = <empty string>"
	}

	//Data := "select count(*) from commuting_basic_information where id_general_information = ?"
	//
	//CheckData := CheckDataById(Data, Request.IdGeneralInformation)
	//log.Println(CheckData)
	//if CheckData >= 1 {
	//	return false, "IdGeneralInformation Must 1"
	//}
	return true, "done"
}

func ValidatorInsertUsageRecordApplyForTravelExpenses(Request *enter_the_information.InsertUsageRecordApplyForTravelExpenses) (valid bool, message string) {

	if Request.RouteProfileName == "" {
		return false, "Missing required field in body request → RouteProfileName  = <empty string>"
	}
	if Request.Date == "" {
		return false, "Missing required field in body request → Date  = <empty string>"
	}
	if Request.Attendance == "" {
		return false, "Missing required field in body request → Attendance  = <empty string>"
	}
	if Request.IdGeneralInformation == "" {
		return false, "Missing required field in body request → IdGeneralInformation  = <empty string>"
	}
	return true, "done"
}

func ValidatorDetailInsertUsageRecordApplyForTravelExpenses(Request *enter_the_information.InsertDetailUsageRecordApplyForTravelExpenses) (valid bool, message string) {

	if Request.IdCommutingTrip == 0 {
		return false, "Missing required field in body request → TypeOfTransport = 0"
	}
	if Request.TypeOfTransport == "" {
		return false, "Missing required field in body request → TypeOfTransport = <empty string>"
	}
	if Request.Purpose == "" {
		return false, "Missing required field in body request → Purpose = <empty string>"
	}
	if Request.DetailFrom == "" {
		return false, "Missing required field in body request → DetailFrom = <empty string>"
	}
	if Request.DetailTo == "" {
		return false, "Missing required field in body request → DetailTo = <empty string>"
	}
	if Request.TransitPoint == "" {
		return false, "Missing required field in body request → TransitPoint = <empty string>"
	}

	return true, "done"

}

func GetAdditionalUsageRecord(store_number string, employee_number string, id_commuting_trip int,Condition string)( DatatypeOfTransportation string, DataRoute string,DataPurpose string){

	var typeOfTransportation string
	var DetailTo string
	var DetailFrom string
	var Purpose string
	var TransitPoint string
	db := db2.Connect()

	if Condition == "usageRecord-CheckData" {

		// Get Data Transportation, detail from, detail to and purpose (horizontal)
			GetDataTypeOfTransportationAndRoute, errGetDataTypeOfTransportationAndRoute := db.Query(`select trans.name_transportation_japanese, 
 										b.detail_from, b.detail_to, b.purpose, b.transit_point
										from basic_information bi, commuting_trip ct, detail_commuting_trip b, store_information si , general_information gi, 
										master_transportation trans
										where ct.id_commuting_trip = b.id_commuting_trip and gi.id_basic_information = bi.id_basic_information
										and b.type_of_transport =  trans.code_transportation
										and gi.id_store_code = si.id_code_store and ct.id_general_information = gi.id_general_information
										and si.code_store =? and bi.employee_code=?
										and ct.submit ='N' and ct.save_trip ='N' and b.id_commuting_trip = ? order by b.id_detail_commuting_trip
										`, store_number, employee_number, id_commuting_trip)

		if errGetDataTypeOfTransportationAndRoute != nil {
			return "", "", ""
		} else {
			for GetDataTypeOfTransportationAndRoute.Next() {
				errGetDataT := GetDataTypeOfTransportationAndRoute.Scan(&typeOfTransportation, &DetailFrom, &DetailTo, &Purpose, &TransitPoint)

				if errGetDataT != nil {
					log.Println(errGetDataT.Error())
				}
				DatatypeOfTransportation += typeOfTransportation + ` - `
				DataRoute += DetailFrom + ` - ` + strings.ReplaceAll(TransitPoint, ";", " - ") + ` - ` + DetailTo
				DataPurpose += Purpose + ` - `

			}
			if typeOfTransportation != "" {
				DatatypeOfTransportation = DatatypeOfTransportation[0 : len(DatatypeOfTransportation)-3]
			}
			//if DataRoute != "" {
			//	DataRoute = DataRoute[0 : len(DataRoute)-3]
			//}
			if DataPurpose != "" {
				DataPurpose = DataPurpose[0 : len(DataPurpose)-3]
			}
			//log.Println(DatatypeOfTransportation)
			//log.Println(DataRoute)
			//log.Println(DataPurpose)
		}
	}

	if Condition == "usageRecordUseRoute"{
		// Get Data Transportation, detail from, detail to and purpose (horizontal)
		GetDataTypeOfTransportationAndRoute, errGetDataTypeOfTransportationAndRoute := db.Query(`select trans.name_transportation_japanese, 
 										b.detail_from, b.detail_to, b.purpose
										from basic_information bi, commuting_trip ct, detail_commuting_trip b, store_information si , general_information gi, 
										master_transportation trans
										where ct.id_commuting_trip = b.id_commuting_trip and gi.id_basic_information = bi.id_basic_information
										and b.type_of_transport =  trans.code_transportation
										and gi.id_store_code = si.id_code_store and ct.id_general_information = gi.id_general_information
										and si.code_store =? and bi.employee_code=?
										and ct.submit ='N' and ct.save_trip ='Y' and b.id_commuting_trip = ?
										`, store_number, employee_number, id_commuting_trip)

		if errGetDataTypeOfTransportationAndRoute != nil {
			log.Println(0)
			typeOfTransportation = ""
			DetailTo = ""
			DetailFrom = ""
			Purpose = ""
		} else {
			for GetDataTypeOfTransportationAndRoute.Next() {
				errGetDataT := GetDataTypeOfTransportationAndRoute.Scan(&typeOfTransportation, &DetailFrom, &DetailTo, &Purpose)

				if errGetDataT != nil {
					log.Println(errGetDataT.Error())
				}
				DatatypeOfTransportation += typeOfTransportation + ` - `
				DataRoute += DetailFrom + ` - ` + DetailTo + ` - `
				DataPurpose += Purpose + ` - `

			}
			if typeOfTransportation != "" {
				DatatypeOfTransportation = DatatypeOfTransportation[0 : len(DatatypeOfTransportation)-3]
			}
			if DataRoute != "" {
				DataRoute = DataRoute[0 : len(DataRoute)-3]
			}
			if DataPurpose != "" {
				DataPurpose = DataPurpose[0 : len(DataPurpose)-3]
			}
			//log.Println(DatatypeOfTransportation)
			//log.Println(DataRoute)
			//log.Println(DataPurpose)
		}
		// end Get Data Transportation, detail from, detail to and purpose (horizontal)

	}

	if Condition == "usageRecordHistory"{
		// Get Data Transportation, detail from, detail to and purpose (horizontal)
		GetDataTypeOfTransportationAndRoute, errGetDataTypeOfTransportationAndRoute := db.Query(`select trans.name_transportation_japanese, 
 										b.detail_from, b.detail_to
										from basic_information bi, commuting_trip ct, detail_commuting_trip b, store_information si , general_information gi, 
										master_transportation trans
										where ct.id_commuting_trip = b.id_commuting_trip and gi.id_basic_information = bi.id_basic_information
										and b.type_of_transport =  trans.code_transportation
										and gi.id_store_code = si.id_code_store and ct.id_general_information = gi.id_general_information
										and si.code_store =? and bi.employee_code=?
										and ct.submit ='Y' and ct.save_trip ='N' and b.id_commuting_trip = ?
										`, store_number, employee_number, id_commuting_trip)

		if errGetDataTypeOfTransportationAndRoute != nil {
			log.Println(0)
			typeOfTransportation = ""
			DetailTo = ""
			DetailFrom = ""
			Purpose = ""
		} else {
			for GetDataTypeOfTransportationAndRoute.Next() {
				errGetDataT := GetDataTypeOfTransportationAndRoute.Scan(&typeOfTransportation, &DetailFrom, &DetailTo)

				if errGetDataT != nil {
					log.Println(errGetDataT.Error())
				}
				DatatypeOfTransportation += typeOfTransportation + ` - `
				DataRoute += DetailFrom + ` - ` + DetailTo + ` - `
				DataPurpose += Purpose + ` - `

			}
			if typeOfTransportation != "" {
				DatatypeOfTransportation = DatatypeOfTransportation[0 : len(DatatypeOfTransportation)-3]
			}
			if DataRoute != "" {
				DataRoute = DataRoute[0 : len(DataRoute)-3]
			}
			if DataPurpose != "" {
				DataPurpose = DataPurpose[0 : len(DataPurpose)-3]
			}
			//log.Println(DatatypeOfTransportation)
			//log.Println(DataRoute)
			//log.Println(DataPurpose)
		}
		// end Get Data Transportation, detail from, detail to and purpose (horizontal)

	}

	return DatatypeOfTransportation, DataRoute, DataPurpose
	// end Get Data Transportation, detail from, detail to and purpose (horizontal)
}