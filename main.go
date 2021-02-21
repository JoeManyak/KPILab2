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
	var path string
	fmt.Print("Введіть назву теки: ")
	_, err := fmt.Scanln(&path)
	fullInfo := takeAllFromFiles(path, err)
	for i := range fullInfo {
		fmt.Println(fullInfo[i])
	}
}

func takeAllFromFiles(path string, err error) [][]string {
	fileArr := getArrOfFiles(path)
	if err != nil {
		log.Fatal(err)
	}
	var fullInfo [][]string
	for _, v := range fileArr {
		appendLines(path, v, &fullInfo)
	}
	return fullInfo
}

func appendLines(path string, v string, fullInfo *[][]string) {
	linesFromCSV := readLinesFromCSV(path, v)
	temp := make([][]string, len(*fullInfo)+len(linesFromCSV), len(*fullInfo)+len(linesFromCSV))
	copy(temp, *fullInfo)
	copy(temp[len(*fullInfo):], linesFromCSV)
	*fullInfo = temp
}

func getArrOfFiles(p string) []string {
	var fileArr []string
	err := filepath.Walk("./"+p, func(path string, info os.FileInfo, err error) error {
		if path != "./"+p {
			for i := utf8.RuneCountInString(path) - 1; i >= 0; i-- {
				if path[i] == 92 {
					path = path[i+1:]
				}
			}
			fileArr = append(fileArr, path)
		}
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
