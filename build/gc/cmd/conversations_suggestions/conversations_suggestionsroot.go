package conversations_suggestions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_suggestions_engagement"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_suggestions_feedback"
)

func init() {
	conversations_suggestionsCmd.AddCommand(conversations_suggestions_engagement.Cmdconversations_suggestions_engagement())
	conversations_suggestionsCmd.AddCommand(conversations_suggestions_feedback.Cmdconversations_suggestions_feedback())
	conversations_suggestionsCmd.Short = utils.GenerateCustomDescription(conversations_suggestionsCmd.Short, conversations_suggestions_engagement.Description, conversations_suggestions_feedback.Description, )
	conversations_suggestionsCmd.Long = conversations_suggestionsCmd.Short
}
