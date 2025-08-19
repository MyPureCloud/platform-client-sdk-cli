package speechandtextanalytics_reprocessing

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_reprocessing_jobs"
)

func init() {
	speechandtextanalytics_reprocessingCmd.AddCommand(speechandtextanalytics_reprocessing_jobs.Cmdspeechandtextanalytics_reprocessing_jobs())
	speechandtextanalytics_reprocessingCmd.Short = utils.GenerateCustomDescription(speechandtextanalytics_reprocessingCmd.Short, speechandtextanalytics_reprocessing_jobs.Description, )
	speechandtextanalytics_reprocessingCmd.Long = speechandtextanalytics_reprocessingCmd.Short
}
