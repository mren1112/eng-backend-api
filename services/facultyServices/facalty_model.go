package facultyServices

import (
	"eng-backend-api/repositories/facultyRepo"
)

type (
	facultyAllService struct {
		facultyAll_Repo facultyRepo.FacultyAllRepoInterface
	}

	FacultyAllRespone struct {
		FACULTY_NO         string `json:"FACULTY_NO"`
		FACULTY_NAME_THAI  string `json:"FACULTY_NAME_THAI"`
		FACULTY_NAME_ENG   string `json:"FACULTY_NAME_ENG"`
		FACULTY_NAME_SHORT string `json:"FACULTY_NAME_SHORT"`
		MAJOR_NO           string `json:"MAJOR_NO"`
		CURR_NO            string `json:"CURR_NO"`
		MAJOR_NAME_THAI    string `json:"MAJOR_NAME_THAI"`
		MAJOR_NAME_ENG     string `json:"MAJOR_NAME_ENG"`
		ISCED_ID           string `json:"ISCED_ID"`
		PROGRAM_ID         string `json:"PROGRAM_ID"`
		ID_MAJOR           string `json:"ID_MAJOR"`
	}

	FacultyAllResponeInterface interface {
		GetFacultyAllData() (*[]FacultyAllRespone, error)
		GetFacultyById(param1 string) (*[]FacultyAllRespone, error)
		GetFacultyByJS(param1 []string) (*[]FacultyAllRespone, error)
	}
)

func NewFacultyAllService(facultyAllRepo facultyRepo.FacultyAllRepoInterface) FacultyAllResponeInterface {
	return &facultyAllService{facultyAll_Repo: facultyAllRepo}

}
