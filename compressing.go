package main

import (
	"compress/gzip"
	"io"
	"sync"

	kgzip "github.com/klauspost/compress/gzip"
)

func CompressGzipWithStdLib(src [][]byte, dst io.Writer) {
	writer, _ := gzip.NewWriterLevel(dst, gzip.DefaultCompression)
	for i := range src {
		writer.Write(src[i])
	}
	writer.Flush()
	writer.Close()
}

func CompressGzipWithCompress(src [][]byte, dst io.Writer) {
	writer, _ := kgzip.NewWriterLevel(dst, kgzip.DefaultCompression)
	for i := range src {
		writer.Write(src[i])
	}
	writer.Flush()
	writer.Close()
}

var gzipPool = sync.Pool{
	New: func() interface{} {
		writer, _ := kgzip.NewWriterLevel(nil, kgzip.DefaultCompression)
		return writer
	},
}

func GetGzipWriter() *kgzip.Writer {
	return gzipPool.Get().(*kgzip.Writer)
}

func PutGzipWriter(wr *kgzip.Writer) {
	wr.Reset(nil)

	gzipPool.Put(wr)
}

func CompressGzipWithCompressPool(src [][]byte, dst io.Writer) {
	writer := GetGzipWriter()

	writer.Reset(dst)

	for i := range src {
		writer.Write(src[i])
	}

	writer.Flush()
	writer.Close()

	PutGzipWriter(writer)
}
