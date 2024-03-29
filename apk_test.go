package appgo

import (
	"errors"
	"testing"
)

func TestAPKParsePath(t *testing.T) {
	fileName := ".test_data/helloworld.apk"
	info, err := APKParsePath(fileName)
	if err != nil {
		t.Fatal(err)
	}
	if info == nil {
		t.Fatal(errors.New("parse error"))
	}
	t.Log(info.Name())
}
