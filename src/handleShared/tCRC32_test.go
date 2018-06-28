package handleShared

import (
	"testing"
)

func BenchmarkExecCRC32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExecCRC32("testing Sara")
	}
}

func BenchmarkExecMD5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExecMD5("testing Sara")
	}
}
