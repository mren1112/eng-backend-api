package employeerepo

import "github.com/jmoiron/sqlx"

type (
	employeeAllRepoDB struct {
		db *sqlx.DB
	}

	EmployeeAll struct {
		YEAR                           string `db:"YEAR"`
		SEMESTER                       string `db:"SEMESTER"`
		UNIV_ID                        string `db:"UNIV_ID"`
		CITIZEN_ID                     string `db:"CITIZEN_ID"`
		PREFIX_NAME_ID                 string `db:"PREFIX_NAME_ID"`
		STF_FNAME                      string `db:"STF_FNAME"`
		STF_MNAME                      string `db:"STF_MNAME"`
		STF_LNAME                      string `db:"STF_LNAME"`
		STF_FNAME_EN                   string `db:"STF_FNAME_EN"`
		STF_MNAME_EN                   string `db:"STF_MNAME_EN"`
		STF_LNAME_EN                   string `db:"STF_LNAME_EN"`
		GENDER_ID                      string `db:"GENDER_ID"`
		BIRTHDAY                       string `db:"BIRTHDAY"`
		HOMEADD                        string `db:"HOMEADD"`
		MOO                            string `db:"MOO"`
		STREET                         string `db:"STREET"`
		SUB_DISTRICT_ID                string `db:"SUB_DISTRICT_ID"`
		TELEPHONE                      string `db:"TELEPHONE"`
		EMAIL                          string `db:"EMAIL"`
		ZIPCODE                        string `db:"ZIPCODE"`
		NATIONALITY_ID                 string `db:"NATIONALITY_ID"`
		STAFFTYPE_ID                   string `db:"STAFFTYPE_ID"`
		TIME_CONTACT_ID                string `db:"TIME_CONTACT_ID"`
		BUDGET_ID                      string `db:"BUDGET_ID"`
		SUBSTAFFTYPE_ID                string `db:"SUBSTAFFTYPE_ID"`
		ADMIN_POSITION_ID              string `db:"ADMIN_POSITION_ID"`
		ACADEMICSTANDING_ID            string `db:"ACADEMICSTANDING_ID"`
		POSITIONLEVEL_ID               string `db:"POSITIONLEVEL_ID"`
		POSITION_WORK                  string `db:"POSITION_WORK"`
		DEPARTMENT_ID                  string `db:"DEPARTMENT_ID"`
		DATE_INWORK                    string `db:"DATE_INWORK"`
		DATE_START_THIS_U              string `db:"DATE_START_THIS_U"`
		SPECIAL_NAME_ID_1              string `db:"SPECIAL_NAME_ID_1"`
		SPECIAL_NAME_ID_2              string `db:"SPECIAL_NAME_ID_2"`
		TEACH_LEVEL_ID                 string `db:"TEACH_LEVEL_ID"`
		DEFORM_ID                      string `db:"DEFORM_ID"`
		INCOME_ID                      string `db:"INCOME_ID"`
		MOVEMENT_TYPE_ID               string `db:"MOVEMENT_TYPE_ID"`
		MOVEMENT_DATE                  string `db:"MOVEMENT_DATE"`
		DECORATION_ID                  string `db:"DECORATION_ID"`
		DECORATION_DATE                string `db:"DECORATION_DATE"`
		PASSPORT_STARTDATE             string `db:"PASSPORT_STARTDATE"`
		PASSPORT_ENDDATE               string `db:"PASSPORT_ENDDATE"`
		APPOINTMENT_ORDER              string `db:"APPOINTMENT_ORDER"`
		APPOINTMENT_ORDER_START_DATE   string `db:"APPOINTMENT_ORDER_START_DATE"`
		APPOINTMENT_DISMISSAL          string `db:"APPOINTMENT_DISMISSAL"`
		APPOINTMENT_DISMISSAL_DATE     string `db:"APPOINTMENT_DISMISSAL_DATE"`
		SCHOLAR_ORDER_ID               string `db:"SCHOLAR_ORDER_ID"`
		SCHOLAR_ORDER_DATE             string `db:"SCHOLAR_ORDER_DATE"`
		ACADEMICSTANDING_SUBJECT_ID    string `db:"ACADEMICSTANDING_SUBJECT_ID"`
		ACADEMICSTANDING_SUBJECT_OTHER string `db:"ACADEMICSTANDING_SUBJECT_OTHER"`
		RESEARCHER_STATUS_ID           string `db:"RESEARCHER_STATUS_ID"`
	}
	EmployeeAllRepoInterface interface {
		GetEmployeeAllUnicon() (*[]EmployeeAll, error)
		GetEmployeeAllUniconMock() (*[]EmployeeAll, error)
	}
)

func NewEmployeeAllRepo(db *sqlx.DB) EmployeeAllRepoInterface {
	return &employeeAllRepoDB{db: db}
}
