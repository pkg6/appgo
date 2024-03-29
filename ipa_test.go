package appgo

import (
	"errors"
	"testing"
)

func TestIPAParsePath(t *testing.T) {
	fileName := ".test_data/ipa.ipa"
	info, err := IPAParsePath(fileName)
	if err != nil {
		t.Fatal(err)
	}
	if info == nil {
		t.Fatal(errors.New("parse error"))
	}
	t.Log(info.Name())
}
