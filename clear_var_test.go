package toolbox

import (
	"runtime"
	"testing"
)

func Benchmark_ClearBytes(b *testing.B) {
	bs := make([]byte, 1000*1000)
	for i := 0; i < b.N; i++ {
		ClearBytes(bs)
	}
}

func simpleClearBytes(bs []byte) {
	for i := range bs {
		bs[i] = 0
	}
	runtime.KeepAlive(bs)
}

func Benchmark_SimpleClearBytes(b *testing.B) {
	bs := make([]byte, 1000*1000)
	for i := 0; i < b.N; i++ {
		simpleClearBytes(bs)
	}
}
