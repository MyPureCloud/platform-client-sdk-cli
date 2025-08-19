package analytics_resolutions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_resolutions_aggregates"
)

func init() {
	analytics_resolutionsCmd.AddCommand(analytics_resolutions_aggregates.Cmdanalytics_resolutions_aggregates())
	analytics_resolutionsCmd.Short = utils.GenerateCustomDescription(analytics_resolutionsCmd.Short, analytics_resolutions_aggregates.Description, )
	analytics_resolutionsCmd.Long = analytics_resolutionsCmd.Short
}
