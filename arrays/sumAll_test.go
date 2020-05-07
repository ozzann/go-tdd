package sumAll

import (
	"reflect"
	"testing"
)

func sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	numbersCount := len(numbersToSum)
	sums := make([]int, len(numbersToSum))
	for i := 0; i < numbersCount; i++ {
		sums[i] = sum(numbersToSum[i])
	}
	return sums
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2}, []int{0, 9})
	}
}
