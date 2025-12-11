package analytics_casemanagement

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_casemanagement_aggregates"
)

func init() {
	analytics_casemanagementCmd.AddCommand(analytics_casemanagement_aggregates.Cmdanalytics_casemanagement_aggregates())
	analytics_casemanagementCmd.Short = utils.GenerateCustomDescription(analytics_casemanagementCmd.Short, analytics_casemanagement_aggregates.Description, )
	analytics_casemanagementCmd.Long = analytics_casemanagementCmd.Short
}
