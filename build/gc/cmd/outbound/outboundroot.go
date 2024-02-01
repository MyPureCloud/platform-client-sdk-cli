package outbound

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_campaigns"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_attemptlimits"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_callanalysisresponsesets"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_callabletimesets"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_campaignrules"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_schedules"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_contactlists"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_contactlistfilters"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_contactlisttemplates"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_audits"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_conversations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_dnclists"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_events"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_filespecificationtemplates"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_importtemplates"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_messagingcampaigns"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_settings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_rulesets"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_sequences"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_wrapupcodemappings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_digitalrulesets"
)

func init() {
	outboundCmd.AddCommand(outbound_campaigns.Cmdoutbound_campaigns())
	outboundCmd.AddCommand(outbound_attemptlimits.Cmdoutbound_attemptlimits())
	outboundCmd.AddCommand(outbound_callanalysisresponsesets.Cmdoutbound_callanalysisresponsesets())
	outboundCmd.AddCommand(outbound_callabletimesets.Cmdoutbound_callabletimesets())
	outboundCmd.AddCommand(outbound_campaignrules.Cmdoutbound_campaignrules())
	outboundCmd.AddCommand(outbound_schedules.Cmdoutbound_schedules())
	outboundCmd.AddCommand(outbound_contactlists.Cmdoutbound_contactlists())
	outboundCmd.AddCommand(outbound_contactlistfilters.Cmdoutbound_contactlistfilters())
	outboundCmd.AddCommand(outbound_contactlisttemplates.Cmdoutbound_contactlisttemplates())
	outboundCmd.AddCommand(outbound_audits.Cmdoutbound_audits())
	outboundCmd.AddCommand(outbound_conversations.Cmdoutbound_conversations())
	outboundCmd.AddCommand(outbound_dnclists.Cmdoutbound_dnclists())
	outboundCmd.AddCommand(outbound_events.Cmdoutbound_events())
	outboundCmd.AddCommand(outbound_filespecificationtemplates.Cmdoutbound_filespecificationtemplates())
	outboundCmd.AddCommand(outbound_importtemplates.Cmdoutbound_importtemplates())
	outboundCmd.AddCommand(outbound_messagingcampaigns.Cmdoutbound_messagingcampaigns())
	outboundCmd.AddCommand(outbound_settings.Cmdoutbound_settings())
	outboundCmd.AddCommand(outbound_rulesets.Cmdoutbound_rulesets())
	outboundCmd.AddCommand(outbound_sequences.Cmdoutbound_sequences())
	outboundCmd.AddCommand(outbound_wrapupcodemappings.Cmdoutbound_wrapupcodemappings())
	outboundCmd.AddCommand(outbound_digitalrulesets.Cmdoutbound_digitalrulesets())
	outboundCmd.Short = utils.GenerateCustomDescription(outboundCmd.Short, outbound_campaigns.Description, outbound_attemptlimits.Description, outbound_callanalysisresponsesets.Description, outbound_callabletimesets.Description, outbound_campaignrules.Description, outbound_schedules.Description, outbound_contactlists.Description, outbound_contactlistfilters.Description, outbound_contactlisttemplates.Description, outbound_audits.Description, outbound_conversations.Description, outbound_dnclists.Description, outbound_events.Description, outbound_filespecificationtemplates.Description, outbound_importtemplates.Description, outbound_messagingcampaigns.Description, outbound_settings.Description, outbound_rulesets.Description, outbound_sequences.Description, outbound_wrapupcodemappings.Description, outbound_digitalrulesets.Description, )
	outboundCmd.Long = outboundCmd.Short
}
