package analytics_dataextraction_downloads

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_dataextraction_downloads_metadata"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_dataextraction_downloads_bulk"
)

func init() {
	analytics_dataextraction_downloadsCmd.AddCommand(analytics_dataextraction_downloads_metadata.Cmdanalytics_dataextraction_downloads_metadata())
	analytics_dataextraction_downloadsCmd.AddCommand(analytics_dataextraction_downloads_bulk.Cmdanalytics_dataextraction_downloads_bulk())
	analytics_dataextraction_downloadsCmd.Short = utils.GenerateCustomDescription(analytics_dataextraction_downloadsCmd.Short, analytics_dataextraction_downloads_metadata.Description, analytics_dataextraction_downloads_bulk.Description, )
	analytics_dataextraction_downloadsCmd.Long = analytics_dataextraction_downloadsCmd.Short
}
