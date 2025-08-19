package speechandtextanalytics_reprocessing_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_reprocessing_jobs_interactions"
)

func init() {
	speechandtextanalytics_reprocessing_jobsCmd.AddCommand(speechandtextanalytics_reprocessing_jobs_interactions.Cmdspeechandtextanalytics_reprocessing_jobs_interactions())
	speechandtextanalytics_reprocessing_jobsCmd.Short = utils.GenerateCustomDescription(speechandtextanalytics_reprocessing_jobsCmd.Short, speechandtextanalytics_reprocessing_jobs_interactions.Description, )
	speechandtextanalytics_reprocessing_jobsCmd.Long = speechandtextanalytics_reprocessing_jobsCmd.Short
}
