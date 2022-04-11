package hw7

import (
	"math/rand"
	"sort"
	"strings"
	"testing"
	"time"
)

func Test_SortInts(t *testing.T) {
	array := []int{2, 42, 1, 46, 7, 0, 9}
	sort.Ints(array)
	got := array
	want := []int{0, 1, 2, 7, 9, 42, 46}
	for i, v := range want {
		if got[i] != v {
			t.Fatalf("Test #1 - got: %+v, want: %+v", got[i], want)
		}
	}
	array = []int{2, 2, 1, 1, 46, 7, 0, 9}
	sort.Ints(array)
	got = array
	want = []int{0, 1, 1, 2, 2, 7, 9, 46}
	for i, v := range want {
		if got[i] != v {
			t.Fatalf("Test #2 - got: %+v, want: %+v", got[i], want)
		}
	}

	array = []int{9, 8, 7, 6, 5, 4, 3}
	sort.Ints(array)
	got = array
	want = []int{3, 4, 5, 6, 7, 8, 9}
	for i, v := range want {
		if got[i] != v {
			t.Fatalf("Test #3 - got: %+v, want: %+v", got[i], want)
		}
	}
}

func Test_SortStrings(t *testing.T) {
	var s []string
	tests := []struct {
		name string
		s    []string
		want []string
	}{
		{"test1", []string{"a", "c", "b"}, []string{"a", "b", "c"}},
		{"test2", []string{"London", "london", "Aberdeen"}, []string{"Aberdeen", "London", "london"}},
		{"test3", []string{"aa", "a", ""}, []string{"", "a", "aa"}},
		{"test4", []string{"3", "1", "2"}, []string{"1", "2", "3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s = tt.s
			sort.Strings(s)
			got := s
			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("%v: got = %v, want=%v", tt.name, v, tt.want[i])
				}
			}
		})
	}
}

func BenchmarkSortInts(b *testing.B) {
	data := sampleData()
	sort.Ints(data)
}

//результат для данного алгоритма - 0.1684 ns/op при сложности О(n)

func BenchmarkSortStrings(b *testing.B) {
	data := sampleStringData()
	sort.Strings(data)
}

//результат для данного алгоритма - 3497976445 ns/op при сложности O(n^2)

func sampleData() []int {
	rand.Seed(time.Now().UnixNano())
	var data []int
	for i := 0; i < 1_000_000; i++ {
		data = append(data, rand.Intn(1000))
	}
	return data
}

func sampleStringData() []string {
	rand.Seed(time.Now().UnixNano())
	var s []string
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 8
	var b strings.Builder
	for j := 0; j < 100000; j++ {
		for i := 0; i < length; i++ {
			b.WriteRune(chars[rand.Intn(len(chars))])
		}
		str := b.String()
		s = append(s, str)
		str = ""
	}
	return s
}
