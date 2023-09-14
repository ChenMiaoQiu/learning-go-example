package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"database":              {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	topsort, err := topoSort(prereqs)

	if err != nil {
		fmt.Println(err)
	} else {
		for i := 0; i < len(topsort); i++ {
			fmt.Println(i+1, ":", topsort[i])
		}
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	d := make(map[string]int)

	for val, edge := range m {
		d[val] = 0
		for _, e := range edge {
			d[e] = 0
		}
	}

	for _, edge := range m {
		for _, e := range edge {
			d[e]++
		}
	}

	q := make([]string, 0)

	for s, val := range d {
		if val == 0 {
			q = append(q, s)
		}
	}

	for i := 0; i < len(q); i++ {
		t := q[i]

		for _, val := range m[t] {
			d[val]--
			if d[val] == 0 {
				q = append(q, val)
			}
		}
	}

	if len(q) != len(d) {
		return q, fmt.Errorf("has ring")
	}
	return q, nil
}
