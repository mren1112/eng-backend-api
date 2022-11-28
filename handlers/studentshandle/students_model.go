package studentshandle

import "eng-backend-api/services/studentservices"

type (
	studentAllHandler struct {
		studentAll_Res studentservices.StudentAllResponeInterface
	}
)

func NewStudentAllHandler(res studentservices.StudentAllResponeInterface) studentAllHandler {
	return studentAllHandler{studentAll_Res: res}
}
