package analytics_reporting_exports

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_reporting_exports_metadata"
)

func init() {
	analytics_reporting_exportsCmd.AddCommand(analytics_reporting_exports_metadata.Cmdanalytics_reporting_exports_metadata())
	analytics_reporting_exportsCmd.Short = utils.GenerateCustomDescription(analytics_reporting_exportsCmd.Short, analytics_reporting_exports_metadata.Description, )
	analytics_reporting_exportsCmd.Long = analytics_reporting_exportsCmd.Short
}
