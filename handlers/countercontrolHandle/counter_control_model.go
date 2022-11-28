package countercontrolhandle

import "eng-backend-api/services/countercontrolservice"

type (
	counterControlAllHandler struct {
		countercontrolAll_Res countercontrolservice.CounterControlAllResponeInterface
	}
)

func NewCounterControlAllHandler(res countercontrolservice.CounterControlAllResponeInterface) counterControlAllHandler {
	return counterControlAllHandler{countercontrolAll_Res: res}
}
