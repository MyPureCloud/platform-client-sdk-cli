package analytics_botflows

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_botflows_reportingturns"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_botflows_sessions"
)

func init() {
	analytics_botflowsCmd.AddCommand(analytics_botflows_reportingturns.Cmdanalytics_botflows_reportingturns())
	analytics_botflowsCmd.AddCommand(analytics_botflows_sessions.Cmdanalytics_botflows_sessions())
	analytics_botflowsCmd.Short = utils.GenerateCustomDescription(analytics_botflowsCmd.Short, analytics_botflows_reportingturns.Description, analytics_botflows_sessions.Description, )
	analytics_botflowsCmd.Long = analytics_botflowsCmd.Short
}
