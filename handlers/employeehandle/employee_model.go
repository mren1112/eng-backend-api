package employeehandle

import "eng-backend-api/services/employeeservices"

type (
	employeeAllHandler struct {
		employeeAll_Res employeeservices.EmployeeAllResponeInterface
	}
)

func NewEmployeeAllHandler(res employeeservices.EmployeeAllResponeInterface) employeeAllHandler {
	return employeeAllHandler{employeeAll_Res: res}
}
