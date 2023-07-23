package iterables

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	want := 15
	if got != want {
		t.Errorf("got:%d wanted:%d, given %v", got, want, numbers)
	}
}

func TestReturnSumArrays(t *testing.T) {
	got := ReturnSumArray([]int{1, 2, 3}, []int{1, 9, 5})
	want := []int{6, 15}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{9, 1, 3})
		want := []int{5, 4}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{5, 7, 3})
		want := []int{0, 10}
		checkSums(t, got, want)
	})

}
