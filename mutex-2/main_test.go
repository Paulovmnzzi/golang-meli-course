package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {

	stdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w

	//By default main is running in his own go routine
	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	t.Log(output)

	if !strings.Contains(output, "$34320") {
		t.Error("Wrong account balance output")
	}

}
