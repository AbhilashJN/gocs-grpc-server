package players_list

import (
	gocsapi "github.com/AbhilashJN/gocs-core/api"
)

func PlayersListWrapper(demoPath string) *PlayersListResponse {
	response := &PlayersListResponse{Names: gocsapi.ListPlayers(demoPath)}
	return response

}
