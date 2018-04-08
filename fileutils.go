package fileutils

import "io/ioutil"

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
