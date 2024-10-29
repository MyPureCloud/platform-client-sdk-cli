package analytics

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_actions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_agentcopilots"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_bots"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_conversations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_evaluations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_flowexecutions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_flows"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_journeys"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_knowledge"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_queues"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_ratelimits"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_routing"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_surveys"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_taskmanagement"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_teams"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_transcripts"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_users"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_botflows"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_dataretention"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_reporting"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_resolutions"
)

func init() {
	analyticsCmd.AddCommand(analytics_actions.Cmdanalytics_actions())
	analyticsCmd.AddCommand(analytics_agentcopilots.Cmdanalytics_agentcopilots())
	analyticsCmd.AddCommand(analytics_bots.Cmdanalytics_bots())
	analyticsCmd.AddCommand(analytics_conversations.Cmdanalytics_conversations())
	analyticsCmd.AddCommand(analytics_evaluations.Cmdanalytics_evaluations())
	analyticsCmd.AddCommand(analytics_flowexecutions.Cmdanalytics_flowexecutions())
	analyticsCmd.AddCommand(analytics_flows.Cmdanalytics_flows())
	analyticsCmd.AddCommand(analytics_journeys.Cmdanalytics_journeys())
	analyticsCmd.AddCommand(analytics_knowledge.Cmdanalytics_knowledge())
	analyticsCmd.AddCommand(analytics_queues.Cmdanalytics_queues())
	analyticsCmd.AddCommand(analytics_ratelimits.Cmdanalytics_ratelimits())
	analyticsCmd.AddCommand(analytics_routing.Cmdanalytics_routing())
	analyticsCmd.AddCommand(analytics_surveys.Cmdanalytics_surveys())
	analyticsCmd.AddCommand(analytics_taskmanagement.Cmdanalytics_taskmanagement())
	analyticsCmd.AddCommand(analytics_teams.Cmdanalytics_teams())
	analyticsCmd.AddCommand(analytics_transcripts.Cmdanalytics_transcripts())
	analyticsCmd.AddCommand(analytics_users.Cmdanalytics_users())
	analyticsCmd.AddCommand(analytics_botflows.Cmdanalytics_botflows())
	analyticsCmd.AddCommand(analytics_dataretention.Cmdanalytics_dataretention())
	analyticsCmd.AddCommand(analytics_reporting.Cmdanalytics_reporting())
	analyticsCmd.AddCommand(analytics_resolutions.Cmdanalytics_resolutions())
	analyticsCmd.Short = utils.GenerateCustomDescription(analyticsCmd.Short, analytics_actions.Description, analytics_agentcopilots.Description, analytics_bots.Description, analytics_conversations.Description, analytics_evaluations.Description, analytics_flowexecutions.Description, analytics_flows.Description, analytics_journeys.Description, analytics_knowledge.Description, analytics_queues.Description, analytics_ratelimits.Description, analytics_routing.Description, analytics_surveys.Description, analytics_taskmanagement.Description, analytics_teams.Description, analytics_transcripts.Description, analytics_users.Description, analytics_botflows.Description, analytics_dataretention.Description, analytics_reporting.Description, analytics_resolutions.Description, )
	analyticsCmd.Long = analyticsCmd.Short
}
