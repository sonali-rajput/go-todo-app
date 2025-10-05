package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlag struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlag {
	cf := CmdFlag{}
	flag.StringVar(&cf.Add, "add", "", "Add a new todo specify title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit todo by index & specify a new title. id:new_title")
	flag.IntVar(&cf.Del, "del", -1, "specify todo by index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "specify todo by index to toggle")
	flag.BoolVar(&cf.List, "list", false, "List all the todos")

	flag.Parse()

	return &cf
}

func (cf *CmdFlag) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.Print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, Invalid format for edit. Please use id:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error: invalid index for edit")
		}

		todos.Edit(index, parts[1])

	case cf.Toggle != -1:
		todos.Toggle(cf.Toggle)
	case cf.Del != -1:
		todos.Delete(cf.Del)
	default:
		fmt.Println("Invalid command")
	}
}
