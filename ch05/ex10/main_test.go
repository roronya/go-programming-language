package main

import (
	"log"
	"testing"
)

// prereqsの依存関係を整理してその通りになっているか確認した
func TestTopoSort(t *testing.T) {
	actual := topoSort(prereqs)

	orderMap := map[string]int{}
	for i, v := range actual {
		orderMap[v] = i
	}
	if !(orderMap["calculus"] > orderMap["linear algebra"]) {
		log.Fatalf("linear algebra must be begger than calculus:\ngot %#v", actual)
	}
	if !(orderMap["discrete math"] > orderMap["intro to programming"]) {
		log.Fatalf("discrete math must be begger than intro to progrraming:\ngot %#v", actual)
	}
	if !(orderMap["formal languages"] > orderMap["discrete math"]) {
		log.Fatalf("error!:\ngot %#v", actual)
	}
	if !(orderMap["algorithms"] > orderMap["data structures"]) {
		log.Fatalf("error!:\ngot %#v", actual)
	}
	if !(orderMap["databases"] > orderMap["data structures"]) {
		log.Fatalf("error!:\ngot %#v", actual)
	}
	if !(orderMap["compilers"] > orderMap["data structures"]) {
		log.Fatalf("error!:\ngot %#v", actual)
	}
	if !(orderMap["compilers"] > orderMap["formal languages"]) {
		log.Fatalf("error!:\ngot %#v", actual)
	}
	if !(orderMap["compilers"] > orderMap["computer organizations"]) {
		log.Fatalf("error!:\ngot %#v", actual)
	}
	if !(orderMap["operating systems"] > orderMap["computer organizations"]) {
		log.Fatalf("error!:\ngot %#v", actual)
	}
	if !(orderMap["operating systems"] > orderMap["data structures"]) {
		log.Fatalf("error!:\ngot %#v", actual)
	}
	if !(orderMap["networks"] > orderMap["operating systems"]) {
		log.Fatalf("error!:\ngot %#v", actual)
	}
	if !(orderMap["programming languages"] > orderMap["data structures"]) {
		log.Fatalf("error!:\ngot %#v", actual)
	}
	if !(orderMap["programming languages"] > orderMap["computer organizations"]) {
		log.Fatalf("error!:\ngot %#v", actual)
	}
}
