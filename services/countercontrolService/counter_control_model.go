package countercontrolservice

import "eng-backend-api/repositories/countercontrolrepo"

type (
	counterControlAllService struct {
		counterCountrolAll_Repo countercontrolrepo.CounterControlAllRepoInterface
	}

	CounterControlAllRespone struct {
		FISCAL_YEAR string `json:"FISCAL_YEAR"`
		SEMESTER    string `json:"SEMESTER"`
		YEAR        string `json:"YEAR"`
	}

	CounterControlAllResponeInterface interface {
		GetCounterControlAllData() (*[]CounterControlAllRespone, error)
		//GetCounterControlById(param1 string) (*[]CounterControlAllRespone, error)
		//GetCounterControlByJS(param1 string) (*[]CounterControlAllRespone, error)
	}
)

func NewCounterControlAllService(countercontrolAllrepo countercontrolrepo.CounterControlAllRepoInterface) CounterControlAllResponeInterface {
	return &counterControlAllService{counterCountrolAll_Repo: countercontrolAllrepo}

}
