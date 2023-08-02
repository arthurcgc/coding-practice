package stack

type Stack interface {
	Push(elem interface{})
	Pop() interface{}
	Peek() interface{}
}

type StringStack struct {
	innerStack []string
}

func (ss *StringStack) Push(elem string) {
	ss.innerStack = append(ss.innerStack, elem)
}

func (ss *StringStack) Pop() string {
	n := len(ss.innerStack)
	if n == 0 {
		return ""
	}
	s := ss.innerStack[n-1]
	ss.innerStack = ss.innerStack[:n-1] // slice ranges are not incluse on the left bound. i.e: [)
	return s
}

func (ss *StringStack) Peek() string {
	if ss.innerStack == nil || len(ss.innerStack) == 0 {
		return "\u0000"
	}
	return ss.innerStack[len(ss.innerStack)-1]
}

func (ss *StringStack) Len() int {
	return len(ss.innerStack)
}

type IntStack struct {
	innerStack []int
}

func (is *IntStack) Push(elem int) {
	is.innerStack = append(is.innerStack, elem)
}

func (is *IntStack) Pop() int {
	n := len(is.innerStack)
	if n == 0 {
		return -1
	}
	v := is.innerStack[n-1]
	is.innerStack = is.innerStack[:n-1] // slice ranges are not incluse on the left bound. i.e: [)
	return v
}

func (is *IntStack) Peek() int {
	return is.innerStack[len(is.innerStack)-1]
}
