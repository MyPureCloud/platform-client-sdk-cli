package speechandtextanalytics_topics

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_topics_publishjobs"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_topics_general"
)

func init() {
	speechandtextanalytics_topicsCmd.AddCommand(speechandtextanalytics_topics_publishjobs.Cmdspeechandtextanalytics_topics_publishjobs())
	speechandtextanalytics_topicsCmd.AddCommand(speechandtextanalytics_topics_general.Cmdspeechandtextanalytics_topics_general())
	speechandtextanalytics_topicsCmd.Short = utils.GenerateCustomDescription(speechandtextanalytics_topicsCmd.Short, speechandtextanalytics_topics_publishjobs.Description, speechandtextanalytics_topics_general.Description, )
	speechandtextanalytics_topicsCmd.Long = speechandtextanalytics_topicsCmd.Short
}
