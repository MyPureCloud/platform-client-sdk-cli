package analytics_teams_activity

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_teams_activity_query"
)

func init() {
	analytics_teams_activityCmd.AddCommand(analytics_teams_activity_query.Cmdanalytics_teams_activity_query())
	analytics_teams_activityCmd.Short = utils.GenerateCustomDescription(analytics_teams_activityCmd.Short, analytics_teams_activity_query.Description, )
	analytics_teams_activityCmd.Long = analytics_teams_activityCmd.Short
}
