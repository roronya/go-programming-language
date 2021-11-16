package main

import (
	"fmt"
	"sort"
	"strings"
)

type Table struct {
	columns []string
	values  [][]string
}

func NewTable(columns []string, values [][]string) (*Table, error) {
	if len(values) > 0 && len(columns) != len(values[0]) {
		return nil, fmt.Errorf("invalid values. must mutch the number of column and values")
	}
	return &Table{columns, values}, nil
}

func (t *Table) String() string {
	var b strings.Builder
	b.WriteString(strings.Join(t.columns, "\t"))
	b.WriteString("\n")
	for _, v := range t.values {
		b.WriteString(strings.Join(v, "\t"))
		b.WriteString("\n")
	}
	return b.String()
}

type byKeys struct {
	t        *Table
	sortkeys []int // クリックしたカラムの順序を保持して比較時に使う
}

// ソートする対象は列なので列の数を返す
func (bk byKeys) Len() int { return len(bk.t.values) }

// ソートするロジックについては特に指定がないので固定したロジックにした
func (bk byKeys) Less(i, j int) bool {
	x, y := bk.t.values[i], bk.t.values[j]
	for _, k := range bk.sortkeys {
		if x[k] != y[k] {
			return x[k] < y[k]
		}
	}
	return false
}
func (bk byKeys) Swap(i, j int) { bk.t.values[i], bk.t.values[j] = bk.t.values[j], bk.t.values[i] }

func main() {
	t := Table{
		columns: []string{"c_a", "c_b", "c_b"},
		values: [][]string{
			{"v_a_0", "v_b_0", "v_c_0"},
			{"v_a_1", "v_b_1", "v_c_1"},
			{"v_a_2", "v_b_2", "v_c_2"},
		},
	}
	sort.Sort(byKeys{&t, []int{0, 1, 2}})
	fmt.Println(t.String())
}
