package main

import (
	"fmt"
	"log"
	"sort"
	"testing"
)

func TestTable_Sort(t *testing.T) {
	columns := []string{"A", "B"}
	values := [][]string{
		{"1", "3"},
		{"1", "1"},
		{"0", "2"},
		{"0", "0"},
	}

	actual, _ := NewTable(columns, values)
	sort.Sort(byKeys{actual, []int{0, 1}})

	values = [][]string{
		{"0", "0"},
		{"0", "2"},
		{"1", "1"},
		{"1", "3"},
	}
	expected, _ := NewTable(columns, values)

	/**
	log.Println(expected)
	log.Println(actual)
	*/
	for i, row := range expected.values {
		for j, _ := range row {
			if expected.values[i][j] != actual.values[i][j] {
				log.Fatalf("expected %s, but got %s\n", expected, actual)
			}
		}
	}
}

// 安定ソートでないと順番が順不同になるというのをどう示したらいいのかわからなかった
func ExampleSort() {
	columns := []string{"A", "B", "C"}
	values := [][]string{
		{"A", "B", "2"},
		{"B", "A", "5"},
		{"1", "foo", "1"},
		{"2", "bar", "4"},
	}

	t, _ := NewTable(columns, values)
	sort.Sort(byKeys{t, []int{0}})
	fmt.Println(t)
	sort.Sort(byKeys{t, []int{1}})
	fmt.Println(t)
	sort.Sort(byKeys{t, []int{2}})
	fmt.Println(t)
	// Output:
}

func ExampleStable() {
	columns := []string{"A", "B", "C"}
	values := [][]string{
		{"A", "B", "2"},
		{"B", "A", "5"},
		{"1", "foo", "1"},
		{"2", "bar", "4"},
	}

	t, _ := NewTable(columns, values)
	sort.Stable(byKeys{t, []int{0}})
	fmt.Println(t)
	sort.Stable(byKeys{t, []int{1}})
	fmt.Println(t)
	sort.Stable(byKeys{t, []int{2}})
	fmt.Println(t)
	// Output:
}
