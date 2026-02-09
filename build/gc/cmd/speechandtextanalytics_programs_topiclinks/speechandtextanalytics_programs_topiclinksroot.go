package speechandtextanalytics_programs_topiclinks

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_programs_topiclinks_jobs"
)

func init() {
	speechandtextanalytics_programs_topiclinksCmd.AddCommand(speechandtextanalytics_programs_topiclinks_jobs.Cmdspeechandtextanalytics_programs_topiclinks_jobs())
	speechandtextanalytics_programs_topiclinksCmd.Short = utils.GenerateCustomDescription(speechandtextanalytics_programs_topiclinksCmd.Short, speechandtextanalytics_programs_topiclinks_jobs.Description, )
	speechandtextanalytics_programs_topiclinksCmd.Long = speechandtextanalytics_programs_topiclinksCmd.Short
}
