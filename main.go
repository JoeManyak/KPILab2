package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	var path string
	fmt.Print("Введіть назву теки: ")
	_, err := fmt.Scanln(&path)
	fullInfo := takeAllFromFiles(path, err)
	fullInfo = setAllMarks(fullInfo)
	addSum(fullInfo)
	sortFullInfo(fullInfo)
	writeToFileResult(fullInfo, 10)
}

func writeToFileResult(fullInfo [][]string, count int) {
	f, err := os.Create("result.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for i := 0; i < count; i++ {
		if i >= len(fullInfo) {
			break
		}
		f.WriteString(fullInfo[i][0] + "," + fullInfo[i][len(fullInfo[i])-1] + "\n")
	}
}

func sortFullInfo(fullInfo [][]string) {
	var marks []int
	var lastIndex = len(fullInfo[0]) - 1
	for i := range fullInfo {
		data, _ := strconv.Atoi(fullInfo[i][lastIndex])
		marks = append(marks, data)
	}
	for i := 0; i < len(marks)-1; i++ {
		for j := 0; j < len(marks)-1-i; j++ {
			if marks[j] < marks[j+1] {
				marks[j], marks[j+1] = marks[j+1], marks[j]
				fullInfo[j], fullInfo[j+1] = fullInfo[j+1], fullInfo[j]
			}
		}
	}

}
func addSum(fullInfo [][]string) {
	for i := range fullInfo {
		var sum = 0
		for j := 1; j < len(fullInfo[i]); j++ {
			s, _ := strconv.Atoi(fullInfo[i][j])
			sum += s
		}
		fullInfo[i] = append(fullInfo[i], strconv.Itoa(sum))
	}
}

func setAllMarks(fullInfo [][]string) [][]string {
	for i := 1; i <= len(fullInfo); i++ {
		fullInfo = setMarks(fullInfo, i)
	}
	return fullInfo
}

func setMarks(fullInfo [][]string, col int) [][]string {
	marks := make([]int, len(fullInfo), len(fullInfo))
	for i := range fullInfo {
		marks[i], _ = strconv.Atoi(fullInfo[i][col])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(marks)))
	for i := range fullInfo {
		val, _ := strconv.Atoi(fullInfo[i][col])
		if val < marks[9] {
			fullInfo[i][col] = "0"
		} else {
			if val >= marks[9] && val < marks[1] {
				for j := 2; j < 10; j++ {
					if marks[j] == val {
						fullInfo[i][col] = strconv.Itoa(10 - j)
					}
				}
			} else {
				if marks[1] == val {
					fullInfo[i][col] = "10"
				} else {
					fullInfo[i][col] = "12"
				}
			}
		}
	}
	return fullInfo
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
