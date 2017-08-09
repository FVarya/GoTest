package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"errors"
	"log"
)

func parseArgs(args []string) (string, error) {
	if len(args) < 2 {
		return "", errors.New("path to file not specifed")
	}

	return args[1], nil
}

func numLines(path string) (int, error) {
	bs, err := ioutil.ReadFile(path)

	if err != nil {
		return -1, err
	}

	return strings.Count(string(bs), "\n"), nil
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	path, err := parseArgs(os.Args)
	checkError(err)
	num, err := numLines(path)
	checkError(err)
	fmt.Println(num)
}