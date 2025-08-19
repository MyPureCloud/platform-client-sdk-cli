package speechandtextanalytics_topics

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_topics_general"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_topics_testphrase"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_topics_publishjobs"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_topics_dialects"
)

func init() {
	speechandtextanalytics_topicsCmd.AddCommand(speechandtextanalytics_topics_general.Cmdspeechandtextanalytics_topics_general())
	speechandtextanalytics_topicsCmd.AddCommand(speechandtextanalytics_topics_testphrase.Cmdspeechandtextanalytics_topics_testphrase())
	speechandtextanalytics_topicsCmd.AddCommand(speechandtextanalytics_topics_publishjobs.Cmdspeechandtextanalytics_topics_publishjobs())
	speechandtextanalytics_topicsCmd.AddCommand(speechandtextanalytics_topics_dialects.Cmdspeechandtextanalytics_topics_dialects())
	speechandtextanalytics_topicsCmd.Short = utils.GenerateCustomDescription(speechandtextanalytics_topicsCmd.Short, speechandtextanalytics_topics_general.Description, speechandtextanalytics_topics_testphrase.Description, speechandtextanalytics_topics_publishjobs.Description, speechandtextanalytics_topics_dialects.Description, )
	speechandtextanalytics_topicsCmd.Long = speechandtextanalytics_topicsCmd.Short
}
