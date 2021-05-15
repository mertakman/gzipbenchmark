package main

import (
	"bytes"
	"compress/gzip"
	"io"
	"sync"

	kgzip "github.com/klauspost/compress/gzip"
)

func DecompressGzipWithStdLib(src [][]byte, dst io.Writer) {
	b := bytes.NewReader(make([]byte, 0))

	for i := range src {
		reader := new(gzip.Reader)
		b.Reset(src[i])
		reader.Reset(b)
		io.Copy(dst, reader)
	}
}

func DecompressGzipWithCompress(src [][]byte, dst io.Writer) {
	b := bytes.NewReader(make([]byte, 0))
	reader := new(kgzip.Reader)

	for i := range src {
		b.Reset(src[i])
		reader.Reset(b)
		reader.Close()
		io.Copy(dst, reader)
	}
}

var gzipReaderPool = sync.Pool{
	New: func() interface{} {
		return new(kgzip.Reader)
	},
}

func GetGzipReader(r io.Reader) *kgzip.Reader {
	return gzipReaderPool.Get().(*kgzip.Reader)
}

func PutGzipReader(zr *kgzip.Reader) {
	gzipReaderPool.Put(zr)
}

func DecompressGzipWithCompressPool(src [][]byte, dst io.Writer) {
	b := bytes.NewReader(make([]byte, 0))
	reader := GetGzipReader(b)

	for i := range src {
		b.Reset(src[i])
		reader.Reset(b)
		io.Copy(dst, reader)
	}

	PutGzipReader(reader)
}
