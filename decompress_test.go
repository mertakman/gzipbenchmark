package main

import (
	"bytes"
	"sync"
	"testing"
)

var compressedData [][]byte

func init() {
	var err error
	if compressedData, err = DecompressTestData(); err != nil {
		panic(err)
	}
}

func BenchmarkDecompressWithStdLib(b *testing.B) {
	b.StopTimer()
	wg := sync.WaitGroup{}
	buffers := make([]*bytes.Buffer, parallelCount)
	for i := range buffers {
		buffers[i] = bytes.NewBuffer(make([]byte, 0))
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		wg.Add(parallelCount)
		for i := range buffers {
			go func(i int, b *bytes.Buffer) {
				DecompressGzipWithStdLib(compressedData, b)
				buffers[i].Reset()
				wg.Done()
			}(i, buffers[i])
		}
		wg.Wait()
	}
}

func BenchmarkDecompressWithCompress(b *testing.B) {
	b.StopTimer()
	wg := sync.WaitGroup{}
	buffers := make([]*bytes.Buffer, parallelCount)
	for i := range buffers {
		buffers[i] = bytes.NewBuffer(make([]byte, 0))
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		wg.Add(parallelCount)
		for i := range buffers {
			go func(i int, b *bytes.Buffer) {
				DecompressGzipWithCompress(compressedData, b)
				buffers[i].Reset()
				wg.Done()
			}(i, buffers[i])
		}
		wg.Wait()
	}
}

func BenchmarkDecompressWithCompressPool(b *testing.B) {
	b.StopTimer()
	wg := sync.WaitGroup{}
	buffers := make([]*bytes.Buffer, parallelCount)
	for i := range buffers {
		buffers[i] = bytes.NewBuffer(make([]byte, 0))
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		wg.Add(parallelCount)
		for i := range buffers {
			go func(i int, b *bytes.Buffer) {
				DecompressGzipWithCompressPool(compressedData, b)
				buffers[i].Reset()
				wg.Done()
			}(i, buffers[i])
		}
		wg.Wait()
	}
}
