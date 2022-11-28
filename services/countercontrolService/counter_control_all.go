package countercontrolservice

//import "fmt"

func (us *counterControlAllService) GetCounterControlAllData() (*[]CounterControlAllRespone, error) {
	responeCounterControlAll := []CounterControlAllRespone{}
	counterControlAll, err := us.counterCountrolAll_Repo.GetCounterControlAllData()

	if err != nil {
		return nil, err
	}
	for _, u := range *counterControlAll {
		responeCounterControlAll = append(responeCounterControlAll, CounterControlAllRespone{
			FISCAL_YEAR: u.FISCAL_YEAR,
			SEMESTER:    u.SEMESTER,
			YEAR:        u.YEAR,
		})
	}

	return &responeCounterControlAll, nil
}
