package main

import "fmt"

func durationCalculate(n int) string {

	if n <= 0 {
		fmt.Print("Durasi penerbangan harus lebih dari 0")
		return ""
	}

	var data = [6]int{1, 7, 3, 4, 8, 9}

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data); j++ {
			if (data[i] + data[j]) == n {
				fmt.Print(" film ", data[i], " & ", data[j])
				return ""
			}
		}
	}
	return ""
}
