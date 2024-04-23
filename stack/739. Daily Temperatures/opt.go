package main

type stack struct {
	values []int
}

func (s *stack) push(v int) {
	s.values = append(s.values, v)
}

func (s *stack) pop() int {
	v := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return v
}

func (s *stack) get() int {
	v := s.values[len(s.values)-1]
	return v
}
func dailyTemperatures(temperatures []int) []int {

	st := stack{values: []int{}}
	return_values := make([]int, len(temperatures))

	for i, v := range temperatures {
		for len(temperatures) > 0 && v > temperatures[st.get()] {
			index := st.pop()
			return_values[index] = i - index
		}
		st.push(i)
	}

	return return_values
}
