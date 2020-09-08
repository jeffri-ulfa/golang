package general_recrutment_model

import (
	"log"

	"../../db"
	initialize "../../initialize/general_recrutment"
	"../../models"
)

type ModelGeneral_init models.DB_init

func (model1 ModelGeneral_init) InsertDataGeneralRecrutment(data *initialize.GeneralRecrutmentJoin) (join []initialize.GeneralRecrutmentJoin, err error) {
	db := db.Connect()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	// {
	var allTest []initialize.BasicInformationGeneral
	var test initialize.BasicInformationGeneral
	stmt, err := tx.Prepare(`INSERT INTO basic_information (employee_code, first_name, last_name, gender, birthdate, add_postal_code, id_prefecture, Adress, Adress_kana, Adress_detail, Adress_detail_kana, add_phone_number, marital_status, dormitory_status)
	VALUES (?,?,?,?,DATE_FORMAT(CONVERT_TZ(NOW(), @@session.time_zone, '+09:00'),'%Y-%m-%d'),?,?,?,?,?,?,?,?,?)
	SELECT basic_information.employee_code, basic_information.id_basic_information FROM basic_information
	INNER JOIN general_information AS employee_code ON basic_information.id_basic_information = general_information.id_basic_information`)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	stmt.Exec(test.Employee_code, test.First_name, test.Last_name, test.Gender, test.Add_postal_code, test.Id_prefecture,
		test.Adress, test.Adress_kana, test.Adress_detail, test.Adress_detail_kana,
		test.Add_phone_number, test.Marital_status, test.Dormitory_status)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	ExcuteData := initialize.BasicInformationGeneral{
		Employee_code:      test.Employee_code,
		First_name:         test.First_name,
		Last_name:          test.Last_name,
		Gender:             test.Gender,
		Add_postal_code:    test.Add_postal_code,
		Id_prefecture:      test.Id_prefecture,
		Adress:             test.Adress,
		Adress_kana:        test.Adress_kana,
		Adress_detail:      test.Adress_detail,
		Adress_detail_kana: test.Adress_detail_kana,
		Add_phone_number:   test.Add_phone_number,
		Marital_status:     test.Marital_status,
		Dormitory_status:   test.Dormitory_status,
	}
	allTest = append(allTest, ExcuteData)

	// 	return join, nil
	// }
	// {
	var arrTest1 []initialize.GeneralInformationGeneral
	var test1 initialize.GeneralInformationGeneral
	stmt1, err := tx.Prepare(`SELECT general_information.basic_information FROM general_information
		INSERT INTO general_information (id_basic_information, id_store_code, id_department, id_code_store, id_unit, join_date, id_bank, account_type, account_number, account_name, first_smester_in_other_company,
		distance_trip, resume_document, written_oath_document, employee_agreement_document, certificate_of_residence_card_document, application_form_of_commuting_method_document, compliance_agreement_document, 
		with_holding_slip_document, dependent_deduction_form_document, pension_book_document, health_check_report_document, office_code) 
		VALUES (?,?,?,?,?,DATE_FORMAT(CONVERT_TZ(NOW(), @@session.time_zone, '+09:00'),'%Y-%m-%d'),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)
		INNER JOIN employment_type ON general_information.id_basic_information = employment_type.id_basic_information
		SELECT employment_type.employment_status FROM employment_type INSERT INTO IF(employment_status = 34) VALUES (?)  INNER JOIN full_time_employee ON employment_type.id_employment_type = full_time_employee.id_employment_type
		SELECT employment_type.employment_status FROM employment_type INSERT INTO IF(employment_status = 41 OR 43) VALUES (?) INNER JOIN part_time_above_18_employee ON employment_type.id_employment_type = part_time_above_18_employee.id_employment_type
		SELECT employment_type.employment_status FROM employment_type INSERT INTO IF(employment_status = 42) VALUES (?)  INNER JOIN part_time_under_18_employee ON employment_type.id_employment_type = part_time_under_18_employee.id_employment_type`)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	stmt1.Exec(test1.Id_general_information, test1.Id_basic_information, test1.Id_store_code, test1.Id_department, test1.Id_store_section, test1.Id_unit,
		test1.Id_bank, test1.Account_type, test1.Account_number, test1.Account_name, test1.First_smester_in_other_company,
		test1.Distance_trip, test1.Resume_document, test1.Written_oath_document, test1.Employee_agreement_document,
		test1.Certificate_of_residence_card_document, test1.Application_form_of_commuting_method_document, test1.Compliance_agreement_document,
		test1.With_holding_slip_document, test1.Dependent_deduction_form_document, test1.Pension_book_document, test1.Health_check_report_document,
		test1.Office_code)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	ExcuteData1 := initialize.GeneralInformationGeneral{
		Id_general_information:                 test1.Id_general_information,
		Id_basic_information:                   test1.Id_basic_information,
		Id_store_code:                          test1.Id_store_code,
		Id_department:                          test1.Id_department,
		Id_store_section:                       test1.Id_store_section,
		Id_unit:                                test1.Id_unit,
		Id_bank:                                test1.Id_bank,
		Account_type:                           test1.Account_type,
		Account_number:                         test1.Account_number,
		First_smester_in_other_company:         test1.First_smester_in_other_company,
		Distance_trip:                          test1.Distance_trip,
		Resume_document:                        test1.Resume_document,
		Written_oath_document:                  test1.Written_oath_document,
		Employee_agreement_document:            test1.Employee_agreement_document,
		Certificate_of_residence_card_document: test1.Certificate_of_residence_card_document,
		Application_form_of_commuting_method_document: test1.Application_form_of_commuting_method_document,
		Compliance_agreement_document:                 test1.Compliance_agreement_document,
		With_holding_slip_document:                    test1.With_holding_slip_document,
		Dependent_deduction_form_document:             test1.Dependent_deduction_form_document,
		Pension_book_document:                         test1.Pension_book_document,
		Health_check_report_document:                  test1.Health_check_report_document,
		Office_code:                                   test1.Office_code,
	}
	arrTest1 = append(arrTest1, ExcuteData1)

	// 	return join, nil
	// }
	// {
	var test2 initialize.GeneralRecrutmentStatusApproval
	var arrTest2 []initialize.GeneralRecrutmentStatusApproval
	stmt2, err := tx.Prepare(`SELECT full_time_employee.id_employment_type,part_time_above_18_emlpoyee.id_employment_type, 
		part_time_under_18_employee.id_employment_type FROM general_recrutment_status_approval
		INSERT INTO general_recrutment_status_approval (id_basic_information, status, reason, date_time, date_time_approve, data_check, date_time_date_check, 
		message_for_edit, by_employee_code, flag, real_data) VALUES (?,?,?,NOW(),NOW(),?,NOW(),?,?,?,?,?)
		INNER JOIN full_time_employee ON general_recrutment_status_approval.id_employment_type = full_time_employee.id_employment_type
		INNER JOIN part_time_above_18_emlpoyee ON general_recrutment_status_approval.id_employment_type = part_time_above_18_emlpoyee.id_employment_type
		INNER JOIN part_time_under_18_employee ON general_recrutment_status_approval.id_employment_type = part_time_under_18_employee.id_employment_type)`)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	stmt2.Exec(test2.Id_basic_information, test2.Status, test2.Reason, test2.Data_check,
		test2.Message_for_edit, test2.By_employee_code, test2.Flag, test2.Real_data)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	ExcuteData2 := initialize.GeneralRecrutmentStatusApproval{
		Id_basic_information: test2.Id_basic_information,
		Status:               test2.Status,
		Reason:               test2.Reason,
		Data_check:           test2.Data_check,
		Message_for_edit:     test2.Message_for_edit,
		By_employee_code:     test2.By_employee_code,
		Flag:                 test2.Flag,
		Real_data:            test2.Real_data,
	}

	arrTest2 = append(arrTest2, ExcuteData2)

	commitTx := tx.Commit()

	if commitTx != nil {
		log.Fatal(commitTx)
	}
	return join, nil
	// }
	// var all initialize.GeneralRecrutmentStatusApproval
	// ctx := context.Background()

	// db := db.Connect()
	// tx, err := db.BeginTx(ctx, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _, err = tx.ExecContext(ctx, "INSERT INTO users (username, password) VALUES ('jeffri', 'asdf'), ('admin', '12345')")
	// if err != nil {
	// 	tx.Rollback()
	// 	return
	// }

	// // Run a query to get a count of all cats
	// row := tx.QueryRow("SELECT count(*) FROM users WHERE username='jeffri'")
	// var catCount int
	// // Store the count in the `catCount` variable
	// err = row.Scan(&catCount)
	// if err != nil {
	// 	tx.Rollback()
	// 	return
	// }

	// // Now update the food table, increasing the quantity of cat food by 10x the number of cats
	// _, err = tx.ExecContext(ctx, "UPDATE last_user SET last_name=last_name+$1 WHERE last_password='admin'", 10*catCount)
	// if err != nil {
	// 	tx.Rollback()
	// 	return
	// }

	// // Commit the change if all queries ran successfully
	// err = tx.Commit()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// {
	// result, err := tx.Exec("INSERT INTO basic_information (employee_code, first_name, last_name, gender, birthdate, add_postal_code, id_prefecture, addres, addres_kana, addres_detail, addres_detail_kana, add_phone_number, marital_status, dormitory_status) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	// if err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// 	log.Println(err)
	// }

	// if _, err := result.Exec(all.Employee_code, first_name, last_name, gender, birthdate, add_postal_code, id_prefecture, addres, addres_kana, addres_detail, addres_detail_kana, add_postal_code, marital_status, dormitory_status); err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// 	log.Println(err)
	// }
	// }
	// {
	// result1, err := tx.Exec("INSERT INTO general_information (id_basic_information, id_store_code, id_department, id_code_store, id_unit, join_date, id_bank, account_type, account_number, account_name, first_smester_in_other_company, distance_trip, resume_document, written_oath_document, employee_agreement_document, certificate_of_residence_card_document, application_form_of_commuting_method_document, compliance_agreement_document, with_holding_slip_document, dependent_deduction_form_document, pension_book_document, health_check_report_document, office_code) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	// if err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// 	log.Println(err)
	// }

	// if _, err := result.Exce(general_informaeExecon, id_store_code, id_department, id_code_store, id_unit, join_date, id_bank, account_type, account_number, account_name, first_smester_in_other_company, distance_trip, resume_document, written_oath_document, employee_agreement_document, certificate_of_residence_card_document, application_form_of_commuting_method_document, compliance_agreement_document, with_holding_slip_document, dependent_deduction_form_document, pension_book_document, health_check_report_document, office_code); err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// 	log.Println(err)
	// }

	// }
	// {
	// result2, err := tx.Exec("INSERT INTO general_recruitment_status_approval (id_basic_information, status, reason, date_time, date_time_approve, date_check, date_time_date_check, message_for_edit, by_employee_code, flag, real_data) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)")
	// if err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// 	log.Println(err)
	// }

	// if _, err := result.Prepare(id_basic_information, status, reason, date_time, date_time_approve, date_check, date_time_date_check, message_for_edit, by_employee_code, flag, real_data); err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// 	log.Println(err)
	// }
	// }

}