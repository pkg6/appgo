package appgo

import (
	"errors"
	"testing"
)

func TestAppParsePath(t *testing.T) {
	fileName := ".test_data/ipa.ipa"
	info, err := AppParsePath(fileName)
	if err != nil {
		t.Fatal(err)
	}
	if info == nil {
		t.Fatal(errors.New("ipa.ipa parse error"))
	}
	t.Log(info.Name)
	fileName2 := ".test_data/helloworld.apk"
	info2, err := AppParsePath(fileName2)
	if err != nil {
		t.Fatal(err)
	}
	if info == nil {
		t.Fatal(errors.New("parse error"))
	}
	t.Log(info2.Name)
}
