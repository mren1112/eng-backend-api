package facultyHandle

import "eng-backend-api/services/facultyServices"

type (
	facultyAllHandler struct {
		facultyAll_Res facultyServices.FacultyAllResponeInterface
	}
)

func NewFacultyAllHandler(res facultyServices.FacultyAllResponeInterface) facultyAllHandler {
	return facultyAllHandler{facultyAll_Res: res}
}
