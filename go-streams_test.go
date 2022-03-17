package streams_test

import (
	"github.com/boreq/go-streams"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestFilter(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := streams.New(slice).
		Filter(onlyEven).
		Collect()

	require.Equal(t, []int{2, 4, 6, 8, 10}, result)
}

func TestConvolutedExample(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := streams.Map(
		streams.New(slice).Filter(onlyEven),
		intToString,
	).
		Filter(onlyOneByte).
		Collect()

	require.Equal(t, []string{"2", "4", "6", "8"}, result)
}

func BenchmarkFilterStream(b *testing.B) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var result []int

	for i := 0; i < b.N; i++ {
		result = streams.New(slice).
			Filter(onlyEven).
			Collect()
	}

	require.Equal(b, []int{2, 4, 6, 8, 10}, result)
}

func BenchmarkFilterLoop(b *testing.B) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var result []int

	for i := 0; i < b.N; i++ {
		result = nil

		for _, element := range slice {
			if element%2 == 0 {
				result = append(result, element)
			}
		}
	}

	require.Equal(b, []int{2, 4, 6, 8, 10}, result)
}

func onlyEven(v int) bool {
	return v%2 == 0
}

func onlyOneByte(v string) bool {
	return len(v) == 1
}

func intToString(v int) string {
	return strconv.Itoa(v)
}
