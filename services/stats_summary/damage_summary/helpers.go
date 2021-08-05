package damage_summary

import (
	gocsapi "github.com/AbhilashJN/gocs-core/api"
)

func DamageSummaryGeneratorWrapper(demoPath string, player string) ([]string, []string, *AllDamageSummaries) {
	resp := gocsapi.GetDamageSummaryForPlayer(demoPath, player)
	damage_summaries := []*DamageSummary{}
	for _, dmg := range resp {
		damage_summary_items := []*DamageSummaryItem{}
		for _, dmgi := range dmg.Items {
			obj := DamageSummaryItem{
				Name:  dmgi.Name,
				Given: int32(dmgi.Given),
				Taken: int32(dmgi.Taken),
			}
			damage_summary_items = append(damage_summary_items, &obj)
		}

		damage_summaries = append(damage_summaries, &DamageSummary{Category: dmg.Category, Items: damage_summary_items})
	}
	dmg_summary := &AllDamageSummaries{
		Data: damage_summaries,
	}
	item_fields := []string{"name", "given", "taken"}
	comparison_fields := []string{"given", "taken"}
	return item_fields, comparison_fields, dmg_summary
}
