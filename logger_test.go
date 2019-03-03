package lit

import (
	"io/ioutil"
	"testing"
)

func init() {
	Writer = ioutil.Discard
}

func BenchmarkError(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		Error("testing error")
	}
}

func BenchmarkDebug(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		Debug("testing error")
	}
}

func BenchmarkCustom(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		Custom(Writer, LogLevel, 0, "testing")
	}
}
