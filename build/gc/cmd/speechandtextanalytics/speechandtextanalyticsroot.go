package speechandtextanalytics

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_categories"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_conversations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_dictionaryfeedback"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_sentimentfeedback"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_programs"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_sentiment"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_settings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_topics"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_transcripts"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_translations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_reprocessing"
)

func init() {
	speechandtextanalyticsCmd.AddCommand(speechandtextanalytics_categories.Cmdspeechandtextanalytics_categories())
	speechandtextanalyticsCmd.AddCommand(speechandtextanalytics_conversations.Cmdspeechandtextanalytics_conversations())
	speechandtextanalyticsCmd.AddCommand(speechandtextanalytics_dictionaryfeedback.Cmdspeechandtextanalytics_dictionaryfeedback())
	speechandtextanalyticsCmd.AddCommand(speechandtextanalytics_sentimentfeedback.Cmdspeechandtextanalytics_sentimentfeedback())
	speechandtextanalyticsCmd.AddCommand(speechandtextanalytics_programs.Cmdspeechandtextanalytics_programs())
	speechandtextanalyticsCmd.AddCommand(speechandtextanalytics_sentiment.Cmdspeechandtextanalytics_sentiment())
	speechandtextanalyticsCmd.AddCommand(speechandtextanalytics_settings.Cmdspeechandtextanalytics_settings())
	speechandtextanalyticsCmd.AddCommand(speechandtextanalytics_topics.Cmdspeechandtextanalytics_topics())
	speechandtextanalyticsCmd.AddCommand(speechandtextanalytics_transcripts.Cmdspeechandtextanalytics_transcripts())
	speechandtextanalyticsCmd.AddCommand(speechandtextanalytics_translations.Cmdspeechandtextanalytics_translations())
	speechandtextanalyticsCmd.AddCommand(speechandtextanalytics_reprocessing.Cmdspeechandtextanalytics_reprocessing())
	speechandtextanalyticsCmd.Short = utils.GenerateCustomDescription(speechandtextanalyticsCmd.Short, speechandtextanalytics_categories.Description, speechandtextanalytics_conversations.Description, speechandtextanalytics_dictionaryfeedback.Description, speechandtextanalytics_sentimentfeedback.Description, speechandtextanalytics_programs.Description, speechandtextanalytics_sentiment.Description, speechandtextanalytics_settings.Description, speechandtextanalytics_topics.Description, speechandtextanalytics_transcripts.Description, speechandtextanalytics_translations.Description, speechandtextanalytics_reprocessing.Description, )
	speechandtextanalyticsCmd.Long = speechandtextanalyticsCmd.Short
}
