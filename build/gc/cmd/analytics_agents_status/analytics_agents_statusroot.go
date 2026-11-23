package analytics_agents_status

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_agents_status_counts"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_agents_status_query"
)

func init() {
	analytics_agents_statusCmd.AddCommand(analytics_agents_status_counts.Cmdanalytics_agents_status_counts())
	analytics_agents_statusCmd.AddCommand(analytics_agents_status_query.Cmdanalytics_agents_status_query())
	analytics_agents_statusCmd.Short = utils.GenerateCustomDescription(analytics_agents_statusCmd.Short, analytics_agents_status_counts.Description, analytics_agents_status_query.Description, )
	analytics_agents_statusCmd.Long = analytics_agents_statusCmd.Short
}
