package analytics_users_observations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_users_observations_query"
)

func init() {
	analytics_users_observationsCmd.AddCommand(analytics_users_observations_query.Cmdanalytics_users_observations_query())
	analytics_users_observationsCmd.Short = utils.GenerateCustomDescription(analytics_users_observationsCmd.Short, analytics_users_observations_query.Description, )
	analytics_users_observationsCmd.Long = analytics_users_observationsCmd.Short
}
