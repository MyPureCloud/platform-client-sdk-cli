package recording_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recording_jobs_failedrecordings"
)

func init() {
	recording_jobsCmd.AddCommand(recording_jobs_failedrecordings.Cmdrecording_jobs_failedrecordings())
	recording_jobsCmd.Short = utils.GenerateCustomDescription(recording_jobsCmd.Short, recording_jobs_failedrecordings.Description, )
	recording_jobsCmd.Long = recording_jobsCmd.Short
}
