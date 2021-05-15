package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"os"
	_ "unsafe"
)

var size = 0

func CompressTestData() ([][]byte, error) {

	testData, err := os.ReadFile("./testjson.json")
	if err != nil {
		panic(err)
	}

	testData = bytes.ReplaceAll(testData, []byte("\t"), []byte(""))
	testData = bytes.ReplaceAll(testData, []byte("\n"), []byte(""))
	testData = bytes.ReplaceAll(testData, []byte(" "), []byte(""))

	gitEvents := new(events)

	if err := json.Unmarshal(testData, &gitEvents.Events); err != nil {
		return nil, err
	}

	bs := make([][]byte, 0)

	for i := range gitEvents.Events {
		if len(gitEvents.Events[i]) < 3500 {
			bs = append(bs, gitEvents.Events[i])

			size += len(gitEvents.Events[i])
		}
	}

	return bs, nil
}

func DecompressTestData() ([][]byte, error) {

	testData, err := os.ReadFile("./testjson.json")
	if err != nil {
		panic(err)
	}

	testData = bytes.ReplaceAll(testData, []byte("\t"), []byte(""))
	testData = bytes.ReplaceAll(testData, []byte("\n"), []byte(""))
	testData = bytes.ReplaceAll(testData, []byte(" "), []byte(""))

	gitEvents := new(events)

	if err := json.Unmarshal(testData, &gitEvents.Events); err != nil {
		return nil, err
	}

	bs := make([][]byte, 0)

	for i := range gitEvents.Events {
		if len(gitEvents.Events[i]) < 3500 {
			buffer := bytes.NewBuffer(make([]byte, 0))
			writer, _ := gzip.NewWriterLevel(buffer, gzip.DefaultCompression)
			writer.Write(gitEvents.Events[i])

			writer.Flush()

			bs = append(bs, buffer.Bytes())

			buffer.Reset()
		}
	}

	return bs, nil
}

type events struct {
	Events []json.RawMessage
}
