package console

import (
	"github.com/CrowdSurge/banner"
	"github.com/fatih/color"
	"strings"
)

func logo() {
	banner.Print("goro")
	color.Red("		     v0.0.1")
	color.Green("POWER BY GIN\n\n")
}

func menu(args []string) {
	if len(args) > 0 {
		row := strings.ToLower(args[0])
		status, commands := checkRow(row)
		if status {
			if len(args) > 1 {
				command := strings.ToLower(args[1])
				if checkCommand(command, commands) {
					handler(row, command, args[1:])
					return
				}
			}else {
				if checkCommand("", commands) {
					handler(row, "", args[1:])
					return
				}
			}
			showCommand(row, Commends[strings.ToLower(row)])
			return
		}
	}
	showMenu()
}

func checkRow(row string) (bool, map[string]string) {
	for r, c := range Commends {
		if r == row {
			return true, c
		}
	}
	return false, nil
}

func checkCommand(command string,commands map[string]string) bool {
	for c, _ := range commands{
		if c == command || c == ""{
			return true
		}
	}
	return false
}

func showMenu() {
	color.Green("\nCOMMAND LIST \n\n")
	for t, r := range Commends {
		showCommand(t, r)
	}
}

func showCommand(t string, r map[string]string) {
	color.Blue("\n%s", strings.ToUpper(t))
	for c, d := range r {
		if c == "" {
			c = "	"
		}
		color.White("	%s		%s				%s", t, c, d)
	}
}
