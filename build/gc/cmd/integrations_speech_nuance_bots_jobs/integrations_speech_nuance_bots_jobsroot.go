package integrations_speech_nuance_bots_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_speech_nuance_bots_jobs_results"
)

func init() {
	integrations_speech_nuance_bots_jobsCmd.AddCommand(integrations_speech_nuance_bots_jobs_results.Cmdintegrations_speech_nuance_bots_jobs_results())
	integrations_speech_nuance_bots_jobsCmd.Short = utils.GenerateCustomDescription(integrations_speech_nuance_bots_jobsCmd.Short, integrations_speech_nuance_bots_jobs_results.Description, )
	integrations_speech_nuance_bots_jobsCmd.Long = integrations_speech_nuance_bots_jobsCmd.Short
}
