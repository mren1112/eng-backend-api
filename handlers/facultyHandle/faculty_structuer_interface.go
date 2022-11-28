package facultyHandle

import (
	"context"
)

var ctx = context.Background()

type ( 
	Faculty_body struct {
		FACULTY_NO         string `db:"FACULTY_NO"` 
		MAJOR_NO           string `db:"MAJOR_NO"` 
	}
)
