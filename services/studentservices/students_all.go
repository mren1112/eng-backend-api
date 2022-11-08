package studentservices

func (us *studentAllService) GetStudentAllUnicon() (*[]StudentAllRespone, error) {
	responeStudentAll := []StudentAllRespone{}
	studentAll, err := us.studentAll_Repo.GetStudentAllUnicon()

	if err != nil {
		return nil, err
	}
	for _, u := range *studentAll {
		responeStudentAll = append(responeStudentAll, StudentAllRespone{
			YEAR:                 u.YEAR,
			SEMESTER:             u.SEMESTER,
			UNIV_ID:              u.UNIV_ID,
			CITIZEN_ID:           u.CITIZEN_ID,
			STD_ID:               u.STD_ID,
			PREFIX_NAME_ID:       u.PREFIX_NAME_ID,
			STD_FNAME:            u.STD_FNAME,
			STD_MNAME:            u.STD_MNAME,
			STD_LNAME:            u.STD_LNAME,
			STD_FNAME_EN:         u.STD_FNAME_EN,
			STD_MNAME_EN:         u.STD_MNAME_EN,
			STD_LNAME_EN:         u.STD_LNAME_EN,
			GENDER_ID:            u.GENDER_ID,
			BIRTHDAY:             u.BIRTHDAY,
			SUB_DISTRICT_ID:      u.SUB_DISTRICT_ID,
			NATIONALITY_ID:       u.NATIONALITY_ID,
			ADMIT_YEAR:           u.ADMIT_YEAR,
			FAC_ID:               u.FAC_ID,
			CURR_ID:              u.CURR_ID,
			STUDY_TYPE_ID:        u.STUDY_TYPE_ID,
			STUDY_TIME_ID:        u.STUDY_TIME_ID,
			CURR_REG_ID:          u.CURR_REG_ID,
			CLASS:                u.CLASS,
			GRAD_STATUS_ID:       u.GRAD_STATUS_ID,
			STD_STATUS_ID:        u.STD_STATUS_ID,
			TERMINATE_STUDY_CASE: u.TERMINATE_STUDY_CASE,
			GPA:                  u.GPA,
			GPAX:                 u.GPAX,
			NUM_CREDIT:           u.NUM_CREDIT,
			ACC_CREDIT:           u.ACC_CREDIT,
			DEFORM_ID:            u.DEFORM_ID,
			FUND_STATUS_ID:       u.FUND_STATUS_ID,
			FUND_NAME:            u.FUND_NAME,
			TALENT_ID:            u.TALENT_ID,
			PASSPORT_NUMBER:      u.PASSPORT_NUMBER,
			PASSPORT_STARTDATE:   u.PASSPORT_STARTDATE,
			PASSPORT_ENDDATE:     u.PASSPORT_ENDDATE,
			DEGREE_NUM:           u.DEGREE_NUM,
		})
	}

	return &responeStudentAll, nil
}

func (us *studentAllService) GetStudentLimitUnicon(param1 string) (*[]StudentAllRespone, error) {
	responeStudentAll := []StudentAllRespone{}
	studentAll, err := us.studentAll_Repo.GetStudentLimitUnicon(param1)

	if err != nil {
		return nil, err
	}
	for _, u := range *studentAll {
		responeStudentAll = append(responeStudentAll, StudentAllRespone{
			YEAR:                 u.YEAR,
			SEMESTER:             u.SEMESTER,
			UNIV_ID:              u.UNIV_ID,
			CITIZEN_ID:           u.CITIZEN_ID,
			STD_ID:               u.STD_ID,
			PREFIX_NAME_ID:       u.PREFIX_NAME_ID,
			STD_FNAME:            u.STD_FNAME,
			STD_MNAME:            u.STD_MNAME,
			STD_LNAME:            u.STD_LNAME,
			STD_FNAME_EN:         u.STD_FNAME_EN,
			STD_MNAME_EN:         u.STD_MNAME_EN,
			STD_LNAME_EN:         u.STD_LNAME_EN,
			GENDER_ID:            u.GENDER_ID,
			BIRTHDAY:             u.BIRTHDAY,
			SUB_DISTRICT_ID:      u.SUB_DISTRICT_ID,
			NATIONALITY_ID:       u.NATIONALITY_ID,
			ADMIT_YEAR:           u.ADMIT_YEAR,
			FAC_ID:               u.FAC_ID,
			CURR_ID:              u.CURR_ID,
			STUDY_TYPE_ID:        u.STUDY_TYPE_ID,
			STUDY_TIME_ID:        u.STUDY_TIME_ID,
			CURR_REG_ID:          u.CURR_REG_ID,
			CLASS:                u.CLASS,
			GRAD_STATUS_ID:       u.GRAD_STATUS_ID,
			STD_STATUS_ID:        u.STD_STATUS_ID,
			TERMINATE_STUDY_CASE: u.TERMINATE_STUDY_CASE,
			GPA:                  u.GPA,
			GPAX:                 u.GPAX,
			NUM_CREDIT:           u.NUM_CREDIT,
			ACC_CREDIT:           u.ACC_CREDIT,
			DEFORM_ID:            u.DEFORM_ID,
			FUND_STATUS_ID:       u.FUND_STATUS_ID,
			FUND_NAME:            u.FUND_NAME,
			TALENT_ID:            u.TALENT_ID,
			PASSPORT_NUMBER:      u.PASSPORT_NUMBER,
			PASSPORT_STARTDATE:   u.PASSPORT_STARTDATE,
			PASSPORT_ENDDATE:     u.PASSPORT_ENDDATE,
			DEGREE_NUM:           u.DEGREE_NUM,
		})
	}

	return &responeStudentAll, nil
}
