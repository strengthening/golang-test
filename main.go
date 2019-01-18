package main

import(
	"fmt"
)

type Task struct {
	Id int64 
	Name string
}


func main() {
	tasks := make([]Task, 0)
	task1 := Task{ 1, "haode" }
	task2 := Task{ 2, "nihao" }

	tasks = append(tasks, task1)
        tasks = append(tasks, task2)


	for _, task := range tasks {
		task_tmp := task
		fmt.Printf("%p\n", &task)
		fmt.Printf("%p\n", &task_tmp)
	}
}



