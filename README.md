# go-streams

An experiment which attempts to create streams similar to what is available in Java now that we have generics.

## Should I use this library?

No.

## Should I use a different library that does this?

No.

## Examples

### Filter

Filter doesn't even look that bad to be frank.

```
func TestFilter(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := streams.New(slice).
		Filter(onlyEven).
		Collect()

	require.Equal(t, []int{2, 4, 6, 8, 10}, result)
}

func onlyEven(v int) bool {
	return v%2 == 0
}
```

### Map and filter

Map and filter looks terrible as
methods [can't be parametrized with extra type parameters](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#methods-may-not-take-additional-type-arguments)
which forced me to use a top level function.

```
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

func onlyEven(v int) bool {
	return v%2 == 0
}

func onlyOneByte(v string) bool {
	return len(v) == 1
}

func intToString(v int) string {
	return strconv.Itoa(v)
}
```

## Slow?

Yes.
```

goos: linux
goarch: amd64
pkg: github.com/boreq/go-streams
cpu: AMD Ryzen 7 3700X 8-Core Processor             
BenchmarkFilterStream
BenchmarkFilterStream-16    	 4389160	       312.2 ns/op


goos: linux
goarch: amd64
pkg: github.com/boreq/go-streams
cpu: AMD Ryzen 7 3700X 8-Core Processor             
BenchmarkFilterLoop
BenchmarkFilterLoop-16    	 6928596	       193.4 ns/op

```
