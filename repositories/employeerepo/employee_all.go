package employeerepo

func (u *employeeAllRepoDB) GetEmployeeAllUnicon() (*[]EmployeeAll, error) {
	employeeAll := []EmployeeAll{}
	query := `select  '2564' as YEAR, '2' as SEMESTER , UNIV_ID, CITIZEN_ID, PREFIX_NAME_ID, 
			STF_FNAME, STF_MNAME, STF_LNAME, STF_FNAME_EN,
			STF_MNAME_EN, STF_LNAME_EN, GENDER_ID,
			BIRTHDAY,
			HOMEADD, MOO, STREET, SUB_DISTRICT_ID, TELEPHONE, EMAIL, ZIPCODE,
			NATIONALITY_ID, STAFFTYPE_ID, TIME_CONTACT_ID, BUDGET_ID,
			 SUBSTAFFTYPE_ID, ADMIN_POSITION_ID, ACADEMICSTANDING_ID, 
			POSITIONLEVEL_ID, POSITION_WORK, DEPARTMENT_ID, 
			 DATE_INWORK, 
			 DATE_START_THIS_U ,
			'' as SPECIAL_NAME_ID_1,
			'' as SPECIAL_NAME_ID_2,
			TEACH_LEVEL_ID, 
			DEFORM_ID, INCOME_ID,
			 MOVEMENT_TYPE_ID, 
			 MOVEMENT_DATE, 
			 DECORATION_ID, DECORATION_DATE,
			'' as PASSPORT_STARTDATE,
			'' as PASSPORT_ENDDATE, 
			APPOINTMENT_ORDER,
			 APPOINTMENT_ORDER_START_DATE,
			 APPOINTMENT_DISMISSAL, APPOINTMENT_DISMISSAL, 
			SCHOLAR_ORDER_ID, 
			'' as SCHOLAR_ORDER_DATE, ACADEMICSTANDING_SUBJECT_ID, ACADEMICSTANDING_SUBJECT_OTHER, RESEARCHER_STATUS_ID 
			from person01.ds2001_test`

	err := u.db.Select(&employeeAll, query)
	if err != nil {
		return nil, err
	}

	return &employeeAll, nil
}

func (u *employeeAllRepoDB) GetEmployeeAllUniconMock() (*[]EmployeeAll, error) {
	employeeAll := []EmployeeAll{
		{
			YEAR:                           "2564",
			SEMESTER:                       "2",
			UNIV_ID:                        "00700",
			CITIZEN_ID:                     "1111111111111",
			PREFIX_NAME_ID:                 "",
			STF_FNAME:                      "",
			STF_MNAME:                      "",
			STF_LNAME:                      "",
			STF_FNAME_EN:                   "",
			STF_MNAME_EN:                   "",
			STF_LNAME_EN:                   "",
			GENDER_ID:                      "",
			BIRTHDAY:                       "",
			HOMEADD:                        "",
			MOO:                            "",
			STREET:                         "",
			SUB_DISTRICT_ID:                "",
			TELEPHONE:                      "",
			EMAIL:                          "",
			ZIPCODE:                        "",
			NATIONALITY_ID:                 "",
			STAFFTYPE_ID:                   "",
			TIME_CONTACT_ID:                "",
			BUDGET_ID:                      "",
			SUBSTAFFTYPE_ID:                "",
			ADMIN_POSITION_ID:              "",
			ACADEMICSTANDING_ID:            "",
			POSITIONLEVEL_ID:               "",
			POSITION_WORK:                  "",
			DEPARTMENT_ID:                  "",
			DATE_INWORK:                    "",
			DATE_START_THIS_U:              "",
			SPECIAL_NAME_ID_1:              "",
			SPECIAL_NAME_ID_2:              "",
			TEACH_LEVEL_ID:                 "",
			DEFORM_ID:                      "",
			INCOME_ID:                      "",
			MOVEMENT_TYPE_ID:               "",
			MOVEMENT_DATE:                  "",
			DECORATION_ID:                  "",
			DECORATION_DATE:                "",
			PASSPORT_STARTDATE:             "",
			PASSPORT_ENDDATE:               "",
			APPOINTMENT_ORDER:              "",
			APPOINTMENT_ORDER_START_DATE:   "",
			APPOINTMENT_DISMISSAL:          "",
			APPOINTMENT_DISMISSAL_DATE:     "",
			SCHOLAR_ORDER_ID:               "",
			SCHOLAR_ORDER_DATE:             "",
			ACADEMICSTANDING_SUBJECT_ID:    "",
			ACADEMICSTANDING_SUBJECT_OTHER: "",
			RESEARCHER_STATUS_ID:           "",
		},
	}

	return &employeeAll, nil
}
