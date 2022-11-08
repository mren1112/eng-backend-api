package studentsrepo

func (u *studentAllRepoDB) GetStudentAllUnicon() (*[]StudentAll, error) {
	studentAll := []StudentAll{}
	query := `select YEAR, SEMESTER, UNIV_ID, CITIZEN_ID, STD_ID, PREFIX_NAME_ID, STD_FNAME, STD_MNAME, STD_LNAME, STD_FNAME_EN,
	STD_MNAME_EN, STD_LNAME_EN, GENDER_ID,  BIRTHDAY, SUB_DISTRICT_ID, NATIONALITY_ID, ADMIT_YEAR, FAC_ID,
	CURR_ID, STUDY_TYPE_ID, STUDY_TIME_ID, CURR_REG_ID, CLASS, GRAD_STATUS_ID, STD_STATUS_ID, TERMINATE_STUDY_CASE, GPA, GPAX,
	NUM_CREDIT, ACC_CREDIT, DEFORM_ID, FUND_STATUS_ID, FUND_NAME, TALENT_ID, PASSPORT_NUMBER, DEGREE_NUM
	from dbbach00.vm_ds1001`

	err := u.db.Select(&studentAll, query)
	if err != nil {
		return nil, err
	}

	return &studentAll, nil
}

func (u *studentAllRepoDB) GetStudentLimitUnicon(param1 string) (*[]StudentAll, error) {
	studentAll := []StudentAll{}
	query := `select YEAR, SEMESTER, UNIV_ID, CITIZEN_ID, STD_ID, PREFIX_NAME_ID, STD_FNAME, STD_MNAME, STD_LNAME, STD_FNAME_EN,
	STD_MNAME_EN, STD_LNAME_EN, GENDER_ID, BIRTHDAY, SUB_DISTRICT_ID, NATIONALITY_ID, ADMIT_YEAR, FAC_ID,
	CURR_ID, STUDY_TYPE_ID, STUDY_TIME_ID, CURR_REG_ID, CLASS, GRAD_STATUS_ID, STD_STATUS_ID, TERMINATE_STUDY_CASE, GPA, GPAX,
	NUM_CREDIT, ACC_CREDIT, DEFORM_ID, FUND_STATUS_ID, FUND_NAME, TALENT_ID, PASSPORT_NUMBER, DEGREE_NUM
	from dbbach00.vm_ds1001  where ROWNUM <= :param1`

	err := u.db.Select(&studentAll, query, param1)
	if err != nil {
		return nil, err
	}

	return &studentAll, nil
}
