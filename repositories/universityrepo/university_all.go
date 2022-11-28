package universityrepo

func (u *universityAllRepoDB) GetUniversityAllUnicon() (*[]UniversityAll, error) {
	universityAll := []UniversityAll{}
	query := `select  ACADEMIC_YEAR,
	UNIV_ID,
	UNIV_NAME,
	UNIV_NAME_EN,
	UNIV_CODE,
	UNIV_CODE_EN,
	UNIV_TYPE,
	CAMPUS_TYPE,
	HOUSE_ID,
	HOUSE_NUMBER,
	VILLAGE_NUMBER,
	TROK,
	SOI,
	STREET,
	SUB_DISTRICT_ID,
	ZIPCODE,
	TELEPHONE_1,
	TELEPHONE_2,
	FAX,
	EMAIL,
	WEBSITE,
	ESTABLISHED_DATE,
	ADMINISTRATOR_NAME,
	LATITUDE,
	LOGITUTE,
	EDUCATION_LEVEL_CODE,
	U_OPEN_1,
	U_CLOSE_1,
	U_OPEN_2,
	U_CLOSE_2,
	U_OPEN_3,
	U_CLOSE_3 from scenter01.unic_ds4001`

	err := u.db.Select(&universityAll, query)
	if err != nil {
		return nil, err
	}

	return &universityAll, nil
}
