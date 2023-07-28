package domain

import (
	"strings"
)

type LocationType string

const (
	Edge  LocationType = "Edge"
	Cloud LocationType = "Cloud"
)

func ParseLocationType(location string) LocationType {
	edge := strings.ToLower(string(Edge))
	cloud := strings.ToLower(string(Cloud))

	input := strings.ToLower(strings.TrimSpace(location))

	if input == edge {
		return Edge
	} else if input == cloud {
		return Cloud
	}
	return ""
}
