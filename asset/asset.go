package asset

import (
	"fmt"
	"github.com/pkg6/appgo/bytess"
)

// tag: 'CTAR'
type CarHeader struct {
	// uint32_t tag;
	Tag bytess.String4
	// uint32_t coreuiVersion;
	CoreuiVersion uint32
	// uint32_t storageVersion;
	StorageVersion uint32
	// uint32_t storageTimestamp;
	StorageTimestamp uint32
	// uint32_t renditionCount;
	RenditionCount uint32
	// char mainVersionString[128];
	MainVersionString bytess.String128
	// char versionString[256];
	VersionString bytess.String256
	// uuid_t uuid;
	UUID bytess.String16
	// uint32_t associatedChecksum;
	AssociatedChecksum uint32
	// uint32_t schemaVersion;
	SchemaVersion uint32
	// uint32_t colorSpaceID;
	ColorSpaceID uint32
	// uint32_t keySemantics;
	KeySemantics uint32
}

// tag: 'META'
type CarextendedMetadata struct {
	// uint32_t tag;
	Tag bytess.String4
	// char thinningArguments[256];
	ThinningArguments bytess.String256
	// char deploymentPlatformVersion[256];
	DeploymentPlatformVersion bytess.String256
	// char deploymentPlatform[256];
	DeploymentPlatform bytess.String256
	// char authoringTool[256];
	AuthoringTool bytess.String256
}

// tag: 'kfmt'
type RenditionKeyFmt struct {
	// uint32_t tag;
	Tag bytess.String4
	// uint32_t version;
	Version uint32
	// uint32_t maximumRenditionKeyTokenCount;
	MaximumRenditionKeyTokenCount uint32
	// uint32_t renditionKeyTokens[];
	RenditionKeyTokens []RenditionAttributeType
}

func (r *RenditionKeyFmt) Keys() []string {
	l := make([]string, len(r.RenditionKeyTokens))
	for i, v := range r.RenditionKeyTokens {
		l[i] = v.String()
	}
	return l
}

// tag: 'FACETKEYS'
type Renditionkeytoken struct {
	CursorHotSpot struct {
		X uint16
		Y uint16
	}
	NumberOfAttributes uint16
	Attributes         []RenditionAttribute
}

type RenditionAttribute struct {
	Name  uint16
	Value uint16hex
}

type uint16hex uint16

func (u uint16hex) String() string {
	return fmt.Sprintf("%04X", uint16(u))
}
