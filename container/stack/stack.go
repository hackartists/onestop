package stack

type Stack struct {
	Items []interface{}
	count int
}

func New() *Stack {
	return &Stack{Items: make([]interface{}, 0), count: 0}
}

func (s *Stack) Push(item interface{}) {
	s.Items = append(s.Items, item)
	s.count++
}

func (s *Stack) Pop() {
	if s.count > 0 {
		s.Items = s.Items[0 : s.count-1]
		s.count--
	}
}

func (s *Stack) Top() interface{} {
	if s.count > 0 && len(s.Items) >= s.count {
		return s.Items[s.count-1]
	}

	return nil
}
