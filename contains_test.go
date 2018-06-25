package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

var (
	size = 500000

	target    int
	targetStr string

	exMap   map[int]struct{}
	exSlice []int
	exStr   string
)

func testMap() {
	_ = exMap[target]
}

func testSlice() {
	for _, v := range exSlice {
		if v == target {
			break
		}
	}
}

func testStrngContains() {
	strings.Contains(exStr, targetStr)
}

func init() {
	target = rand.Intn(size)
	targetStr = fmt.Sprintf(",%d,", target)

	var v struct{}
	exMap = make(map[int]struct{}, size)
	exSlice = make([]int, 0, size)
	str := make([]string, 0, size)

	for i := 0; i < size; i++ {
		exMap[i] = v
		exSlice = append(exSlice, i)
		str = append(str, strconv.Itoa(i))
	}

	exStr = fmt.Sprintf(",%s,", strings.Join(str, ","))
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testMap()
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testSlice()
	}
}

func BenchmarkStringContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testStrngContains()
	}
}
