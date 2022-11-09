package facultiesrepo

import "github.com/jmoiron/sqlx"

type (
	facultyAllRepoDB struct {
		db *sqlx.DB
	}

	FacultyAll struct {
		YEAR     string `db:"YEAR"`
		SEMESTER string `db:"SEMESTER"`
	}
	FacultyAllRepoInterface interface {
		GetFacultyAllData() (*[]FacultyAll, error)
	}
)

func NewFacultyAllRepo(db *sqlx.DB) FacultyAllRepoInterface {
	return &facultyAllRepoDB{db: db}
}
