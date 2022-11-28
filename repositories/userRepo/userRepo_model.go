package userrepo

import "github.com/jmoiron/sqlx"

type (
	userRepoDB struct {
		db *sqlx.DB
	}

	UserAll struct {
		FACULTY_NO         string `db:"FACULTY_NO"`
		FACULTY_NAME_THAI  string `db:"FACULTY_NAME_THAI"` 
	}

	UserRepoInterface interface {
		//GetFacultyAllData() (*[]FacultyAll, error)
		//GetFacultyById(param1 string) (*[]FacultyAll, error) 
	}
)

func NewFacultyAllRepo(db *sqlx.DB) UserRepoInterface {
	return &userRepoDB{db: db}
}
