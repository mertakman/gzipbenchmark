package main

import (
	"bytes"
	"sync"
	"testing"
)

const parallelCount = 50

var decompressedData [][]byte

func init() {
	var err error
	if decompressedData, err = CompressTestData(); err != nil {
		panic(err)
	}
}

func BenchmarkGZipWithStdLib(b *testing.B) {
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
				CompressGzipWithStdLib(decompressedData, b)
				buffers[i].Reset()
				wg.Done()
			}(i, buffers[i])
		}
		wg.Wait()
	}
}

func BenchmarkGZipWithCompress(b *testing.B) {
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
				CompressGzipWithCompress(decompressedData, b)
				buffers[i].Reset()
				wg.Done()
			}(i, buffers[i])
		}
		wg.Wait()
	}
}

func BenchmarkGZipWithCompressPool(b *testing.B) {
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
				CompressGzipWithCompressPool(decompressedData, b)
				buffers[i].Reset()
				wg.Done()
			}(i, buffers[i])
		}
		wg.Wait()
	}
}
