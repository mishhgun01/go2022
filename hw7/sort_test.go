package hw7

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

//отличия тестовых примеров :
// 1 - подаём на вход обычный слайс разных чисел
// 2 - слайс с повторяющимися элементами
// 3 - слайс с повторяющимися элементами и нестандартной записью числа
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

	array = []int{3, 03, 0, 1, 1, 4, 3}
	sort.Ints(array)
	got = array
	want = []int{0, 1, 1, 3, 3, 3, 4}
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
		{"test1", []string{"a", "c", "b"}, []string{"a", "b", "c"}},                                   // Проверяем сортировкуи букв
		{"test2", []string{"London", "london", "Aberdeen"}, []string{"Aberdeen", "London", "london"}}, // Проверяем как сортируются заглавные
		{"test3", []string{"aa", "a", ""}, []string{"", "a", "aa"}},                                   // как сортируется слайс с пустой строкой
		{"test4", []string{"3", "1", "2"}, []string{"1", "2", "3"}},                                   //как сортируются цифры в строковом виде
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
	for i := 0; i < b.N; i++ {
		data := sampleData()
		sort.Ints(data)
	}
}

//goos: linux
//goarch: amd64
//pkg: go2022/hw7
//cpu: AMD Ryzen 7 3700U with Radeon Vega Mobile Gfx
//BenchmarkSortInts
//BenchmarkSortInts-8   	       6	 183576711 ns/op

func BenchmarkSortFloat64S(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := sampleFloat64Data()
		sort.Float64s(data)
	}
}

//goos: linux
//goarch: amd64
//pkg: go2022/hw7
//cpu: AMD Ryzen 7 3700U with Radeon Vega Mobile Gfx
//BenchmarkSortFloat64S
//BenchmarkSortFloat64S-8   	       3	 342160908 ns/op

func sampleData() []int {
	rand.Seed(time.Now().UnixNano())
	var data []int
	for i := 0; i < 1_000_000; i++ {
		data = append(data, rand.Intn(1000))
	}
	return data
}

func sampleFloat64Data() []float64 {
	rand.Seed(time.Now().UnixNano())
	var data []float64
	for i := 0; i < 1_000_000; i++ {
		data = append(data, rand.Float64())
	}
	return data
}
