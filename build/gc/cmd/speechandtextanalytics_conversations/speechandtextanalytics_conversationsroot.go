package speechandtextanalytics_conversations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_conversations_summaries"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_conversations_categories"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_conversations_communications"
)

func init() {
	speechandtextanalytics_conversationsCmd.AddCommand(speechandtextanalytics_conversations_summaries.Cmdspeechandtextanalytics_conversations_summaries())
	speechandtextanalytics_conversationsCmd.AddCommand(speechandtextanalytics_conversations_categories.Cmdspeechandtextanalytics_conversations_categories())
	speechandtextanalytics_conversationsCmd.AddCommand(speechandtextanalytics_conversations_communications.Cmdspeechandtextanalytics_conversations_communications())
	speechandtextanalytics_conversationsCmd.Short = utils.GenerateCustomDescription(speechandtextanalytics_conversationsCmd.Short, speechandtextanalytics_conversations_summaries.Description, speechandtextanalytics_conversations_categories.Description, speechandtextanalytics_conversations_communications.Description, )
	speechandtextanalytics_conversationsCmd.Long = speechandtextanalytics_conversationsCmd.Short
}
