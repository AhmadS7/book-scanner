package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	WordbyWordScan()
}
func WordbyWordScan() {
	file, err := os.Open("file.txt.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		count++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d\n", count)
}
