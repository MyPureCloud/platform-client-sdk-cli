package speechandtextanalytics

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_topics"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_programs"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_settings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_dialects"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_transcripts"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_conversations"
)

func init() {
	speechandtextanalyticsCmd.AddCommand(speechandtextanalytics_topics.Cmdspeechandtextanalytics_topics())
	speechandtextanalyticsCmd.AddCommand(speechandtextanalytics_programs.Cmdspeechandtextanalytics_programs())
	speechandtextanalyticsCmd.AddCommand(speechandtextanalytics_settings.Cmdspeechandtextanalytics_settings())
	speechandtextanalyticsCmd.AddCommand(speechandtextanalytics_dialects.Cmdspeechandtextanalytics_dialects())
	speechandtextanalyticsCmd.AddCommand(speechandtextanalytics_transcripts.Cmdspeechandtextanalytics_transcripts())
	speechandtextanalyticsCmd.AddCommand(speechandtextanalytics_conversations.Cmdspeechandtextanalytics_conversations())
	speechandtextanalyticsCmd.Short = utils.GenerateCustomDescription(speechandtextanalyticsCmd.Short, speechandtextanalytics_topics.Description, speechandtextanalytics_programs.Description, speechandtextanalytics_settings.Description, speechandtextanalytics_dialects.Description, speechandtextanalytics_transcripts.Description, speechandtextanalytics_conversations.Description, )
	speechandtextanalyticsCmd.Long = speechandtextanalyticsCmd.Short
}
