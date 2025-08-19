package speechandtextanalytics_topics_general

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_topics_general_status"
)

func init() {
	speechandtextanalytics_topics_generalCmd.AddCommand(speechandtextanalytics_topics_general_status.Cmdspeechandtextanalytics_topics_general_status())
	speechandtextanalytics_topics_generalCmd.Short = utils.GenerateCustomDescription(speechandtextanalytics_topics_generalCmd.Short, speechandtextanalytics_topics_general_status.Description, )
	speechandtextanalytics_topics_generalCmd.Long = speechandtextanalytics_topics_generalCmd.Short
}
