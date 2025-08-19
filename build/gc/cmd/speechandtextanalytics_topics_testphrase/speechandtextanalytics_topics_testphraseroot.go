package speechandtextanalytics_topics_testphrase

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_topics_testphrase_jobs"
)

func init() {
	speechandtextanalytics_topics_testphraseCmd.AddCommand(speechandtextanalytics_topics_testphrase_jobs.Cmdspeechandtextanalytics_topics_testphrase_jobs())
	speechandtextanalytics_topics_testphraseCmd.Short = utils.GenerateCustomDescription(speechandtextanalytics_topics_testphraseCmd.Short, speechandtextanalytics_topics_testphrase_jobs.Description, )
	speechandtextanalytics_topics_testphraseCmd.Long = speechandtextanalytics_topics_testphraseCmd.Short
}
