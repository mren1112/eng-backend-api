package universityservices

func (us *universityAllService) GetUniversityAllUnicon() (*[]UniversityAllRespone, error) {
	responeUniversityAll := []UniversityAllRespone{}
	universityAll, err := us.universityAll_Repo.GetUniversityAllUnicon()

	if err != nil {
		return nil, err
	}
	for _, u := range *universityAll {
		responeUniversityAll = append(responeUniversityAll, UniversityAllRespone{
			ACADEMIC_YEAR:        u.ACADEMIC_YEAR,
			UNIV_ID:              u.UNIV_ID,
			UNIV_NAME:            u.UNIV_NAME,
			UNIV_NAME_EN:         u.UNIV_NAME_EN,
			UNIV_CODE:            u.UNIV_CODE,
			UNIV_CODE_EN:         u.UNIV_CODE_EN,
			UNIV_TYPE:            u.UNIV_TYPE,
			CAMPUS_TYPE:          u.CAMPUS_TYPE,
			HOUSE_ID:             u.HOUSE_ID,
			HOUSE_NUMBER:         u.HOUSE_NUMBER,
			VILLAGE_NUMBER:       u.VILLAGE_NUMBER,
			TROK:                 u.TROK,
			SOI:                  u.SOI,
			STREET:               u.STREET,
			SUB_DISTRICT_ID:      u.SUB_DISTRICT_ID,
			ZIPCODE:              u.ZIPCODE,
			TELEPHONE_1:          u.TELEPHONE_1,
			TELEPHONE_2:          u.TELEPHONE_2,
			FAX:                  u.FAX,
			EMAIL:                u.EMAIL,
			WEBSITE:              u.WEBSITE,
			ESTABLISHED_DATE:     u.ESTABLISHED_DATE,
			ADMINISTRATOR_NAME:   u.ADMINISTRATOR_NAME,
			LATITUDE:             u.LATITUDE,
			LOGITUTE:             u.LOGITUTE,
			EDUCATION_LEVEL_CODE: u.EDUCATION_LEVEL_CODE,
			U_OPEN_1:             u.U_OPEN_1,
			U_CLOSE_1:            u.U_CLOSE_1,
			U_OPEN_2:             u.U_OPEN_2,
			U_CLOSE_2:            u.U_CLOSE_2,
			U_OPEN_3:             u.U_OPEN_3,
			U_CLOSE_3:            u.U_CLOSE_3,
		})
	}

	return &responeUniversityAll, nil
}
