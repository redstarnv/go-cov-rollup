package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var examples = []string{
	"atomic", "set",
}

func runTestExample(t *testing.T, name string) {
	in, err := os.Open("examples/" + name + "-in.txt")
	if err != nil {
		t.Fatalf("%s", err)
	}
	defer in.Close()

	var out bytes.Buffer
	run(in, &out)

	expected, err := ioutil.ReadFile("examples/" + name + "-out.txt")
	if err != nil {
		t.Fatalf("%s", err)
	}

	if strings.Compare(out.String(), string(expected)) != 0 {
		fmt.Println(name + " example received output that does not match expected output")
		fmt.Printf("Expected:\n------\n%s\n------\n\n", expected)
		fmt.Printf("Received:\n------\n%s\n------\n\n", out.String())
		t.FailNow()
	}
}

func TestMain(t *testing.T) {
	runTestExample(t, "atomic")
	runTestExample(t, "set")
}
