// Eric Rumer
// SRE Programming Test
//
// code.go:  main source file

package main

import "os"
import "fmt"
import "bufio"

// GetCommands reads each line of input from stdin and returns an array of the commands
func GetCommands(s *bufio.Scanner) (cmds []string) {
    for s.Scan() {
        cmds = append(cmds, s.Text())
	}
	return cmds
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	cmds := GetCommands(s)

	for _, cmd := range cmds {
		fmt.Println(cmd)
	}
}
