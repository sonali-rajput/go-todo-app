package main

// import "fmt"

func main() {
	// fmt.Println("Inside main in todo app")
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	todos.Print()
	storage.Save(todos)
}
