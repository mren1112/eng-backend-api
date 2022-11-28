package facultyServices

import (
	//"eng-backend-api/repositories/facultyRepo"
	"fmt"
)

func (us *facultyAllService) GetFacultyAllData() (*[]FacultyAllRespone, error) {
	responeFacultyAll := []FacultyAllRespone{}
	facultyAll, err := us.facultyAll_Repo.GetFacultyAllData()

	if err != nil {
		return nil, err
	}
	for _, u := range *facultyAll {
		responeFacultyAll = append(responeFacultyAll, FacultyAllRespone{
			FACULTY_NO:         u.FACULTY_NO,
			FACULTY_NAME_THAI:  u.FACULTY_NAME_THAI,
			FACULTY_NAME_ENG:   u.FACULTY_NAME_ENG,
			FACULTY_NAME_SHORT: u.FACULTY_NAME_SHORT,
			MAJOR_NO:           u.MAJOR_NO,
			CURR_NO:            u.CURR_NO,
			MAJOR_NAME_THAI:    u.MAJOR_NAME_THAI,
			MAJOR_NAME_ENG:     u.MAJOR_NAME_ENG,
			ISCED_ID:           u.ISCED_ID,
			PROGRAM_ID:         u.PROGRAM_ID,
			ID_MAJOR:           u.ID_MAJOR,
		})
	}

	return &responeFacultyAll, nil
}

func (us *facultyAllService) GetFacultyById(param1 string) (*[]FacultyAllRespone, error) {
	responeFacultyAll := []FacultyAllRespone{}
	studentAll, err := us.facultyAll_Repo.GetFacultyById(param1)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for _, u := range *studentAll {
		responeFacultyAll = append(responeFacultyAll, FacultyAllRespone{
			FACULTY_NO:         u.FACULTY_NO,
			FACULTY_NAME_THAI:  u.FACULTY_NAME_THAI,
			FACULTY_NAME_ENG:   u.FACULTY_NAME_ENG,
			FACULTY_NAME_SHORT: u.FACULTY_NAME_SHORT,
			MAJOR_NO:           u.MAJOR_NO,
			CURR_NO:            u.CURR_NO,
			MAJOR_NAME_THAI:    u.MAJOR_NAME_THAI,
			MAJOR_NAME_ENG:     u.MAJOR_NAME_ENG,
			ISCED_ID:           u.ISCED_ID,
			PROGRAM_ID:         u.PROGRAM_ID,
			ID_MAJOR:           u.ID_MAJOR,
		})
	}

	return &responeFacultyAll, nil
}

func (us *facultyAllService) GetFacultyByJS(param1 []string) (*[]FacultyAllRespone, error) {
	responeFacultyAll := []FacultyAllRespone{}
	studentAll, err := us.facultyAll_Repo.GetFacultyByJS(param1)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for _, u := range *studentAll {
		responeFacultyAll = append(responeFacultyAll, FacultyAllRespone{
			FACULTY_NO:         u.FACULTY_NO,
			FACULTY_NAME_THAI:  u.FACULTY_NAME_THAI,
			FACULTY_NAME_ENG:   u.FACULTY_NAME_ENG,
			FACULTY_NAME_SHORT: u.FACULTY_NAME_SHORT,
			MAJOR_NO:           u.MAJOR_NO,
			CURR_NO:            u.CURR_NO,
			MAJOR_NAME_THAI:    u.MAJOR_NAME_THAI,
			MAJOR_NAME_ENG:     u.MAJOR_NAME_ENG,
			ISCED_ID:           u.ISCED_ID,
			PROGRAM_ID:         u.PROGRAM_ID,
			ID_MAJOR:           u.ID_MAJOR,
		})
	}

	return &responeFacultyAll, nil
}
