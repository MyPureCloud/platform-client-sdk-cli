package analytics_teams

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_teams_activity"
)

func init() {
	analytics_teamsCmd.AddCommand(analytics_teams_activity.Cmdanalytics_teams_activity())
	analytics_teamsCmd.Short = utils.GenerateCustomDescription(analytics_teamsCmd.Short, analytics_teams_activity.Description, )
	analytics_teamsCmd.Long = analytics_teamsCmd.Short
}
