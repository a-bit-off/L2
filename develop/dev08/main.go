package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	gops "github.com/mitchellh/go-ps"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах
*/

func main() {
	shell()
}

func shell() {
	stdin := bufio.NewScanner(os.Stdin)
	for {
		dir, err := os.Getwd()
		if err != nil {
			return
		}
		fmt.Print(dir, " % ")
		if stdin.Scan() {
			cmdSlice := strings.Split(stdin.Text(), "|")
			choseAction(cmdSlice)
		}
	}
}

func choseAction(actions []string) {
	for _, action := range actions {
		command := strings.Split(action, " ")
		switch command[0] {
		case "cd":
			cd(command)
		case "pwd":
			pwd()
		case "echo":
			echo(command)
		case "kill":
			kill(command)
		case "ps":
			ps()
		default:
			fork(command)
		}
	}
}

func cd(command []string) {
	if len(command) != 2 {
		return
	}

	err := os.Chdir(command[1])
	if err != nil {
		return
	}

}

func pwd() {
	dir, err := os.Getwd()
	if err != nil {
		return
	}
	fmt.Println(dir)
}

func echo(command []string) {
	size := len(command)
	if size == 1 {
		fmt.Println()
		return
	}

	for i := 1; i < size; i++ {
		fmt.Printf(command[i])
		if i+1 < size {
			fmt.Printf(" ")
		}
	}
	fmt.Println()
}

func kill(command []string) {
	if len(command) != 2 {
		fmt.Fprintln(os.Stderr, "kill: not enough arguments")
		return
	}
	pid, err := strconv.Atoi(command[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "kill: illegal pid: qsd")
		return
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	proc.Kill()
}

func ps() {
	fmt.Printf("%5s %-7s\n", "PID", "TTY")
	res, err := gops.Processes()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	for _, v := range res {
		fmt.Printf("%5d %-7s\n", v.Pid(), v.Executable())
	}
}

func fork(command []string) {
	cmd := exec.Command(command[0], command[1:]...)

	out, err := cmd.Output()
	if err != nil {
		fmt.Println("zsh: command not found:", command)
	}
	fmt.Println(string(out))
}
