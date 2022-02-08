package speechandtextanalytics

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_conversations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_sentimentfeedback"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_sentiment"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_settings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_programs"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_topics"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_transcripts"
)

func init() {
	speechandtextanalyticsCmd.AddCommand(speechandtextanalytics_conversations.Cmdspeechandtextanalytics_conversations())
	speechandtextanalyticsCmd.AddCommand(speechandtextanalytics_sentimentfeedback.Cmdspeechandtextanalytics_sentimentfeedback())
	speechandtextanalyticsCmd.AddCommand(speechandtextanalytics_sentiment.Cmdspeechandtextanalytics_sentiment())
	speechandtextanalyticsCmd.AddCommand(speechandtextanalytics_settings.Cmdspeechandtextanalytics_settings())
	speechandtextanalyticsCmd.AddCommand(speechandtextanalytics_programs.Cmdspeechandtextanalytics_programs())
	speechandtextanalyticsCmd.AddCommand(speechandtextanalytics_topics.Cmdspeechandtextanalytics_topics())
	speechandtextanalyticsCmd.AddCommand(speechandtextanalytics_transcripts.Cmdspeechandtextanalytics_transcripts())
	speechandtextanalyticsCmd.Short = utils.GenerateCustomDescription(speechandtextanalyticsCmd.Short, speechandtextanalytics_conversations.Description, speechandtextanalytics_sentimentfeedback.Description, speechandtextanalytics_sentiment.Description, speechandtextanalytics_settings.Description, speechandtextanalytics_programs.Description, speechandtextanalytics_topics.Description, speechandtextanalytics_transcripts.Description, )
	speechandtextanalyticsCmd.Long = speechandtextanalyticsCmd.Short
}
