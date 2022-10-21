package users_search

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_search_teams"
)

func init() {
	users_searchCmd.AddCommand(users_search_teams.Cmdusers_search_teams())
	users_searchCmd.Short = utils.GenerateCustomDescription(users_searchCmd.Short, users_search_teams.Description, )
	users_searchCmd.Long = users_searchCmd.Short
}
