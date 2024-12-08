package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.ReadFile("in.txt")

	if err != nil {
		log.Fatal(err)
	}

	data := string(file)
	regex := regexp.MustCompile("do\\(\\)|don't\\(\\)|mul\\((\\d+),(\\d+)\\)")
	matches := regex.FindAllStringSubmatch(data, -1)

	result := 0

	enabled := true

	for _, match := range matches {
		switch match[0] {

		case "do()":
			enabled = true

		case "don't()":
			enabled = false

		default:
			if enabled {
				var num1, num2 int
				num1, err = strconv.Atoi(match[1])
				if err != nil {
					log.Fatal(err)
				}

				num2, err = strconv.Atoi(match[2])
				if err != nil {
					log.Fatal(err)
				}

				result += num1 * num2
			}
		}
	}
	fmt.Println(result)
}
