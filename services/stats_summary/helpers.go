package stats_summary

import (
	accuracy_summary "github.com/AbhilashJN/gocs-grpc-server/services/stats_summary/accuracy_summary"
	damage_summary "github.com/AbhilashJN/gocs-grpc-server/services/stats_summary/damage_summary"
	deaths_summary "github.com/AbhilashJN/gocs-grpc-server/services/stats_summary/deaths_summary"
)

func StatsSummaryWrapper(demoPath string, player string, statType string) *GetStatsSummaryResponse {
	var response *GetStatsSummaryResponse
	switch statType {
	case "Damage":
		item_fields, comparison_fields, data := damage_summary.DamageSummaryGeneratorWrapper(demoPath, player)
		response = &GetStatsSummaryResponse{
			ItemFields:       item_fields,
			ComparisonFields: comparison_fields,
			Data: &GetStatsSummaryResponse_AllDamageSummaries{
				AllDamageSummaries: data},
		}
	case "Deaths/Kills":
		item_fields, comparison_fields, data := deaths_summary.DeathsSummaryGeneratorWrapper(demoPath, player)
		response = &GetStatsSummaryResponse{
			ItemFields:       item_fields,
			ComparisonFields: comparison_fields,
			Data: &GetStatsSummaryResponse_AllDeathsSummaries{
				AllDeathsSummaries: data},
		}
	case "Accuracy":
		item_fields, comparison_fields, data := accuracy_summary.AccuracySummaryGeneratorWrapper(demoPath, player)
		response = &GetStatsSummaryResponse{
			ItemFields:       item_fields,
			ComparisonFields: comparison_fields,
			Data: &GetStatsSummaryResponse_AllAccuracySummaries{
				AllAccuracySummaries: data},
		}
	default:
		response = &GetStatsSummaryResponse{}
	}
	return response
}
