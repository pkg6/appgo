package appgo

import (
	"path"
	"strings"
)

const (
	AppTypeIPA     = AppType(0)
	AppTypeAPK     = AppType(1)
	AppTypeUnknown = AppType(-1)
)

func AppTypeFileName(filename string) AppType {
	ext := strings.ToLower(path.Ext(filename))
	switch ext {
	case ".ipa":
		return AppTypeIPA
	case ".apk":
		return AppTypeAPK
	default:
		return AppTypeUnknown
	}
}

type AppType int

func (t AppType) Name() string {
	switch t {
	case AppTypeIPA:
		return "ipa.ipa"
	case AppTypeAPK:
		return "apk.apk"
	default:
		return "unknown"
	}
}
