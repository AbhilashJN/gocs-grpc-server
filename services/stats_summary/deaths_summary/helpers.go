package deaths_summary

import (
	gocsapi "github.com/AbhilashJN/gocs-core/api"
)

func DeathsSummaryGeneratorWrapper(demoPath string, player string) ([]string, []string, *AllDeathsSummaries) {
	resp := gocsapi.GetDeathsSummaryForPlayer(demoPath, player)
	deaths_summaries := []*DeathsSummary{}
	for _, de := range resp {
		deaths_summary_items := []*DeathsSummaryItem{}
		for _, dei := range de.Items {
			obj := DeathsSummaryItem{
				Name:   dei.Name,
				Kills:  int32(dei.Kills),
				Deaths: int32(dei.Deaths),
			}
			deaths_summary_items = append(deaths_summary_items, &obj)
		}

		deaths_summaries = append(deaths_summaries, &DeathsSummary{Category: de.Category, Items: deaths_summary_items})
	}
	acc_summary := &AllDeathsSummaries{
		Data: deaths_summaries,
	}
	item_fields := []string{"name", "kills", "deaths"}
	comparison_fields := []string{"kills", "deaths"}
	return item_fields, comparison_fields, acc_summary
}
