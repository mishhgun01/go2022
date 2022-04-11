package fileworker

import (
	"go2022/hw5/pkg/crawler"
	"io"
)

func InputInFile(writer io.Writer, data []crawler.Document) (int, error) {
	bytesRead := 0
	if data == nil {
		return 0, nil
	}
	var err error
	for _, v := range data {
		bytesRead, err = writer.Write([]byte(v.URL + "\n"))
		if err != nil {
			return 0, err
		}
	}
	return bytesRead, nil
}

func ReadFromFile(reader io.Reader) ([]string, error) {
	output := make([]byte, 64)
	var URLS []string
	for {
		_, err := reader.Read(output)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		URLS = append(URLS, string(output))
	}
	return URLS, nil
}
