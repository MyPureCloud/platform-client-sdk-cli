package users_search_teams

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_search_teams_assign"
)

func init() {
	users_search_teamsCmd.AddCommand(users_search_teams_assign.Cmdusers_search_teams_assign())
	users_search_teamsCmd.Short = utils.GenerateCustomDescription(users_search_teamsCmd.Short, users_search_teams_assign.Description, )
	users_search_teamsCmd.Long = users_search_teamsCmd.Short
}
