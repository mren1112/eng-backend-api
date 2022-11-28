package countercontrolrepo

func (u *counterRepoDB) GetCounterControlAllData() (*[]CounterControlAll, error) {
	counterControlAll := []CounterControlAll{}
	query := `select FISCAL_YEAR,SEMESTER,YEAR from ENG_COUNTER_ADMIN `

	err := u.db.Select(&counterControlAll, query) 
	if err != nil {
		return nil, err
	}

	return &counterControlAll, nil
}
