package outbound_campaigns

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_campaigns_agents"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_campaigns_callback"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_campaigns_divisionviews"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_campaigns_interactions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_campaigns_linedistribution"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_campaigns_progress"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_campaigns_agentownedmappingpreview"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_campaigns_diagnostics"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_campaigns_stats"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_campaigns_all"
)

func init() {
	outbound_campaignsCmd.AddCommand(outbound_campaigns_agents.Cmdoutbound_campaigns_agents())
	outbound_campaignsCmd.AddCommand(outbound_campaigns_callback.Cmdoutbound_campaigns_callback())
	outbound_campaignsCmd.AddCommand(outbound_campaigns_divisionviews.Cmdoutbound_campaigns_divisionviews())
	outbound_campaignsCmd.AddCommand(outbound_campaigns_interactions.Cmdoutbound_campaigns_interactions())
	outbound_campaignsCmd.AddCommand(outbound_campaigns_linedistribution.Cmdoutbound_campaigns_linedistribution())
	outbound_campaignsCmd.AddCommand(outbound_campaigns_progress.Cmdoutbound_campaigns_progress())
	outbound_campaignsCmd.AddCommand(outbound_campaigns_agentownedmappingpreview.Cmdoutbound_campaigns_agentownedmappingpreview())
	outbound_campaignsCmd.AddCommand(outbound_campaigns_diagnostics.Cmdoutbound_campaigns_diagnostics())
	outbound_campaignsCmd.AddCommand(outbound_campaigns_stats.Cmdoutbound_campaigns_stats())
	outbound_campaignsCmd.AddCommand(outbound_campaigns_all.Cmdoutbound_campaigns_all())
	outbound_campaignsCmd.Short = utils.GenerateCustomDescription(outbound_campaignsCmd.Short, outbound_campaigns_agents.Description, outbound_campaigns_callback.Description, outbound_campaigns_divisionviews.Description, outbound_campaigns_interactions.Description, outbound_campaigns_linedistribution.Description, outbound_campaigns_progress.Description, outbound_campaigns_agentownedmappingpreview.Description, outbound_campaigns_diagnostics.Description, outbound_campaigns_stats.Description, outbound_campaigns_all.Description, )
	outbound_campaignsCmd.Long = outbound_campaignsCmd.Short
}
