package handleShared
import (
	"testing"
)
func BenchmarkExecShell(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExecShell(`ps`)
	}
}