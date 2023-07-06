package main

import "fmt"

func main() {
	t := NewTodo()
	t.Add("go shopping")
	t.Add("go to school")
	t.Add("pay cheque")
	fmt.Println(t.List())
	t.Delete(2)
	fmt.Println(t.List())
	t.Update(1, "go to market")
	fmt.Println(t.List())
}

func (t *Todo) Add(task string) {
	t.tasks = append(t.tasks, task)
}
func (t *Todo) Delete(id int) {
	t.tasks = append(t.tasks[:id-1], t.tasks[id:]...)
}
func (t *Todo) List() []string {
	return t.tasks
}
func (t *Todo) Update(id int, task string) {
	t.tasks[id-1] = task
}

type Todo struct {
	tasks []string
}

func NewTodo() *Todo {
	return &Todo{}
}
