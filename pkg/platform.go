package pkg

import "strings"

const (
	PLATFORM_iOS     = "ios"
	PLATFORM_macOS   = "macos"
	PLATFORM_tvOS    = "tvos"
	PLATFORM_watchOS = "watchos"
)

func IsValidPlatform(platform string) bool {
	p := strings.ToLower(platform)
	return p == PLATFORM_iOS || p == PLATFORM_macOS || p == PLATFORM_tvOS || p == PLATFORM_watchOS
}
