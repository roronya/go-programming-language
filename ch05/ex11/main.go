package main

import (
	"fmt"
	"log"
)

// valueは依存関係のある講義はtrueがセットされているmapとする
var prereqs = map[string]map[string]bool{
	"algorithms": {"data structures": true},
	"calculus":   {"linear algebra": true},
	"compilers": {
		"data structures":       true,
		"formal languages":      true,
		"computer organization": true,
	},
	"data structures":       {"discrete math": true},
	"databases":             {"data structures": true},
	"discrete math":         {"intro to programming": true},
	"formal languages":      {"discrete math": true},
	"networks":              {"operating systems": true},
	"operating systems":     {"data structures": true, "computer organization": true},
	"programming languages": {"data structures": true, "computer organization": true},
	"linear algebra":        {"calculus": true}, // これが循環構造を作るので報告するようにする
}

func main() {
	courses, err := topoSort(prereqs)
	if err != nil {
		log.Fatal(err)
	}
	for i, course := range courses {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string]map[string]bool) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	tmp := make(map[string]bool)
	var visitAll func(items map[string]bool) error

	visitAll = func(items map[string]bool) error {
		for item, hasEdge := range items {
			if hasEdge && tmp[item] {
				return fmt.Errorf("acyclic path is detected. %s cannot be have edge\n", item)
			}
			if hasEdge && !tmp[item] && !seen[item] {
				tmp[item] = true
				err := visitAll(m[item])
				if err != nil {
					return err
				}
				order = append(order, item)
			}
		}
		return nil
	}

	// 各ノードから探索を始める
	// 各探索で訪れるノードは一時的に色を付けていって、探索から抜けたあとに恒久的に訪れないようにする
	// 一時的に色を付けることで、その探索の中で循環したかどうかを検出できる
	for k, _ := range m {
		err := visitAll(m[k])
		if err != nil {
			return nil, err
		}
		for k, v := range tmp {
			seen[k] = seen[k] || v
			tmp[k] = false
		}
	}
	return order, nil
}
