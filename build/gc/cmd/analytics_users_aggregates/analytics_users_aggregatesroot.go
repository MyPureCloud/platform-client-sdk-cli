package analytics_users_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_users_aggregates_query"
)

func init() {
	analytics_users_aggregatesCmd.AddCommand(analytics_users_aggregates_query.Cmdanalytics_users_aggregates_query())
	analytics_users_aggregatesCmd.Short = utils.GenerateCustomDescription(analytics_users_aggregatesCmd.Short, analytics_users_aggregates_query.Description, )
	analytics_users_aggregatesCmd.Long = analytics_users_aggregatesCmd.Short
}
