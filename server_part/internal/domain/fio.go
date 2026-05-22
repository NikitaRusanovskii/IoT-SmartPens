package domain

// FIO struct ---------------------------------------------------------------------------------------------------------

type Fio struct {
	FirstName  string `db:"fname" json:"first_name"`
	MiddleName string `db:"mname" json:"middle_name"`
	LastName   string `db:"lname" json:"last_name"`
}

// --------------------------------------------------------------------------------------------------------------------
