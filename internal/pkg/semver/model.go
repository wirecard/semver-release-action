package semver

import (
	"errors"
	"fmt"
	"strings"

	version "github.com/blang/semver"
)

var ErrInvalidIncrement = errors.New("invalid increment")

type Increment string

const (
	IncrementPatch Increment = "patch"
	IncrementMinor Increment = "minor"
	IncrementMajor Increment = "major"
)

type Version struct {
	major uint64
	minor uint64
	patch uint64
}

func (v Version) bump(inc Increment) Version {
	switch inc {
	case IncrementPatch:
		return Version{
			patch: v.patch + 1,
			minor: v.minor,
			major: v.major,
		}
	case IncrementMinor:
		return Version{
			patch: 0,
			minor: v.minor + 1,
			major: v.major,
		}
	case IncrementMajor:
		return Version{
			patch: 0,
			minor: 0,
			major: v.major + 1,
		}
	}

	return v
}

func (v Version) String() string {
	return fmt.Sprintf("v%d.%d.%d", v.major, v.minor, v.patch)
}

func ParseVersion(input string) (Version, error) {
	v, err := version.ParseTolerant(input)
	if err != nil {
		return Version{}, err
	}

	return Version{
		major: v.Major,
		minor: v.Minor,
		patch: v.Patch,
	}, nil
}

func ParseIncrement(inc string) (Increment, error) {
	switch strings.ToLower(inc) {
	case "patch":
		return IncrementPatch, nil
	case "minor":
		return IncrementMinor, nil
	case "major":
		return IncrementMajor, nil
	}

	return IncrementPatch, ErrInvalidIncrement
}
