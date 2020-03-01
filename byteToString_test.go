package test

import (
	"fmt"
	"testing"
	"unsafe"
)

var data = []byte("HelloTatsuyaです。")

func Benchmark_String(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// []byte型をstring型にキャスト
		_ = string(data)
	}
}

func Benchmark_Sprintf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s", data)
	}
}

func Benchmark_Unsafe(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// []byte型アドレスをunsafeパッケージを使用して、一度無名型のポインタにする。
		// 無名型のポインタをstring型のポインタにキャストする。キャストしたポインタの値を取り出すことで、string型の文字列を取得する。
		// そうすることでstring(data)の時に発生したメモリコピーが発生しなくなる。
		_ = *(*string)(unsafe.Pointer(&data))
	}
}
