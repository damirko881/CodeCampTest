package main

import "fmt"

type Task struct{
	Task string
}

type Employee struct{
	Name string
	Position string
	Tasks []Task
}

func (e *Employee) AddTask(taskName string) {
	e.Tasks = append(e.Tasks, Task{Task: taskName})
}

func (e* Employee) CompleteTasks(){
	fmt.Printf("Employee %s, completed tasks: %s\n", e.Name, e.Tasks)
}

func main(){
	zaposlenik1 := Employee{
		Name: "Pero Perić",
		Position: "Software developer",
		Tasks: []Task{
			{Task: "Debugiranje"},
			{Task: "Pisanje koda"},
			{Task: "Standup sastanak"},
		},
	}

	zaposlenik2 := Employee{
		Name: "Marko Markić",
		Position: "Team leader",
		Tasks: []Task{
			{Task: "Smišlja zadatke"},
			{Task: "Pisanje koda"},
			{Task: "Vođa tima"},
		},
	}

	zaposlenik1.CompleteTasks()
	zaposlenik2.CompleteTasks()
	fmt.Printf("------------------------------------------\n")

	zaposlenik1.AddTask("test")
	zaposlenik2.AddTask("test")

	zaposlenik1.CompleteTasks()
	zaposlenik2.CompleteTasks()
	fmt.Printf("------------------------------------------\n")
}