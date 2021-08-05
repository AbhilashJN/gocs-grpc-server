package accuracy_summary

import (
	gocsapi "github.com/AbhilashJN/gocs-core/api"
)

func AccuracySummaryGeneratorWrapper(demoPath string, player string) ([]string, []string, *AllAccuracySummaries) {
	resp := gocsapi.GenerateAccuracySummaryForPlayer(demoPath, player)
	accuracy_summaries := []*AccuracySummary{}
	for _, ac := range resp {
		accuracy_summary_items := []*AccuracySummaryItem{}
		for _, aci := range ac.Items {
			obj := AccuracySummaryItem{
				Name:          aci.Name,
				Fired:         int32(aci.Fired),
				Hits:          int32(aci.Hits),
				HitPercentage: aci.HitPercentage,
				Headshots:     int32(aci.Headshots),
			}
			accuracy_summary_items = append(accuracy_summary_items, &obj)
		}

		accuracy_summaries = append(accuracy_summaries, &AccuracySummary{Category: ac.Category, Items: accuracy_summary_items})
	}
	acc_summary := &AllAccuracySummaries{
		Data: accuracy_summaries,
	}
	item_fields := []string{"name", "fired", "hits", "hitPercentage", "headshots"}
	comparison_fields := []string{}
	return item_fields, comparison_fields, acc_summary
}
