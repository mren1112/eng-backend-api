package facultyRepo

import (
	//"fmt"

)

func (u *facultyAllRepoDB) GetFacultyAllData() (*[]FacultyAll, error) {
	facultyAll := []FacultyAll{}
	query := `select a.FACULTY_NO,FACULTY_NAME_THAI,b.MAJOR_NO,FACULTY_NAME_ENG,FACULTY_NAME_SHORT, 
	 CURR_NO,MAJOR_NAME_THAI,b.MAJOR_NAME_ENG,ISCED_ID,PROGRAM_ID,ID_MAJOR 
	 from dbbach00.ugb_faculty a , dbbach00.ugb_major b 
	 where a.FACULTY_NO = b.FACULTY_NO and a.FACULTY_NO = '51'`

	//fmt.Println("GetFacultyAllData = ", query)

	err := u.db.Select(&facultyAll, query)
	if err != nil {
		return nil, err
	}

	return &facultyAll, nil
}

func (u *facultyAllRepoDB) GetFacultyById(param1 string) (*[]FacultyAll, error) {
	facultyAll := []FacultyAll{}
	query := `select a.FACULTY_NO,FACULTY_NAME_THAI,b.MAJOR_NO,FACULTY_NAME_ENG,FACULTY_NAME_SHORT, 
	CURR_NO,MAJOR_NAME_THAI,b.MAJOR_NAME_ENG,ISCED_ID,PROGRAM_ID,ID_MAJOR 
	from dbbach00.ugb_faculty a , dbbach00.ugb_major b 
	where a.FACULTY_NO = b.FACULTY_NO and a.FACULTY_NO = :param1`

	//fmt.Println("GetFacultyById = ", query)

	err := u.db.Select(&facultyAll, query, param1)
	if err != nil {
		return nil, err
	}

	return &facultyAll, nil
}

func (u *facultyAllRepoDB) GetFacultyByJS(param1 []string) (*[]FacultyAll, error) {
	facultyAll := []FacultyAll{}
	query := `select a.FACULTY_NO,FACULTY_NAME_THAI,b.MAJOR_NO,FACULTY_NAME_ENG,FACULTY_NAME_SHORT, 
	CURR_NO,MAJOR_NAME_THAI,b.MAJOR_NAME_ENG,ISCED_ID,PROGRAM_ID,ID_MAJOR 
	from dbbach00.ugb_faculty a , dbbach00.ugb_major b 
	where a.FACULTY_NO = b.FACULTY_NO and a.FACULTY_NO = :1 and b.MAJOR_NO = :2`

	//fmt.Println("GetFacultyByJS = ", query)

	err := u.db.Select(&facultyAll, query, param1[0],param1[1])
	if err != nil {
		return nil, err
	}

	return &facultyAll, nil
}
