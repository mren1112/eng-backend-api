package universityrepo

import "github.com/jmoiron/sqlx"

type (
	universityAllRepoDB struct {
		db *sqlx.DB
	}

	UniversityAll struct {
		ACADEMIC_YEAR        string `db:"ACADEMIC_YEAR"`
		UNIV_ID              string `db:"UNIV_ID"`
		UNIV_NAME            string `db:"UNIV_NAME"`
		UNIV_NAME_EN         string `db:"UNIV_NAME_EN"`
		UNIV_CODE            string `db:"UNIV_CODE"`
		UNIV_CODE_EN         string `db:"UNIV_CODE_EN"`
		UNIV_TYPE            string `db:"UNIV_TYPE"`
		CAMPUS_TYPE          string `db:"CAMPUS_TYPE"`
		HOUSE_ID             string `db:"HOUSE_ID"`
		HOUSE_NUMBER         string `db:"HOUSE_NUMBER"`
		VILLAGE_NUMBER       string `db:"VILLAGE_NUMBER"`
		TROK                 string `db:"TROK"`
		SOI                  string `db:"SOI"`
		STREET               string `db:"STREET"`
		SUB_DISTRICT_ID      string `db:"SUB_DISTRICT_ID"`
		ZIPCODE              string `db:"ZIPCODE"`
		TELEPHONE_1          string `db:"TELEPHONE_1"`
		TELEPHONE_2          string `db:"TELEPHONE_2"`
		FAX                  string `db:"FAX"`
		EMAIL                string `db:"EMAIL"`
		WEBSITE              string `db:"WEBSITE"`
		ESTABLISHED_DATE     string `db:"ESTABLISHED_DATE"`
		ADMINISTRATOR_NAME   string `db:"ADMINISTRATOR_NAME"`
		LATITUDE             string `db:"LATITUDE"`
		LOGITUTE             string `db:"LOGITUTE"`
		EDUCATION_LEVEL_CODE string `db:"EDUCATION_LEVEL_CODE"`
		U_OPEN_1             string `db:"U_OPEN_1"`
		U_CLOSE_1            string `db:"U_CLOSE_1"`
		U_OPEN_2             string `db:"U_OPEN_2"`
		U_CLOSE_2            string `db:"U_CLOSE_2"`
		U_OPEN_3             string `db:"U_OPEN_3"`
		U_CLOSE_3            string `db:"U_CLOSE_3"`
	}
	UniversityAllRepoInterface interface {
		GetUniversityAllUnicon() (*[]UniversityAll, error)
	}
)

func NewUniversityAllRepo(db *sqlx.DB) UniversityAllRepoInterface {
	return &universityAllRepoDB{db: db}
}
