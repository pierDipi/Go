package files

import (
	"bytes"
	"io/ioutil"
)

func SameFiles(filenames ... string) (bool, error){
	if len(filenames) < 2 {
		return true, nil
	}
	file, err := ioutil.ReadFile(filenames[0])
	if err != nil {
		return false, err
	}
	for i := 1; i < len(filenames); i++ {
		f2, err := ioutil.ReadFile(filenames[i])
		if err != nil || !bytes.Equal(file, f2) {
			return false, err
		}
	}
	return true, nil
}
