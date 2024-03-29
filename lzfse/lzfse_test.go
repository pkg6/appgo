package lzfse

import (
	"io/ioutil"
	"log"
	"testing"
)

func Test_lzfse(t *testing.T) {
	dat, err := ioutil.ReadFile("encoded.file")
	if err != nil {
		log.Fatal(err)
	}

	decompressed := DecodeBuffer(dat)

	err = ioutil.WriteFile("decoded.file", decompressed, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
