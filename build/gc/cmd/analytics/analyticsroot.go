package analytics

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_reporting"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_users"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_flows"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_conversations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_queues"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_evaluations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_bots"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_journeys"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_surveys"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_transcripts"
)

func init() {
	analyticsCmd.AddCommand(analytics_reporting.Cmdanalytics_reporting())
	analyticsCmd.AddCommand(analytics_users.Cmdanalytics_users())
	analyticsCmd.AddCommand(analytics_flows.Cmdanalytics_flows())
	analyticsCmd.AddCommand(analytics_conversations.Cmdanalytics_conversations())
	analyticsCmd.AddCommand(analytics_queues.Cmdanalytics_queues())
	analyticsCmd.AddCommand(analytics_evaluations.Cmdanalytics_evaluations())
	analyticsCmd.AddCommand(analytics_bots.Cmdanalytics_bots())
	analyticsCmd.AddCommand(analytics_journeys.Cmdanalytics_journeys())
	analyticsCmd.AddCommand(analytics_surveys.Cmdanalytics_surveys())
	analyticsCmd.AddCommand(analytics_transcripts.Cmdanalytics_transcripts())
	analyticsCmd.Short = utils.GenerateCustomDescription(analyticsCmd.Short, analytics_reporting.Description, analytics_users.Description, analytics_flows.Description, analytics_conversations.Description, analytics_queues.Description, analytics_evaluations.Description, analytics_bots.Description, analytics_journeys.Description, analytics_surveys.Description, analytics_transcripts.Description, )
	analyticsCmd.Long = analyticsCmd.Short
}
