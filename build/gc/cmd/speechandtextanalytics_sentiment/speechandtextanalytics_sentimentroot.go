package speechandtextanalytics_sentiment

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_sentiment_dialects"
)

func init() {
	speechandtextanalytics_sentimentCmd.AddCommand(speechandtextanalytics_sentiment_dialects.Cmdspeechandtextanalytics_sentiment_dialects())
	speechandtextanalytics_sentimentCmd.Short = utils.GenerateCustomDescription(speechandtextanalytics_sentimentCmd.Short, speechandtextanalytics_sentiment_dialects.Description, )
	speechandtextanalytics_sentimentCmd.Long = speechandtextanalytics_sentimentCmd.Short
}
