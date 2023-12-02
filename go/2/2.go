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

	red := 0
	green := 0
	blue := 0
	for _, line := range data {
		trimmed := strings.Split(strings.ReplaceAll(strings.Split(line, ": ")[1], ";", ","), ", ")

		red = 0
		green = 0
		blue = 0

		for _, v := range trimmed {
			amountAndColor := strings.Split(v, " ")

			switch amountAndColor[1] {
			case "red":
				checkRed, _ := strconv.Atoi(amountAndColor[0])
				if checkRed >= red {
					red = checkRed
				}
			case "green":
				checkGreen, _ := strconv.Atoi(amountAndColor[0])
				if checkGreen >= green {
					green = checkGreen
				}
			case "blue":
				checkBlue, _ := strconv.Atoi(amountAndColor[0])
				if checkBlue >= blue {
					blue = checkBlue
				}
			}
		}

		sum += red * green * blue
	}
	fmt.Println(sum)
}

// func checkForMaxAmount(amount string, max int) bool {
// 	value, _ := strconv.Atoi(amount)

// 	return value > max
// }
