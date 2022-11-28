package countercontrolrepo

import "github.com/jmoiron/sqlx"

type (
	counterRepoDB struct {
		db *sqlx.DB
	}

	CounterControlAll struct {
		FISCAL_YEAR string `db:"FISCAL_YEAR"`
		SEMESTER    string `db:"SEMESTER"`
		YEAR        string `db:"YEAR"`
	}

	CounterControlAllRepoInterface interface {
		GetCounterControlAllData() (*[]CounterControlAll, error)
		//GetFacultyById(param1 string) (*[]CounterControlAll, error)
		//GetFacultyByJS(param1 string) (*[]CounterControlAll, error)
	}
)

func NewCounterControlAllRepo(db *sqlx.DB) CounterControlAllRepoInterface {
	return &counterRepoDB{db: db}
}
