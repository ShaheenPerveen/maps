package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getAndSplit(location string, w int, h int, zoom int, splitsize int) {
	location = parseLocation(location)
	globalImage, err := getImage(location, w, h, zoom)
	handleErr(err)

	splitImages := splitFromTop(globalImage, splitsize)

	saveSplitImages("./img/", location, splitImages)

}

func readCSV(path string) {
	file, err := os.Open(path)
	handleErr(err)

	defer file.Close()
	maxGoRoutines := 50

	guard := make(chan struct{}, maxGoRoutines)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		guard <- struct{}{}
		text := scanner.Text()
		go func() {
			fmt.Println("working on " + text)
			getAndSplit(text, 448, 468, 15, 224)
			<-guard
		}()

	}

	err = scanner.Err()
	handleErr(err)

}

func parseLocation(location string) string {
	return strings.Replace(location, " ", "+", -1)
}

func main() {
	readCSV("./data/top1000cities.csv")
}
