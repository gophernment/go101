package main

func main() {
	// t := NewTodo()
	// t.Add("go shopping")
	// t.Add("go to school")
	// t.Add("pay cheque")
	// fmt.Println(t.List())
	// t.Delete(2)
	// fmt.Println(t.List())
	// t.Update(1, "go to market")
	// fmt.Println(t.List())
}

type Todo struct {
	tasks []string
}

func NewTodo() *Todo {
	return &Todo{}
}
