package appgo

import (
	"bytes"
	"errors"
	"github.com/google/uuid"
	"image/png"
	"path/filepath"
	"time"
)

func NewAppInfo(pkg Package, t AppType) *AppInfo {
	return &AppInfo{
		ID:         uuid.New().String(),
		Name:       pkg.Name(),
		Version:    pkg.Version(),
		Identifier: pkg.Identifier(),
		Build:      pkg.Build(),
		Channel:    pkg.Channel(),
		Date:       time.Now(),
		Size:       pkg.Size(),
		Type:       t,
		Package:    pkg,
	}
}

type AppInfo struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Version    string    `json:"version"`
	Identifier string    `json:"identifier"`
	Build      string    `json:"build"`
	Channel    string    `json:"channel"`
	Date       time.Time `json:"date"`
	Size       int64     `json:"size"`
	Type       AppType   `json:"type"`
	Package    Package   `json:"package"`
}

//Icon
// filename := "./.test_data/ipa.ipa"
// f, err := os.Open(filename)
// if err != nil {
// 	fmt.Println(err)
// }
// defer f.Close()
// info, _ := appgo.APPInfo(f)
// filename, buf, _ := info.Icon()
// file, _ := os.Create(filename)
// defer file.Close()
// io.Copy(file, buf)
func (a *AppInfo) Icon() (string, *bytes.Buffer, error) {
	if a.Package.Icon() == nil {
		return "", nil, errors.New("unable to find ICO")
	}
	buf := &bytes.Buffer{}
	err := png.Encode(buf, a.Package.Icon())
	return filepath.Join(a.Identifier, a.ID, "icon.png"), buf, err
}

func (a *AppInfo) PackageName() string {
	return filepath.Join(a.Identifier, a.ID, a.Type.Name())
}
