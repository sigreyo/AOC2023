package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	testdata := []string{}
	for scanner.Scan() {
		testdata = append(testdata, scanner.Text())
	}

	// testdata := []string{
	// 	"467..114..",
	// 	"...*......",
	// 	"..35..633.",
	// 	"......#...",
	// 	"617*......",
	// 	".....+.58.",
	// 	"..592.....",
	// 	".......755",
	// 	"...$..*...",
	// 	".664.598..",
	// }

	sum := 0
	for lineIndex, line := range testdata {
		tempNum := ""
		tempIndex := []int{}
		for i, _ := range line {
			if unicode.IsDigit(rune(line[i])) && string(line[i]) != "." {
				tempNum += string(line[i])
				if i < len(line)-1 {
					if !unicode.IsDigit(rune(line[i+1])) {
						for j := i - len(tempNum); j <= i+1; j++ {
							if j != -1 {
								tempIndex = append(tempIndex, j)
							}
						}

						if lineIndex == 0 {
							sum += checkValidNumber(tempNum, testdata[lineIndex:lineIndex+2], tempIndex)
						} else if lineIndex < len(testdata)-1 {
							sum += checkValidNumber(tempNum, testdata[lineIndex-1:lineIndex+2], tempIndex)
						} else {
							sum += checkValidNumber(tempNum, testdata[lineIndex-1:lineIndex+1], tempIndex)
						}
						tempIndex = nil
						tempNum = ""
					}
				} else {
					for j := i - len(tempNum); j <= i; j++ {
						if j != -1 {
							tempIndex = append(tempIndex, j)
						}
					}
					if lineIndex == 0 {
						sum += checkValidNumber(tempNum, testdata[lineIndex:lineIndex+2], tempIndex)
					} else if lineIndex < len(testdata)-1 {
						sum += checkValidNumber(tempNum, testdata[lineIndex-1:lineIndex+2], tempIndex)
					} else {
						sum += checkValidNumber(tempNum, testdata[lineIndex-1:lineIndex+1], tempIndex)
					}

					tempIndex = nil
					tempNum = ""
				}
			}
		}
	}
	fmt.Println(sum)
}

func checkValidNumber(word string, lines []string, indexes []int) int {
	for i := 0; i < len(lines); i++ {
		for _, v := range indexes {
			if !unicode.IsDigit(rune(lines[i][v])) && string(lines[i][v]) != "." {
				add, _ := strconv.Atoi(word)
				return add
			}
		}
	}
	return 0
}
