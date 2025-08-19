package teams

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/teams_members"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/teams_search"
)

func init() {
	teamsCmd.AddCommand(teams_members.Cmdteams_members())
	teamsCmd.AddCommand(teams_search.Cmdteams_search())
	teamsCmd.Short = utils.GenerateCustomDescription(teamsCmd.Short, teams_members.Description, teams_search.Description, )
	teamsCmd.Long = teamsCmd.Short
}
