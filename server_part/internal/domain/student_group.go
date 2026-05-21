package domain

// StudentGroup -------------------------------------------------------------------------------------------------------

type StudentGroup struct {
	GroupID int    `db:"group_id"`
	Name    string `db:"name"`
}

func NewStudentGroup(gi int, nm string) StudentGroup {
	return StudentGroup{gi, nm}
}
func (s *StudentGroup) SetName(nnm string) {
	s.Name = nnm
}
func (s StudentGroup) GetID() int { return s.GroupID }

// --------------------------------------------------------------------------------------------------------------------
