package pkg

import (
	"errors"
	"strings"
)

type Resolution struct {
	Origin   string
	Location string
	Version  string
}

var ErrParsingError = errors.New("couldn't parse the framework")

func NewResolution(line string) (*Resolution, error) {
	rows := strings.Split(line, " ")

	if len(rows) != 3 {
		return nil, ErrParsingError
	}

	location := rows[1][1 : len(rows[1])-1] // stripping out the double quotes
	version := rows[2][1 : len(rows[2])-1]

	return &Resolution{
		Origin:   rows[0],
		Location: location,
		Version:  version,
	}, nil
}

func (r *Resolution) RepoName() string {
	rows := strings.Split(r.Location, "/")
	name := rows[len(rows)-1]

	if strings.HasSuffix(name, ".git") { // Removing the .git suffix
		return name[:len(name)-4]
	}

	return name
}
