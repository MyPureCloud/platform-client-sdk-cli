package users_search

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_search_teams"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_search_conversation"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_search_queuemembers"
)

func init() {
	users_searchCmd.AddCommand(users_search_teams.Cmdusers_search_teams())
	users_searchCmd.AddCommand(users_search_conversation.Cmdusers_search_conversation())
	users_searchCmd.AddCommand(users_search_queuemembers.Cmdusers_search_queuemembers())
	users_searchCmd.Short = utils.GenerateCustomDescription(users_searchCmd.Short, users_search_teams.Description, users_search_conversation.Description, users_search_queuemembers.Description, )
	users_searchCmd.Long = users_searchCmd.Short
}
