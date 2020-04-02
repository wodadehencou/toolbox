package toolbox

import "testing"
import "crypto/elliptic"
import "crypto/rand"

func Benchmark_BigToSlice(b *testing.B) {
	_, x, _, _ := elliptic.GenerateKey(elliptic.P256(), rand.Reader)
	bs := make([]byte, 32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BigToSlice(x, bs)
	}
}
