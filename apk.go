package appgo

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"

	"github.com/pkg6/appgo/seekbuf"
	"github.com/shogo82148/androidbinary"
	"github.com/shogo82148/androidbinary/apk"
)

func APKParsePath(file string) (*APK, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return APKParseFile(f)
}

func APKParseFile(f *os.File) (*APK, error) {
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	buf, err := seekbuf.Open(f, seekbuf.MemoryMode)
	if err != nil {
		return nil, err
	}
	defer buf.Close()
	return APKParseReader(buf, fi.Size())
}

func APKParseMultipartFile(file multipart.File) (*APK, error) {
	var buf = new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		return nil, err
	}
	sbuf, err := seekbuf.Open(buf, seekbuf.MemoryMode)
	if err != nil {
		return nil, err
	}
	return APKParseReader(sbuf, int64(buf.Len()))
}

func APKParseReader(readerAt io.ReaderAt, size int64) (*APK, error) {
	pkg, err := apk.OpenZipReader(readerAt, size)
	if err != nil {
		return nil, err
	}
	defer pkg.Close()

	icon, err := pkg.Icon(&androidbinary.ResTableConfig{
		Density: 720,
	})
	if err != nil {
		// NOTE: ignore error
	}

	return &APK{
		icon:     icon,
		manifest: pkg.Manifest(),
		size:     size,
	}, nil
}
