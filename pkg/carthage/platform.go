package carthage

import "strings"

const (
	PLATFORM_iOS     = "iOS"
	PLATFORM_macOS   = "macOS"
	PLATFORM_tvOS    = "tvOS"
	PLATFORM_watchOS = "watchOS"
)

func IsValidPlatform(platform string) bool {
	return strings.EqualFold(platform, PLATFORM_iOS) ||
		strings.EqualFold(platform, PLATFORM_macOS) ||
		strings.EqualFold(platform, PLATFORM_tvOS) ||
		strings.EqualFold(platform, PLATFORM_watchOS)
}
