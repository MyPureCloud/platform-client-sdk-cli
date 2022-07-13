package analytics_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_users_aggregates"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_users_details"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_users_observations"
)

func init() {
	analytics_usersCmd.AddCommand(analytics_users_aggregates.Cmdanalytics_users_aggregates())
	analytics_usersCmd.AddCommand(analytics_users_details.Cmdanalytics_users_details())
	analytics_usersCmd.AddCommand(analytics_users_observations.Cmdanalytics_users_observations())
	analytics_usersCmd.Short = utils.GenerateCustomDescription(analytics_usersCmd.Short, analytics_users_aggregates.Description, analytics_users_details.Description, analytics_users_observations.Description, )
	analytics_usersCmd.Long = analytics_usersCmd.Short
}
