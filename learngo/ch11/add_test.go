package ch11

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	var dataset = []struct {
		a        int
		b        int
		expected int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{3, 2, 5},
		{4, 2, 6},
		{-1, 9, 8},
		{0, 0, 0},
		{-1, -1, -2},
	}

	for _, data := range dataset {
		actual := add(data.a, data.b)
		if actual != data.expected {
			t.Errorf("验证不通过")
		}
	}

}

func TestAdd2(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	fmt.Println("yes")
	a := add(1, 2)
	if a != 3 {
		t.Error("add(1,2) should be 3")
	}
}

func BenchmarkAdd(bb *testing.B) {
	var a, b int
	a = 123
	b = 456

	for i := 0; i < bb.N; i++ {
		_ = add(a, b)
	}
}

const numbers = 10000

func BenchmarkStringSprintf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < numbers; j++ {
			str += fmt.Sprintf("%d", j)
		}
	}
	b.StopTimer()
}

func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < numbers; j++ {
			str = str + strconv.Itoa(j)
		}
	}
	b.StopTimer()
}

func BenchmarkStringBuilder(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//var str string
		var builder strings.Builder
		for j := 0; j < numbers; j++ {
			builder.WriteString(strconv.Itoa(j))
			//str = str + strconv.Itoa(j)
		}
		_ = builder.String()
	}
	b.StopTimer()
}
