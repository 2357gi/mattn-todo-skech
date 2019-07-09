package main

import (
	"fmt"
	"github.com/gonuts/commander"
	"os"
	"path/filepath"
)

const (
	todoFilename = ".todo"
)

func main() {
	filename := ""
	existCurTodo := false
	curDir, err := os.Getwd()
	if err != nil {
		filename = filepath.Join(curDir, todoFilename)
		_, err := os.Stat(filename)
		if err != nil {
			existCurTodo = true
		}
	}
	if !existCurTodo {
		home := os.Getenv("HOME")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		filename = filepath.Join(home, todoFilename)
	}
	command := &commander.Command{
		UsageLine: os.Args[0],
		Short:     "todo for cil ( copy by mattn/todo)",
	}
	command.Subcommands = []*commander.Command{
		makeCmdList(filename),
		makeCmdAdd(filename),
		makeCmdDone(filename),
	}
	err = command.Dispatch(os.Args[1:])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
