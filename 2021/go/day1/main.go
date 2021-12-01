package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func DepthIncreases(d []int, w int) int {
	inc := 0
	sum := 0
	for i := 0; i < w; i++ {
		sum += d[i]
	}
	for i := w; i < len(d); i++ {
		diff := d[i] - d[i-w]
		if diff > 0 {
			inc++
		}
		sum += diff
	}
	return inc
}

func ParseInputFile(filename string) ([]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	depths := []int{}

	s := bufio.NewScanner(f)
	for s.Scan() {
		d, err := strconv.Atoi(s.Text())
		if err != nil {
			return nil, err
		}
		depths = append(depths, d)
	}

	return depths, nil
}

func main() {
	depths, err := ParseInputFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(DepthIncreases(depths, 1))
	fmt.Println(DepthIncreases(depths, 3))
}
