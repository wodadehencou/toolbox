package toolbox

import (
	"crypto/rand"
	"fmt"
	"testing"
)

func BenchmarkCopyConcatBytes(b *testing.B) {
	n := 2
	f := CopyConcatBytes
	b.Run(fmt.Sprintf("32B * %d", n), func(b *testing.B) { runConcatBytes(b, f, 32, n) })
	b.Run(fmt.Sprintf("1KB * %d", n), func(b *testing.B) { runConcatBytes(b, f, 1024, n) })
	b.Run(fmt.Sprintf("32KB * %d", n), func(b *testing.B) { runConcatBytes(b, f, 32*1024, n) })
	b.Run(fmt.Sprintf("1MB * %d", n), func(b *testing.B) { runConcatBytes(b, f, 1024*1024, n) })
	n = 3
	b.Run(fmt.Sprintf("32B * %d", n), func(b *testing.B) { runConcatBytes(b, f, 32, n) })
	b.Run(fmt.Sprintf("1KB * %d", n), func(b *testing.B) { runConcatBytes(b, f, 1024, n) })
	b.Run(fmt.Sprintf("32KB * %d", n), func(b *testing.B) { runConcatBytes(b, f, 32*1024, n) })
	b.Run(fmt.Sprintf("1MB * %d", n), func(b *testing.B) { runConcatBytes(b, f, 1024*1024, n) })
	n = 4
	b.Run(fmt.Sprintf("32B * %d", n), func(b *testing.B) { runConcatBytes(b, f, 32, n) })
	b.Run(fmt.Sprintf("1KB * %d", n), func(b *testing.B) { runConcatBytes(b, f, 1024, n) })
	b.Run(fmt.Sprintf("32KB * %d", n), func(b *testing.B) { runConcatBytes(b, f, 32*1024, n) })
	b.Run(fmt.Sprintf("1MB * %d", n), func(b *testing.B) { runConcatBytes(b, f, 1024*1024, n) })
}

func BenchmarkAppendConcatBytes(b *testing.B) {
	n := 2
	f := AppendConcatBytes
	b.Run(fmt.Sprintf("32B * %d", n), func(b *testing.B) { runConcatBytes(b, f, 32, n) })
	b.Run(fmt.Sprintf("1KB * %d", n), func(b *testing.B) { runConcatBytes(b, f, 1024, n) })
	b.Run(fmt.Sprintf("32KB * %d", n), func(b *testing.B) { runConcatBytes(b, f, 32*1024, n) })
	b.Run(fmt.Sprintf("1MB * %d", n), func(b *testing.B) { runConcatBytes(b, f, 1024*1024, n) })
	n = 3
	b.Run(fmt.Sprintf("32B * %d", n), func(b *testing.B) { runConcatBytes(b, f, 32, n) })
	b.Run(fmt.Sprintf("1KB * %d", n), func(b *testing.B) { runConcatBytes(b, f, 1024, n) })
	b.Run(fmt.Sprintf("32KB * %d", n), func(b *testing.B) { runConcatBytes(b, f, 32*1024, n) })
	b.Run(fmt.Sprintf("1MB * %d", n), func(b *testing.B) { runConcatBytes(b, f, 1024*1024, n) })
	n = 4
	b.Run(fmt.Sprintf("32B * %d", n), func(b *testing.B) { runConcatBytes(b, f, 32, n) })
	b.Run(fmt.Sprintf("1KB * %d", n), func(b *testing.B) { runConcatBytes(b, f, 1024, n) })
	b.Run(fmt.Sprintf("32KB * %d", n), func(b *testing.B) { runConcatBytes(b, f, 32*1024, n) })
	b.Run(fmt.Sprintf("1MB * %d", n), func(b *testing.B) { runConcatBytes(b, f, 1024*1024, n) })
}

func BenchmarkCopyConcat(b *testing.B) {
	f := CopyConcat
	len1 := 32
	len2 := 32
	b.Run(fmt.Sprintf("%d and %d", len1, len2), func(b *testing.B) { runConcat(b, f, len1, len2) })
	len1 = 1
	len2 = 32
	b.Run(fmt.Sprintf("%d and %d", len1, len2), func(b *testing.B) { runConcat(b, f, len1, len2) })
	len1 = 1
	len2 = 1024
	b.Run(fmt.Sprintf("%d and %d", len1, len2), func(b *testing.B) { runConcat(b, f, len1, len2) })
	len1 = 1
	len2 = 32 * 1024
	b.Run(fmt.Sprintf("%d and %d", len1, len2), func(b *testing.B) { runConcat(b, f, len1, len2) })
	len1 = 1
	len2 = 1024 * 1024
	b.Run(fmt.Sprintf("%d and %d", len1, len2), func(b *testing.B) { runConcat(b, f, len1, len2) })
}

func BenchmarkAppendConcat(b *testing.B) {
	f := AppendConcat
	len1 := 32
	len2 := 32
	b.Run(fmt.Sprintf("%d and %d", len1, len2), func(b *testing.B) { runConcat(b, f, len1, len2) })
	len1 = 1
	len2 = 32
	b.Run(fmt.Sprintf("%d and %d", len1, len2), func(b *testing.B) { runConcat(b, f, len1, len2) })
	len1 = 1
	len2 = 1024
	b.Run(fmt.Sprintf("%d and %d", len1, len2), func(b *testing.B) { runConcat(b, f, len1, len2) })
	len1 = 1
	len2 = 32 * 1024
	b.Run(fmt.Sprintf("%d and %d", len1, len2), func(b *testing.B) { runConcat(b, f, len1, len2) })
	len1 = 1
	len2 = 1024 * 1024
	b.Run(fmt.Sprintf("%d and %d", len1, len2), func(b *testing.B) { runConcat(b, f, len1, len2) })
}

func runConcatBytes(b *testing.B, f func(...[]byte) []byte, length, n int) {
	sl := make([][]byte, n)
	for i := range sl {
		sl[i] = make([]byte, length)
		rand.Read(sl[i])
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		o := f(sl...)
		_ = o
	}
}

func runConcat(b *testing.B, f func([]byte, []byte) []byte, len1, len2 int) {
	sl1 := make([]byte, len1)
	sl2 := make([]byte, len2)
	rand.Read(sl1)
	rand.Read(sl2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		o := f(sl1, sl2)
		_ = o
	}
}
