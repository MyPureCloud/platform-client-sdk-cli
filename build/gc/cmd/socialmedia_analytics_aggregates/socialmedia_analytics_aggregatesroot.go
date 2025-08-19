package socialmedia_analytics_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/socialmedia_analytics_aggregates_jobs"
)

func init() {
	socialmedia_analytics_aggregatesCmd.AddCommand(socialmedia_analytics_aggregates_jobs.Cmdsocialmedia_analytics_aggregates_jobs())
	socialmedia_analytics_aggregatesCmd.Short = utils.GenerateCustomDescription(socialmedia_analytics_aggregatesCmd.Short, socialmedia_analytics_aggregates_jobs.Description, )
	socialmedia_analytics_aggregatesCmd.Long = socialmedia_analytics_aggregatesCmd.Short
}
