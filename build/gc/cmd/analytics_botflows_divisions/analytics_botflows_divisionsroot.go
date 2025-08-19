package analytics_botflows_divisions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_botflows_divisions_reportingturns"
)

func init() {
	analytics_botflows_divisionsCmd.AddCommand(analytics_botflows_divisions_reportingturns.Cmdanalytics_botflows_divisions_reportingturns())
	analytics_botflows_divisionsCmd.Short = utils.GenerateCustomDescription(analytics_botflows_divisionsCmd.Short, analytics_botflows_divisions_reportingturns.Description, )
	analytics_botflows_divisionsCmd.Long = analytics_botflows_divisionsCmd.Short
}
