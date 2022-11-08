package employeeservices

import "eng-backend-api/repositories/employeerepo"

type (
	employeeAllService struct {
		employeeAll_Repo employeerepo.EmployeeAllRepoInterface
	}

	EmployeeAllRespone struct {
		YEAR                           string `json:"YEAR"`
		SEMESTER                       string `json:"SEMESTER"`
		UNIV_ID                        string `json:"UNIV_ID"`
		CITIZEN_ID                     string `json:"CITIZEN_ID"`
		PREFIX_NAME_ID                 string `json:"PREFIX_NAME_ID"`
		STF_FNAME                      string `json:"STF_FNAME"`
		STF_MNAME                      string `json:"STF_MNAME"`
		STF_LNAME                      string `json:"STF_LNAME"`
		STF_FNAME_EN                   string `json:"STF_FNAME_EN"`
		STF_MNAME_EN                   string `json:"STF_MNAME_EN"`
		STF_LNAME_EN                   string `json:"STF_LNAME_EN"`
		GENDER_ID                      string `json:"GENDER_ID"`
		BIRTHDAY                       string `json:"BIRTHDAY"`
		HOMEADD                        string `json:"HOMEADD"`
		MOO                            string `json:"MOO"`
		STREET                         string `json:"STREET"`
		SUB_DISTRICT_ID                string `json:"SUB_DISTRICT_ID"`
		TELEPHONE                      string `json:"TELEPHONE"`
		EMAIL                          string `json:"EMAIL"`
		ZIPCODE                        string `json:"ZIPCODE"`
		NATIONALITY_ID                 string `json:"NATIONALITY_ID"`
		STAFFTYPE_ID                   string `json:"STAFFTYPE_ID"`
		TIME_CONTACT_ID                string `json:"TIME_CONTACT_ID"`
		BUDGET_ID                      string `json:"BUDGET_ID"`
		SUBSTAFFTYPE_ID                string `json:"SUBSTAFFTYPE_ID"`
		ADMIN_POSITION_ID              string `json:"ADMIN_POSITION_ID"`
		ACADEMICSTANDING_ID            string `json:"ACADEMICSTANDING_ID"`
		POSITIONLEVEL_ID               string `json:"POSITIONLEVEL_ID"`
		POSITION_WORK                  string `json:"POSITION_WORK"`
		DEPARTMENT_ID                  string `json:"DEPARTMENT_ID"`
		DATE_INWORK                    string `json:"DATE_INWORK"`
		DATE_START_THIS_U              string `json:"DATE_START_THIS_U"`
		SPECIAL_NAME_ID_1              string `json:"SPECIAL_NAME_ID_1"`
		SPECIAL_NAME_ID_2              string `json:"SPECIAL_NAME_ID_2"`
		TEACH_LEVEL_ID                 string `json:"TEACH_LEVEL_ID"`
		DEFORM_ID                      string `json:"DEFORM_ID"`
		INCOME_ID                      string `json:"INCOME_ID"`
		MOVEMENT_TYPE_ID               string `json:"MOVEMENT_TYPE_ID"`
		MOVEMENT_DATE                  string `json:"MOVEMENT_DATE"`
		DECORATION_ID                  string `json:"DECORATION_ID"`
		DECORATION_DATE                string `json:"DECORATION_DATE"`
		PASSPORT_STARTDATE             string `json:"PASSPORT_STARTDATE"`
		PASSPORT_ENDDATE               string `json:"PASSPORT_ENDDATE"`
		APPOINTMENT_ORDER              string `json:"APPOINTMENT_ORDER"`
		APPOINTMENT_ORDER_START_DATE   string `json:"APPOINTMENT_ORDER_START_DATE"`
		APPOINTMENT_DISMISSAL          string `json:"APPOINTMENT_DISMISSAL"`
		APPOINTMENT_DISMISSAL_DATE     string `json:"APPOINTMENT_DISMISSAL_DATE"`
		SCHOLAR_ORDER_ID               string `json:"SCHOLAR_ORDER_ID"`
		SCHOLAR_ORDER_DATE             string `json:"SCHOLAR_ORDER_DATE"`
		ACADEMICSTANDING_SUBJECT_ID    string `json:"ACADEMICSTANDING_SUBJECT_ID"`
		ACADEMICSTANDING_SUBJECT_OTHER string `json:"ACADEMICSTANDING_SUBJECT_OTHER"`
		RESEARCHER_STATUS_ID           string `json:"RESEARCHER_STATUS_ID"`
	}
	EmployeeAllResponeInterface interface {
		GetEmployeeAllUnicon() (*[]EmployeeAllRespone, error)
	}
)

func NewEmployeeAllService(employeeAllRepo employeerepo.EmployeeAllRepoInterface) EmployeeAllResponeInterface {
	return &employeeAllService{employeeAll_Repo: employeeAllRepo}

}
