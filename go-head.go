package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func getFile(fileName string) (arrOfStrings []string, length int) {
	if fileName == "-" {
		arrOfStrings = readPipe()
		length = len(arrOfStrings)
		return
	}
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	str := string(file)
	arrOfStrings = strings.Split(str, "\n")
	length = len(arrOfStrings)
	return
}

func head(file []string, num int) {
	if num > len(file) {
		num = len(file)
	}
	for i := 0; i < num; i++ {
		fmt.Println(string(file[i]))
	}
}

func readPipe() (arrayOfStrings []string) {
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err, "failed at readPipe")
	}
	str := string(stdin)
	arrayOfStrings = strings.Split(str, "\n")
	return
}

func main() {
	var lines int = 10
	var fileList []string

	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("Not enough arguments: expected at least one argument, recieved 0")
	}

	if args[0] == "-h" || args[0] == "--help" {
		fmt.Print(manPage)
		os.Exit(0)
	}

	if args[0] == "-n" {
		l, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal("if argument '-n' is used it must be followed with a number")
		}
		lines = l

		fileList = append(fileList, args[2:]...)
	} else {
		fileList = append(fileList, args...)
	}

	for _, fi := range fileList {
		file, length := getFile(fi)
		if lines < length {
			head(file, lines)
		} else {
			head(file, length)
		}

	}
	os.Exit(0)

}
