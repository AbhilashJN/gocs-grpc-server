package map_details

import (
	gocsapi "github.com/AbhilashJN/gocs-core/api"
)

func GetMapName(demoPath string) *MapNameResponse {
	map_name := gocsapi.GetMapName(demoPath)
	return &MapNameResponse{MapName: map_name}
}
