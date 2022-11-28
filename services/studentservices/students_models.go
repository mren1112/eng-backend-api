package studentservices

import "eng-backend-api/repositories/studentsrepo"

type (
	studentAllService struct {
		studentAll_Repo studentsrepo.StudentAllRepoInterface
	}

	StudentAllRespone struct {
		YEAR                 string `json:"YEAR"`
		SEMESTER             string `json:"SEMESTER"`
		UNIV_ID              string `json:"UNIV_ID"`
		CITIZEN_ID           string `json:"CITIZEN_ID"`
		STD_ID               string `json:"STD_ID"`
		PREFIX_NAME_ID       string `json:"PREFIX_NAME_ID"`
		STD_FNAME            string `json:"STD_FNAME"`
		STD_MNAME            string `json:"STD_MNAME"`
		STD_LNAME            string `json:"STD_LNAME"`
		STD_FNAME_EN         string `json:"STD_FNAME_EN"`
		STD_MNAME_EN         string `json:"STD_MNAME_EN"`
		STD_LNAME_EN         string `json:"STD_LNAME_EN"`
		GENDER_ID            string `json:"GENDER_ID"`
		BIRTHDAY             string `json:"BIRTHDAY"`
		SUB_DISTRICT_ID      string `json:"SUB_DISTRICT_ID"`
		NATIONALITY_ID       string `json:"NATIONALITY_ID"`
		ADMIT_YEAR           string `json:"ADMIT_YEAR"`
		FAC_ID               string `json:"FAC_ID"`
		CURR_ID              string `json:"CURR_ID"`
		STUDY_TYPE_ID        string `json:"STUDY_TYPE_ID"`
		STUDY_TIME_ID        string `json:"STUDY_TIME_ID"`
		STUDY_REG_ID         string `json:"STUDY_REG_ID"`
		CURR_REG_ID          string `json:"CURR_REG_ID"`
		CLASS                string `json:"CLASS"`
		GRAD_STATUS_ID       string `json:"GRAD_STATUS_ID"`
		STD_STATUS_ID        string `json:"STD_STATUS_ID"`
		TERMINATE_STUDY_CASE string `json:"TERMINATE_STUDY_CAU"`
		GPA                  string `json:"GPA"`
		GPAX                 string `json:"GPAX"`
		NUM_CREDIT           string `json:"NUM_CREDIT"`
		ACC_CREDIT           string `json:"ACC_CREDIT"`
		DEFORM_ID            string `json:"DEFORM_ID"`
		FUND_STATUS_ID       string `json:"FUND_STATUS_ID"`
		FUND_NAME            string `json:"FUND_NAME"`
		TALENT_ID            string `json:"TALENT_ID"`
		PASSPORT_NUMBER      string `json:"PASSPORT_NUMBER"`
		PASSPORT_STARTDATE   string `json:"PASSPORT_STARTDATE"`
		PASSPORT_ENDDATE     string `json:"PASSPORT_ENDDATE"`
		DEGREE_NUM           string `json:"DEGREE_NUM"`
	}
	StudentAllResponeInterface interface {
		GetStudentAllUnicon() (*[]StudentAllRespone, error)
		GetStudentLimitUnicon(param1 string) (*[]StudentAllRespone, error)
	}
)

func NewStudentAllService(studentAllRepo studentsrepo.StudentAllRepoInterface) StudentAllResponeInterface {
	return &studentAllService{studentAll_Repo: studentAllRepo}

}
