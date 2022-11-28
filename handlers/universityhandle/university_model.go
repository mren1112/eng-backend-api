package universityhandle

import "eng-backend-api/services/universityservices"

type (
	universityAllHandler struct {
		universityAll_Res universityservices.UniversityAllResponeInterface
	}
)

func NewUniversityAllHandler(res universityservices.UniversityAllResponeInterface) universityAllHandler {
	return universityAllHandler{universityAll_Res: res}
}
