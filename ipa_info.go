package appgo

import "image"

type IPA struct {
	info *InfoPlist
	icon image.Image
	size int64
}
type InfoPlistIcon struct {
	CFBundlePrimaryIcon struct {
		CFBundleIconFiles []string `json:"CFBundleIconFiles,omitempty"`
		CFBundleIconName  string   `json:"CFBundleIconName,omitempty"`
	} `json:"CFBundlePrimaryIcon,omitempty"`
}
type InfoPlist struct {
	CFBundleDisplayName        string        `json:"CFBundleDisplayName,omitempty"`
	CFBundleExecutable         string        `json:"CFBundleExecutable,omitempty"`
	CFBundleIconName           string        `json:"CFBundleIconName,omitempty"`
	CFBundleIcons              InfoPlistIcon `json:"CFBundleIcons,omitempty"`
	CFBundleIconsIpad          InfoPlistIcon `json:"CFBundleIcons~ipad,omitempty"`
	CFBundleIdentifier         string        `json:"CFBundleIdentifier,omitempty"`
	CFBundleName               string        `json:"CFBundleName,omitempty"`
	CFBundleShortVersionString string        `json:"CFBundleShortVersionString,omitempty"`
	CFBundleSupportedPlatforms []string      `json:"CFBundleSupportedPlatforms,omitempty"`
	CFBundleVersion            string        `json:"CFBundleVersion,omitempty"`
	// not standard
	Channel string `json:"channel"`
}

func (i *IPA) Name() string {
	return def(i.info.CFBundleDisplayName, i.info.CFBundleName, i.info.CFBundleExecutable)
}

func (i *IPA) Version() string {
	return i.info.CFBundleShortVersionString
}

func (i *IPA) Identifier() string {
	return i.info.CFBundleIdentifier
}

func (i *IPA) Build() string {
	return i.info.CFBundleVersion
}

func (i *IPA) Channel() string {
	return i.info.Channel
}

func (i *IPA) Icon() image.Image {
	return i.icon
}

func (i *IPA) Size() int64 {
	return i.size
}

// def get args until arg is not empty
func def(args ...string) string {
	for _, v := range args {
		if v != "" {
			return v
		}
	}
	return ""
}
