package main

import (
	"io"
	"fmt"
	"strings"
	"terminal"
)

func main() {
	term, _ := terminal.NewWithStdInOut()
	defer term.ReleaseFromStdInOut() // defer this  
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
