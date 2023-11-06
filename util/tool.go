package util

import (
	"bytes"
	"os"
)

func ReadFileToLineStr(filePath string) []string {

	var codeData []string
	dat, _ := os.ReadFile(filePath)
	lines := bytes.Split(dat, []byte("\n"))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		codeData = append(codeData, string(line))
	}
	return codeData
}
