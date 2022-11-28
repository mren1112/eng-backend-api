package universityservices

import "eng-backend-api/repositories/universityrepo"

type (
	universityAllService struct {
		universityAll_Repo universityrepo.UniversityAllRepoInterface
	}

	UniversityAllRespone struct {
		ACADEMIC_YEAR        string `json:"ACADEMIC_YEAR"`
		UNIV_ID              string `json:"UNIV_ID"`
		UNIV_NAME            string `json:"UNIV_NAME"`
		UNIV_NAME_EN         string `json:"UNIV_NAME_EN"`
		UNIV_CODE            string `json:"UNIV_CODE"`
		UNIV_CODE_EN         string `json:"UNIV_CODE_EN"`
		UNIV_TYPE            string `json:"UNIV_TYPE"`
		CAMPUS_TYPE          string `json:"CAMPUS_TYPE"`
		HOUSE_ID             string `json:"HOUSE_ID"`
		HOUSE_NUMBER         string `json:"HOUSE_NUMBER"`
		VILLAGE_NUMBER       string `json:"VILLAGE_NUMBER"`
		TROK                 string `json:"TROK"`
		SOI                  string `json:"SOI"`
		STREET               string `json:"STREET"`
		SUB_DISTRICT_ID      string `json:"SUB_DISTRICT_ID"`
		ZIPCODE              string `json:"ZIPCODE"`
		TELEPHONE_1          string `json:"TELEPHONE_1"`
		TELEPHONE_2          string `json:"TELEPHONE_2"`
		FAX                  string `json:"FAX"`
		EMAIL                string `json:"EMAIL"`
		WEBSITE              string `json:"WEBSITE"`
		ESTABLISHED_DATE     string `json:"ESTABLISHED_DATE"`
		ADMINISTRATOR_NAME   string `json:"ADMINISTRATOR_NAME"`
		LATITUDE             string `json:"LATITUDE"`
		LOGITUTE             string `json:"LOGITUTE"`
		EDUCATION_LEVEL_CODE string `json:"EDUCATION_LEVEL_CODE"`
		U_OPEN_1             string `json:"U_OPEN_1"`
		U_CLOSE_1            string `json:"U_CLOSE_1"`
		U_OPEN_2             string `json:"U_OPEN_2"`
		U_CLOSE_2            string `json:"U_CLOSE_2"`
		U_OPEN_3             string `json:"U_OPEN_3"`
		U_CLOSE_3            string `json:"U_CLOSE_3"`
	}
	UniversityAllResponeInterface interface {
		GetUniversityAllUnicon() (*[]UniversityAllRespone, error)
	}
)

func NewUniversityAllService(universityAllRepo universityrepo.UniversityAllRepoInterface) UniversityAllResponeInterface {
	return &universityAllService{universityAll_Repo: universityAllRepo}

}
