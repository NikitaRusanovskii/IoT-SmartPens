package domain

// Subject --------------------------------------------------------------------------------------------------------------

type Subject struct {
	SubjectID int    `db:"subject_id"`
	Name      string `db:"name"`
}

func NewSubject(gi int, nm string) Subject {
	return Subject{gi, nm}
}
func (s *Subject) SetName(nnm string) {
	s.Name = nnm
}
func (s Subject) GetID() int { return s.SubjectID }

// --------------------------------------------------------------------------------------------------------------------
