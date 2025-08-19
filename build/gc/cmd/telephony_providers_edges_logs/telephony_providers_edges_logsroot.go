package telephony_providers_edges_logs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_providers_edges_logs_jobs"
)

func init() {
	telephony_providers_edges_logsCmd.AddCommand(telephony_providers_edges_logs_jobs.Cmdtelephony_providers_edges_logs_jobs())
	telephony_providers_edges_logsCmd.Short = utils.GenerateCustomDescription(telephony_providers_edges_logsCmd.Short, telephony_providers_edges_logs_jobs.Description, )
	telephony_providers_edges_logsCmd.Long = telephony_providers_edges_logsCmd.Short
}
