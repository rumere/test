// Eric Rumer
// SRE Programming Test
//
// code_test.go:  unit test source file

package main

import "testing"
import "bufio"
import "strings"




func TestGetCommands(t *testing.T) {
	// Test data sets
	var TestInputs = []string {
		"TEST 1\nTEST 2\nTEST3\n",
		"TEST 3\n",
		"TEST 1\n, TEST2\n",
	}

	for _, TestInput := range TestInputs {
		s := bufio.NewScanner(strings.NewReader(TestInput))
		cmds := GetCommands(s)
		var TestResult string
		for _, cmd := range cmds {
			TestResult += (cmd + "\n")
		}
		if strings.Compare(TestResult, TestInput) != 0 {
			t.Errorf("GetCommands did not process input correctly:\nEXPECTED:\n", TestInput, "\nRESULT:\n", TestResult)
		}
	}
}