package speechandtextanalytics_transcripts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_transcripts_search"
)

func init() {
	speechandtextanalytics_transcriptsCmd.AddCommand(speechandtextanalytics_transcripts_search.Cmdspeechandtextanalytics_transcripts_search())
	speechandtextanalytics_transcriptsCmd.Short = utils.GenerateCustomDescription(speechandtextanalytics_transcriptsCmd.Short, speechandtextanalytics_transcripts_search.Description, )
	speechandtextanalytics_transcriptsCmd.Long = speechandtextanalytics_transcriptsCmd.Short
}
