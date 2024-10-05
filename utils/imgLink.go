package utils

import (
	"bufio"
	"fmt"
	"os"
)

func CreateImgLink() {
	file, err := os.Create("imgLink")

	if err != nil {
		panic(err)
	}

	defer file.Close()
}

func AddImgLink(link string) {
	_, err := os.Stat("imgLink")

	if os.IsNotExist(err) {
		CreateImgLink()
		fmt.Println("Creating imgLink")
	}

	file, errO := os.OpenFile("imgLink", os.O_APPEND|os.O_WRONLY, 0644)

	if errO != nil {
		panic(errO)
	}
	
	_, errW := file.WriteString(link)

	if errW != nil {
		panic(errW)
	}
}

func GetImgLink(i int) string {
	_, statErr := os.Stat("imgLink")

	if os.IsNotExist(statErr) {
		CreateImgLink()
		fmt.Println("Creating imgLink")
	}

	file, err := os.Open("imgLink")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if scanErr := scanner.Err(); scanErr != nil {
		panic(scanErr)
	}

	if i >= 0 && i < len(lines) {
		return lines[i]
	}

	return ""
}