package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)
 

func main() {
	var v struct{}
	sizes := []int{500, 5000, 50000, 500000, 5000000}
	for _, size := range sizes {
		useSlice := make([]string, 0, size)
		useMap := make(map[string]struct{}, size)

		for i := 0; i < size; i++ {
			iStr := strconv.Itoa(i)
			useSlice = append(useSlice, iStr)
			useMap[iStr] = v
		}

		useString := strings.Join(useSlice, ", ")
		useString = fmt.Sprintf(" %s,", useString)

		target := strconv.Itoa(size - 1)
		// slice
		start := time.Now()
		for i, _ := range useSlice {
			if useSlice[i] == target {
				break
			}
		}
		fmt.Printf("use slice for size (%d): %s \n", size, time.Now().Sub(start))

		// map
		start = time.Now()
		_, exist := useMap[target]
		fmt.Printf("use map for size (%d) exist (%v): %s \n", size, exist, time.Now().Sub(start))

		start = time.Now()
		_, exist = useMap[strconv.Itoa(size)]
		fmt.Printf("use map for size (%d) exist (%v): %s \n", size, exist, time.Now().Sub(start))

		// use string contains
		targetContains := fmt.Sprintf(" %s,", target)
		start = time.Now()
		exist = strings.Contains(useString, targetContains)
		fmt.Printf("use string contains for size (%d) exist (%v): %s \n", size, exist, time.Now().Sub(start))

		targetContains = fmt.Sprintf(" %s,", strconv.Itoa(size))
		start = time.Now()
		exist = strings.Contains(useString, targetContains)
		fmt.Printf("use string contains for size (%d) exist (%v): %s \n", size, exist, time.Now().Sub(start))

		fmt.Println()
	}

}
