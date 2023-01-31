package workspace

type Student struct {
	Name  string
	Age   int
	Grade float64
}

func (s *Student) GetName() string {
	return s.Name
}
