package main

import (
	"io"
	"fmt"
	"strings"

	"github.com/carmark/pseudo-terminal-go/terminal"
)

func main() {
	term, err := terminal.NewWithStdInOut()
	if err != nil {
		panic(err)
	}
	defer term.ReleaseFromStdInOut() // defer this  
	fmt.Println("Ctrl-D to break")
	term.SetPrompt("root@hello: # ")
	line, err:= term.ReadLine()
	for {
		if err == io.EOF {
			term.Write([]byte(line))
			fmt.Println()
			return
		}
		if (err != nil && strings.Contains(err.Error(), "control-c break")) || len(line) == 0{
			line, err = term.ReadLine()
		} else {
			term.Write([]byte(line+"\r\n"))
			line, err = term.ReadLine()
		}
	}
	term.Write([]byte(line))
}
