package colour

import (
	"testing"
	// . "github.com/smartystreets/goconvey/convey"
)

func TestTemp(t *testing.T) {
	testrun()
	testrun2()
	testrun2020()

}

func BenchmarkNRGBA64Area(b *testing.B) {
	// decode to get the colour values
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		testrun()
	}
}

func BenchmarkNRGBA64ACESSet(b *testing.B) {
	// decode to get the colour values

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		testrun2()
	}
}
