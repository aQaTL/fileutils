package fileutils

import (
	"io/ioutil"
	"strings"
)

func MustLoadToString(filename string) string {
	str, err := LoadToString(filename)
	if err != nil {
		panic(err)
	}
	return string(str)
}

func LoadToString(filename string) (string, error) {
	bytes, err := ioutil.ReadFile(filename)
	return string(bytes), err
}

func MustLoadLines(filename string, windowsLineEnding bool) []string {
	lines, err := LoadLines(filename, windowsLineEnding)
	if err != nil {
		panic(err)
	}
	return lines
}

func LoadLines(filename string, windowsLineEnding bool) ([]string, error) {
	var lineSeparator string
	if windowsLineEnding {
		lineSeparator = "\r\n"
	} else {
		lineSeparator = "\n"
	}

	str, err := LoadToString(filename)
	if err != nil {
		return nil, err
	}

	return strings.Split(str, lineSeparator), nil
}
