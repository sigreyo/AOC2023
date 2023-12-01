package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0
	for scanner.Scan() {
		lineWithNumbers := scanner.Text()

		// numbers := []int{}
		// for _, char := range lineWithNumbers {
		// 	if unicode.IsDigit(char) {
		// 		number, _ := strconv.Atoi(string(char))
		// 		numbers = append(numbers, number)
		// 	}
		// }

		numbers := checkForWordedNumbers(lineWithNumbers)

		lineSum, _ := addFirstAndLast(numbers[0], numbers[len(numbers)-1])
		sum += lineSum
	}

	fmt.Println(sum)
}

func addFirstAndLast(first, last int) (int, error) {
	return strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(last))
}

func checkForWordedNumbers(line string) []int {
	wordedNumbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	lineNumbers := []int{}

	for index, char := range line {
		if unicode.IsDigit(char) {
			value, _ := strconv.Atoi(string(char))
			lineNumbers = append(lineNumbers, value)
		} else {
			for _, word := range wordedNumbers {
				if strings.Contains(line, word) {

					for _, letter := range word {
						if char == letter {
							if len(word)+index > len(line) {
								break
							}
							if line[index:(index+len(word))] == word {
								lineNumbers = append(lineNumbers, slices.Index(wordedNumbers, word)+1)
							}
						} else {
							break
						}
					}
				}
			}
		}
	}
	return lineNumbers
}
