package analytics_surveys

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_surveys_aggregates"
)

func init() {
	analytics_surveysCmd.AddCommand(analytics_surveys_aggregates.Cmdanalytics_surveys_aggregates())
	analytics_surveysCmd.Short = utils.GenerateCustomDescription(analytics_surveysCmd.Short, analytics_surveys_aggregates.Description, )
	analytics_surveysCmd.Long = analytics_surveysCmd.Short
}
