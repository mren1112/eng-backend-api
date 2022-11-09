package facultiesServices

import "eng-backend-api/repositorries/facultiesrepo"

type (
	facultyAllService struct {
		facultyAll_Repo facultiesrepo.FacultyAllRepoInterface
	}

	FacultyAllRespone struct {
		YEAR                           string `json:"YEAR"`
		SEMESTER                       string `json:"SEMESTER"` 
	}
	FacultyAllResponeInterface interface {
		GetEmployeeAllUnicon() (*[]FacultyAllRespone, error)
	}
)