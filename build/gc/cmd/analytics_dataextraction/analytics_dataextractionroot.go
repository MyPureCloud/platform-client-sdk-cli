package analytics_dataextraction

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_dataextraction_downloads"
)

func init() {
	analytics_dataextractionCmd.AddCommand(analytics_dataextraction_downloads.Cmdanalytics_dataextraction_downloads())
	analytics_dataextractionCmd.Short = utils.GenerateCustomDescription(analytics_dataextractionCmd.Short, analytics_dataextraction_downloads.Description, )
	analytics_dataextractionCmd.Long = analytics_dataextractionCmd.Short
}
