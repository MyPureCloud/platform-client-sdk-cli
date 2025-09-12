package conversations_summaries

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_summaries_engagements"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_summaries_feedback"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_summaries_preview"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_summaries_settings"
)

func init() {
	conversations_summariesCmd.AddCommand(conversations_summaries_engagements.Cmdconversations_summaries_engagements())
	conversations_summariesCmd.AddCommand(conversations_summaries_feedback.Cmdconversations_summaries_feedback())
	conversations_summariesCmd.AddCommand(conversations_summaries_preview.Cmdconversations_summaries_preview())
	conversations_summariesCmd.AddCommand(conversations_summaries_settings.Cmdconversations_summaries_settings())
	conversations_summariesCmd.Short = utils.GenerateCustomDescription(conversations_summariesCmd.Short, conversations_summaries_engagements.Description, conversations_summaries_feedback.Description, conversations_summaries_preview.Description, conversations_summaries_settings.Description, )
	conversations_summariesCmd.Long = conversations_summariesCmd.Short
}
