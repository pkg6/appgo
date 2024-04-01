package appgo

import (
	"errors"
	"github.com/pkg6/appgo/seekbuf"
	"io"
	"os"
	"path"
	"strings"
)

const (
	AppTypeIPA     = AppType(0)
	AppTypeAPK     = AppType(1)
	AppTypeUnknown = AppType(-1)
)

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

func AppParsePath(path string) (*AppInfo, error) {
	var (
		err error
	)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return AppParseFile(file)
}

func AppParseFile(file *os.File) (*AppInfo, error) {
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	buf, err := seekbuf.Open(file, seekbuf.MemoryMode)
	if err != nil {
		return nil, err
	}
	defer buf.Close()
	return AppParseReader(buf, AppTypeFileName(file.Name()), fi.Size())
}

func AppParseReader(readerAt io.ReaderAt, appType AppType, size int64) (*AppInfo, error) {
	var (
		pkg Package
		err error
	)
	switch appType {
	case AppTypeIPA:
		pkg, err = IPAParseReader(readerAt, size)
	case AppTypeAPK:
		pkg, err = APKParseReader(readerAt, size)
	default:
		err = errors.New("unknown APP")
	}
	if err != nil {
		return nil, err
	}
	return NewAppInfo(pkg, appType), nil
}
