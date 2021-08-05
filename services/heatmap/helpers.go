package heatmap

import (
	gocsapi "github.com/AbhilashJN/gocs-core/api"
	"github.com/golang/geo/r2"
)

func pointsConverter(input []r2.Point) []*HeatmapPoint {
	output := make([]*HeatmapPoint, len(input))
	for i, pt := range input {
		output[i] = &HeatmapPoint{X: pt.X, Y: pt.Y}
	}
	return output
}

func PlayersListWrapper(demoPath string, player string) *HeatmapResponse {
	heatmapPoints := gocsapi.GetHeatMapPositions(demoPath, player)
	killPoints := pointsConverter(heatmapPoints["kill"])
	deathPoints := pointsConverter(heatmapPoints["death"])
	bombPlantPoints := pointsConverter(heatmapPoints["bomb_plant"])
	response := &HeatmapResponse{
		Kill:      killPoints,
		Death:     deathPoints,
		BombPlant: bombPlantPoints,
	}
	return response
}
