package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readLinesFromCSV()
	for i := range lines {
		fmt.Println(lines[i])
	}
}

func readLinesFromCSV() [][]string {
	file, err := os.Open("./data/eurovision1.csv")
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
