package main

import "fmt"

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("invalid argument")
	}

	result := vals[0]
	for _, val := range vals {
		if result < val {
			result = val
		}
	}
	return result, nil
}

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("invalid argument")
	}

	result := vals[0]
	for _, val := range vals {
		if result > val {
			result = val
		}
	}
	return result, nil
}

func max2(first int, vals ...int) int {
	result := first
	for _, val := range vals {
		if result < val {
			result = val
		}
	}
	return result
}

func min2(first int, vals ...int) int {
	result := first
	for _, val := range vals {
		if result > val {
			result = val
		}
	}
	return result
}
