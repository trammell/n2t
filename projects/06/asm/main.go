package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := regexp.MustCompile(`//.*`).ReplaceAllString(scanner.Text(), "")
		txt = regexp.MustCompile(`\s`).ReplaceAllString(txt, "")
		if len(txt) > 0 {
			fmt.Fprintf(os.Stderr, "> %s\n", txt)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func isLabel(inst string) bool {
	return regexp.MustCompile(`^(.*)`).MatchString(inst)
}

func isAInstruction(inst string) bool {
	return regexp.MustCompile(`^@.*`).MatchString(inst)
}

func isCInstruction(inst string) bool {
	return regexp.MustCompile(`[ADM]*=?([^;]*);())`).MatchString(inst)
}
