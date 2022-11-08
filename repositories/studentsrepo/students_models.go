package studentsrepo

import "github.com/jmoiron/sqlx"

type (
	studentAllRepoDB struct {
		db *sqlx.DB
	}

	StudentAll struct {
		YEAR                 string `db:"YEAR"`
		SEMESTER             string `db:"SEMESTER"`
		UNIV_ID              string `db:"UNIV_ID"`
		CITIZEN_ID           string `db:"CITIZEN_ID"`
		STD_ID               string `db:"STD_ID"`
		PREFIX_NAME_ID       string `db:"PREFIX_NAME_ID"`
		STD_FNAME            string `db:"STD_FNAME"`
		STD_MNAME            string `db:"STD_MNAME"`
		STD_LNAME            string `db:"STD_LNAME"`
		STD_FNAME_EN         string `db:"STD_FNAME_EN"`
		STD_MNAME_EN         string `db:"STD_MNAME_EN"`
		STD_LNAME_EN         string `db:"STD_LNAME_EN"`
		GENDER_ID            string `db:"GENDER_ID"`
		BIRTHDAY             string `db:"BIRTHDAY"`
		SUB_DISTRICT_ID      string `db:"SUB_DISTRICT_ID"`
		NATIONALITY_ID       string `db:"NATIONALITY_ID"`
		ADMIT_YEAR           string `db:"ADMIT_YEAR"`
		FAC_ID               string `db:"FAC_ID"`
		CURR_ID              string `db:"CURR_ID"`
		STUDY_TYPE_ID        string `db:"STUDY_TYPE_ID"`
		STUDY_TIME_ID        string `db:"STUDY_TIME_ID"`
		STUDY_REG_ID         string `db:"STUDY_REG_ID"`
		CURR_REG_ID          string `db:"CURR_REG_ID"`
		CLASS                string `db:"CLASS"`
		GRAD_STATUS_ID       string `db:"GRAD_STATUS_ID"`
		STD_STATUS_ID        string `db:"STD_STATUS_ID"`
		TERMINATE_STUDY_CASE string `db:"TERMINATE_STUDY_CASE"`
		GPA                  string `db:"GPA"`
		GPAX                 string `db:"GPAX"`
		NUM_CREDIT           string `db:"NUM_CREDIT"`
		ACC_CREDIT           string `db:"ACC_CREDIT"`
		DEFORM_ID            string `db:"DEFORM_ID"`
		FUND_STATUS_ID       string `db:"FUND_STATUS_ID"`
		FUND_NAME            string `db:"FUND_NAME"`
		TALENT_ID            string `db:"TALENT_ID"`
		PASSPORT_NUMBER      string `db:"PASSPORT_NUMBER"`
		PASSPORT_STARTDATE   string `db:"PASSPORT_STARTDATE"`
		PASSPORT_ENDDATE     string `db:"PASSPORT_ENDDATE"`
		DEGREE_NUM           string `db:"DEGREE_NUM"`
	}
	StudentAllRepoInterface interface {
		GetStudentAllUnicon() (*[]StudentAll, error)
		GetStudentLimitUnicon(param1 string) (*[]StudentAll, error)
	}
)

func NewStudentAllRepo(db *sqlx.DB) StudentAllRepoInterface {
	return &studentAllRepoDB{db: db}
}
