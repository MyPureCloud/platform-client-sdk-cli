package analytics_dataretention

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_dataretention_settings"
)

func init() {
	analytics_dataretentionCmd.AddCommand(analytics_dataretention_settings.Cmdanalytics_dataretention_settings())
	analytics_dataretentionCmd.Short = utils.GenerateCustomDescription(analytics_dataretentionCmd.Short, analytics_dataretention_settings.Description, )
	analytics_dataretentionCmd.Long = analytics_dataretentionCmd.Short
}
