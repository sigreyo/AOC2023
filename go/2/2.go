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
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0
	data := []string{}
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	for index, line := range data {
		trimmed := strings.Split(strings.ReplaceAll(strings.Split(line, ": ")[1], ";", ","), ", ")

		valid := true
		for _, v := range trimmed {
			amountAndColor := strings.Split(v, " ")

			switch amountAndColor[1] {
			case "red":
				if checkForMaxAmount(amountAndColor[0], 12) {
					valid = false
					break
				}
			case "green":
				if checkForMaxAmount(amountAndColor[0], 13) {
					valid = false
					break
				}
			case "blue":
				if checkForMaxAmount(amountAndColor[0], 14) {
					valid = false
					break
				}
			}
		}
		if valid {
			sum += index + 1
		}

	}
	fmt.Println(sum)
}

func checkForMaxAmount(amount string, max int) bool {
	value, _ := strconv.Atoi(amount)

	return value > max
}
