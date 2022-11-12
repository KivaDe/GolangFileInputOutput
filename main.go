package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	file_inp, err := os.ReadFile("input.txt")

	if err != nil {
		panic("cannot read file")
	}

	arr := strings.Split(strings.TrimSuffix(string(file_inp), "\n"), " ")
	n1, _ := strconv.Atoi(string(arr[0]))
	n2, _ := strconv.Atoi(string(arr[1]))

	result := n1 + n2

	data := []byte(strconv.Itoa(result))

	file_writer := os.WriteFile("output.txt", data, 0600)

	if file_writer != nil {
		panic("Cannot write file")
	}

}
