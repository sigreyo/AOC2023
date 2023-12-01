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

	sum := 0
	for scanner.Scan() {
		lineWithNumbers := []rune(scanner.Text())
		numbers := []int{}
		for _, char := range lineWithNumbers {
			if unicode.IsDigit(char) {
				number, _ := strconv.Atoi(string(char))
				numbers = append(numbers, number)
			}
		}

		lineSum, _ := addFirstAndLast(numbers[0], numbers[len(numbers)-1])
		sum += lineSum

	}

	fmt.Println(sum)
}

func addFirstAndLast(first, last int) (int, error) {
	return strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(last))

}
