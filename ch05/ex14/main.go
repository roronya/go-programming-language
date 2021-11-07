package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	abspath, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	detects := find(abspath, os.Args[2])

	if len(detects) > 0 {
		fmt.Printf("%s is detected!\n", os.Args[2])
		for _, detect := range detects {
			fmt.Println(detect)
		}
	} else {
		fmt.Printf("%s is not detected.", os.Args[2])
	}
}

func find(root string, target string) []string {
	var detects []string

	f := func(item string) []string {
		if filepath.Base(item) == target {
			detects = append(detects, item)
			return nil
		}

		stat, err := os.Stat(item)
		if err != nil {
			log.Print(err)
		}
		if !stat.IsDir() {
			return nil
		}

		files, err := os.ReadDir(item)
		if err != nil {
			return nil
		}

		var worklist []string
		for _, file := range files {
			worklist = append(worklist, filepath.Join(item, file.Name()))
		}
		return worklist
	}

	breadthFirst(f, []string{root})
	return detects
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}
