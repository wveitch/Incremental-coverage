package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParse(t *testing.T) {
	byt, err := ioutil.ReadFile("gitlab.diff")
	if err != nil {
		fmt.Println(err)
	}

	q, _ := ParseDiff(string(byt))

	fmt.Println(q)
}