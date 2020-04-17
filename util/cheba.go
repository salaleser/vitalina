package util

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
)

func GetRandomLines() string {
	dir, err := ioutil.ReadDir("lyrics")
	if err != nil {
		fmt.Fprintf(os.Stderr, "reading lyrics directory: %v\n", err)
		return err.Error()
	}

	fileName := dir[rand.Intn(len(dir))].Name()
	file, err := os.Open("lyrics/" + fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "reading lyrics file: %v\n", err)
		return err.Error()

	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Fprintf(os.Stderr, "scan lyrics lines: %v\n", err)
		return err.Error()
	}

	rnd := rand.Intn(len(lines))
	if rnd == 0 {
		rnd = 1
	}
	text := lines[rnd-1] + "\n" + lines[rnd]

	return text
}
