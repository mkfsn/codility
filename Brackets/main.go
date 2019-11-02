package main

func main() {

}

type Stack struct {
	data []string
}

func (s *Stack) Push(b string) {
	s.data = append(s.data, b)
}

func (s *Stack) Top() string {
	return s.data[len(s.data)-1]
}

func (s *Stack) Pop() {
	s.data = s.data[:len(s.data)-1]
}

func (s *Stack) Size() int {
	return len(s.data)
}

func Solution(S string) int {
	stack := &Stack{}
	pairs := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}

	for _, b := range S {
		switch b {
		case '{', '[', '(':
			stack.Push(string(b))

		case '}', ']', ')':
			if stack.Size() > 0 && stack.Top() == pairs[string(b)] {
				stack.Pop()
			} else {
				return 0
			}
		}
	}

	if stack.Size() == 0 {
		return 1
	}
	return 0
}
