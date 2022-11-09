package facultiesrepo

func (u *facultyAllRepoDB) GetFacultyAllData() (*[]FacultyAll, error) {
	facultyAll := []FacultyAll{}
	query := ``

	err := u.db.Select(&facultyAll, query)
	if err != nil {
		return nil, err
	}

	return &facultyAll, nil
}

func (u *facultyAllRepoDB) GetFacultyAllDataMock() (*[]FacultyAll, error) {
	facultyAll := []FacultyAll{
		{
			YEAR:     "2564",
			SEMESTER: "2",
		},
	}

	return &facultyAll, nil
}
