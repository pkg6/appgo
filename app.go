package appgo

import (
	"bytes"
	"errors"
	"github.com/pkg6/appgo/seekbuf"
	"io"
	"mime/multipart"
	"os"
)

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

func AppParseMultipartFile(file multipart.File, appType AppType) (*AppInfo, error) {
	var buf = new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		return nil, err
	}
	sbuf, err := seekbuf.Open(buf, seekbuf.MemoryMode)
	if err != nil {
		return nil, err
	}
	return AppParseReader(sbuf, appType, int64(buf.Len()))
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
