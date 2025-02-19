package testdataprovider

import (
	"bytes"
	"compress/gzip"
	"io"
	"os"
	"sync"
)

func DecompressGz(inputPath string) ([]byte, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	gzReader, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}
	defer gzReader.Close()

	pr, pw := io.Pipe()
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer pw.Close()
		buf := make([]byte, 64*1024) // 64KB
		for {
			n, err := gzReader.Read(buf)
			if n > 0 {
				if _, writeErr := pw.Write(buf[:n]); writeErr != nil {
					break
				}
			}
			if err != nil {
				break
			}
		}
		wg.Done()
	}()

	var outBuf bytes.Buffer
	if _, err := io.Copy(&outBuf, pr); err != nil {
		return nil, err
	}
	wg.Wait()
	return outBuf.Bytes(), nil
}
