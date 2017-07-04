package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

// TestParseResult tests the ParseClamAvOutput function.
func TestParseResult(t *testing.T) {

	r, err := ioutil.ReadFile("tests/av_scan.out")
	if err != nil {
		fmt.Print(err)
	}

	results := ParseClamAvOutput(string(r), nil)

	if true {
		t.Log("Result: ", results.Result)
		t.Log("Engine: ", results.Engine)
	}

}
