package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	var s string
	fmt.Print("Введіть назву теки: ")
	_, err := fmt.Scanln(&s)
	fileArr := getArrOfFiles(s)
	if err != nil {
		log.Fatal(err)
	}
	readLinesFromCSV(s, "eurovision1.csv")
	for i := range fileArr {
		fmt.Println(fileArr[i])
	}
}

func getArrOfFiles(path string) []string {
	var fileArr []string
	err := filepath.Walk("./"+path, func(path string, info os.FileInfo, err error) error {
		for i := utf8.RuneCountInString(path) - 1; i >= 0; i-- {
			if path[i] == 92 {
				path = path[i+1:]
			}
		}
		fileArr = append(fileArr, path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return fileArr
}

func readLinesFromCSV(path, filename string) [][]string {
	file, err := os.Open("./" + path + "/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	n := findN(scanner)
	return readLines(n, scanner)
}

func readLines(n int, scanner *bufio.Scanner) [][]string {
	lines := make([][]string, n, n)
	for i := 0; scanner.Scan(); i++ {
		lines[i] = strings.Split(scanner.Text(), ",")
	}
	return lines
}

func findN(scanner *bufio.Scanner) int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
