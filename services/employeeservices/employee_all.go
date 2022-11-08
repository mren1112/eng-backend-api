package employeeservices

func (us *employeeAllService) GetEmployeeAllUnicon() (*[]EmployeeAllRespone, error) {
	responeEmployeeAll := []EmployeeAllRespone{}
	employeeAll, err := us.employeeAll_Repo.GetEmployeeAllUnicon()

	if err != nil {
		return nil, err
	}
	for _, u := range *employeeAll {
		responeEmployeeAll = append(responeEmployeeAll, EmployeeAllRespone{
			YEAR:                           u.YEAR,
			SEMESTER:                       u.SEMESTER,
			UNIV_ID:                        u.UNIV_ID,
			CITIZEN_ID:                     u.CITIZEN_ID,
			PREFIX_NAME_ID:                 u.PREFIX_NAME_ID,
			STF_FNAME:                      u.STF_FNAME,
			STF_MNAME:                      u.STF_MNAME,
			STF_LNAME:                      u.STF_LNAME,
			STF_FNAME_EN:                   u.STF_FNAME_EN,
			STF_MNAME_EN:                   u.STF_MNAME_EN,
			STF_LNAME_EN:                   u.STF_LNAME_EN,
			GENDER_ID:                      u.GENDER_ID,
			BIRTHDAY:                       u.BIRTHDAY,
			HOMEADD:                        u.HOMEADD,
			MOO:                            u.MOO,
			STREET:                         u.STREET,
			SUB_DISTRICT_ID:                u.SUB_DISTRICT_ID,
			TELEPHONE:                      u.TELEPHONE,
			EMAIL:                          u.EMAIL,
			ZIPCODE:                        u.ZIPCODE,
			NATIONALITY_ID:                 u.NATIONALITY_ID,
			STAFFTYPE_ID:                   u.STAFFTYPE_ID,
			TIME_CONTACT_ID:                u.TIME_CONTACT_ID,
			BUDGET_ID:                      u.BUDGET_ID,
			SUBSTAFFTYPE_ID:                u.SUBSTAFFTYPE_ID,
			ADMIN_POSITION_ID:              u.ADMIN_POSITION_ID,
			ACADEMICSTANDING_ID:            u.ACADEMICSTANDING_ID,
			POSITIONLEVEL_ID:               u.POSITIONLEVEL_ID,
			POSITION_WORK:                  u.POSITION_WORK,
			DEPARTMENT_ID:                  u.DEPARTMENT_ID,
			DATE_INWORK:                    u.DATE_INWORK,
			DATE_START_THIS_U:              u.DATE_START_THIS_U,
			SPECIAL_NAME_ID_1:              u.SPECIAL_NAME_ID_1,
			SPECIAL_NAME_ID_2:              u.SPECIAL_NAME_ID_2,
			TEACH_LEVEL_ID:                 u.TEACH_LEVEL_ID,
			DEFORM_ID:                      u.DEFORM_ID,
			INCOME_ID:                      u.INCOME_ID,
			MOVEMENT_TYPE_ID:               u.MOVEMENT_TYPE_ID,
			MOVEMENT_DATE:                  u.MOVEMENT_DATE,
			DECORATION_ID:                  u.DECORATION_ID,
			DECORATION_DATE:                u.DECORATION_DATE,
			PASSPORT_STARTDATE:             u.PASSPORT_STARTDATE,
			PASSPORT_ENDDATE:               u.PASSPORT_ENDDATE,
			APPOINTMENT_ORDER:              u.APPOINTMENT_ORDER,
			APPOINTMENT_ORDER_START_DATE:   u.APPOINTMENT_ORDER_START_DATE,
			APPOINTMENT_DISMISSAL:          u.APPOINTMENT_DISMISSAL,
			APPOINTMENT_DISMISSAL_DATE:     u.APPOINTMENT_DISMISSAL_DATE,
			SCHOLAR_ORDER_ID:               u.SCHOLAR_ORDER_ID,
			SCHOLAR_ORDER_DATE:             u.SCHOLAR_ORDER_DATE,
			ACADEMICSTANDING_SUBJECT_ID:    u.ACADEMICSTANDING_SUBJECT_ID,
			ACADEMICSTANDING_SUBJECT_OTHER: u.ACADEMICSTANDING_SUBJECT_OTHER,
			RESEARCHER_STATUS_ID:           u.RESEARCHER_STATUS_ID,
		})
	}

	return &responeEmployeeAll, nil
}
