package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	arguments := os.Args

	if len(os.Args) == 1 {
		fmt.Printf("File path not provided\n")
		os.Exit(1)
	}
	file := arguments[1]
	f, err := os.OpenFile(file, os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Printf("error opening the file %v\n", err)
		os.Exit(1)
	}
	defer f.Close()
	lines, err := ReadFile(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = WriteFile(f, lines)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ReadFile(file *os.File) ([]string, error) {
	r := bufio.NewReader(file)
	lines := make([]string, 0)
	var err error
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		if err = MatchLineString(&line); err != nil {
			break
		}
		lines = append(lines, line)
		continue
	}

	return lines, err
}

func MatchLineString(line *string) error {
	dateRegex := regexp.MustCompile(`.*\w+ \d\d \d\d:\d\d:\d\d`)
	if dateRegex.MatchString(*line) {
		match := dateRegex.FindStringSubmatch(*line)
		date, err := time.Parse("Jan 02 15:04:05", match[0])
		if err != nil {
			return err
		}
		newFormat := date.Format(time.RFC850)
		*line = strings.Replace(*line, match[0], newFormat, 1)
	}
	return nil
}

func WriteFile(file *os.File, lines []string) error {
	w := bufio.NewWriter(file)
	file.Seek(0, 0)
	for _, line := range lines {
		_, _ = w.WriteString(line)
	}
	err := w.Flush()
	if err != nil {
		return err
	}
	return nil
}
