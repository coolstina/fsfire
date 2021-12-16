package fsfire

import (
	"fmt"
	"strings"

	"github.com/coolstina/expression"
)

// VersionGithub Github version number format.
type VersionGithub struct {
	Major string `json:"major"`
	Minor string `json:"minor"`
	Patch string `json:"patch"`
}

// Join version section, default use dot.
func (v *VersionGithub) Join(element ...string) string {
	return strings.Join(element, ".")
}

// GetFullVersion Get a full version string.
// Such as 1.16.6, will return 1.16.6
func (v *VersionGithub) GetFullVersion() string {
	return strings.Join([]string{v.Major, v.Minor, v.Patch}, ".")
}

// GetMinorVersion Get a minor version string.
// Such as 1.16.6, will return 1.16
func (v *VersionGithub) GetMinorVersion() string {
	return strings.Join([]string{v.Major, v.Minor}, ".")
}

// GetVersion Get Github version information in version number format.
func GetVersion(version string) (*VersionGithub, error) {
	replaced, err := expression.RegularReplaceMatch(`go`, version, "")
	if err != nil {
		return nil, err
	}

	slice := strings.Split(replaced, ".")
	if len(slice) != 3 {
		return nil, fmt.Errorf("go version invalid")
	}

	v := &VersionGithub{
		Major: slice[0],
		Minor: slice[1],
		Patch: slice[2],
	}

	return v, nil
}
