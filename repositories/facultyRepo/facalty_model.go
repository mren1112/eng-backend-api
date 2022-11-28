package facultyRepo

import (
	"github.com/jmoiron/sqlx"
	
)

type (
	facultyAllRepoDB struct {
		db *sqlx.DB
	}

	FacultyAll struct {
		FACULTY_NO         string `db:"FACULTY_NO"`
		FACULTY_NAME_THAI  string `db:"FACULTY_NAME_THAI"`
		FACULTY_NAME_ENG   string `db:"FACULTY_NAME_ENG"`
		FACULTY_NAME_SHORT string `db:"FACULTY_NAME_SHORT"`
		MAJOR_NO           string `db:"MAJOR_NO"`
		CURR_NO            string `db:"CURR_NO"`
		MAJOR_NAME_THAI    string `db:"MAJOR_NAME_THAI"`
		MAJOR_NAME_ENG     string `db:"MAJOR_NAME_ENG"`
		ISCED_ID           string `db:"ISCED_ID"`
		PROGRAM_ID         string `db:"PROGRAM_ID"`
		ID_MAJOR           string `db:"ID_MAJOR"`
	}

	FacultyJsonRespone struct {
		FACULTY_NO         string `json:"FACULTY_NO"` 
		MAJOR_NO           string `json:"MAJOR_NO"` 
	}

	FacultyAllRepoInterface interface {
		GetFacultyAllData() (*[]FacultyAll, error)
		GetFacultyById(param1 string) (*[]FacultyAll, error)
		GetFacultyByJS(param1 []string) (*[]FacultyAll, error)
	}
)

func NewFacultyAllRepo(db *sqlx.DB) FacultyAllRepoInterface {
	return &facultyAllRepoDB{db: db}
}
