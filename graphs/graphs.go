package main

type stack struct {
	values []string
}

func (s *stack) push(value string) {
	s.values = append(s.values, value)
}

func (s *stack) pop() string {
	poped := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return poped
}

func dfs_traversal(graph map[string][]string, firstElem string) []string {

	result := []string{}
	s := stack{
		values: []string{firstElem},
	}
	for {
		if len(s.values) == 0 {
			break
		}
		value := s.pop()
		result = append(result, value)
		if len(graph[value]) == 0 {
			continue
		}
		for _, v := range graph[value] {
			s.push(v)
		}

	}
	return result
}

type queue struct {
	values []string
}

func (q *queue) enqueue(value string) {
	q.values = append(q.values, value)
}

func (q *queue) dequeue() string {
	if len(q.values) == 0 {
		return ""
	} else if len(q.values) == 1 {
		val := q.values[0]
		q.values = []string{}
		return val
	}
	que := q.values[0]
	q.values = q.values[1:]

	return que
}

func bfs_traversal(graph map[string][]string, initial string) []string {
	result := []string{}
	q := queue{
		values: []string{initial},
	}

	for {
		if len(q.values) == 0 {
			break
		}
		val := q.dequeue()
		result = append(result, val)
		values := graph[val]
		for _, v := range values {
			q.enqueue(v)
		}
	}

	return result
}

func has_path(graph map[string][]string, src, dest string) bool {
	s := stack{
		values: []string{src},
	}
	for {
		if len(s.values) == 0 {
			break
		}
		value := s.pop()
		if value == dest {
			return true
		}
		if len(graph[value]) == 0 {
			continue
		}
		for _, v := range graph[value] {
			s.push(v)
		}

	}
	return false
}
