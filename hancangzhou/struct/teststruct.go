package main

import
	(
		"fmt"
	)



type Command struct {
	Name    string    // 指令名称
	Var     *int      // 指令绑定的变量
	Comment string    // 指令的注释
	}

var version int = 1
cmd := &Command{}
cmd.Name = "version"
cmd.Var = &version
cmd.Comment = "show version"

func newCommand(name string, varref *int, comment string) *Command {
	return &Command{
	Name:    name,
	Var:     varref,
	Comment: comment,
	}
	}

cmd = newCommand(
	"version",
	&version,
	"show version",
	)
	
