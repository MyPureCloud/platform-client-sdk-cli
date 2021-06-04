package telephony_providers_edges_logs_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_providers_edges_logs_jobs_upload"
)

func init() {
	telephony_providers_edges_logs_jobsCmd.AddCommand(telephony_providers_edges_logs_jobs_upload.Cmdtelephony_providers_edges_logs_jobs_upload())
	telephony_providers_edges_logs_jobsCmd.Short = utils.GenerateCustomDescription(telephony_providers_edges_logs_jobsCmd.Short, telephony_providers_edges_logs_jobs_upload.Description, )
	telephony_providers_edges_logs_jobsCmd.Long = telephony_providers_edges_logs_jobsCmd.Short
}
