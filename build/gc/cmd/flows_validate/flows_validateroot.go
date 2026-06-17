package flows_validate

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_validate_jobs"
)

func init() {
	flows_validateCmd.AddCommand(flows_validate_jobs.Cmdflows_validate_jobs())
	flows_validateCmd.Short = utils.GenerateCustomDescription(flows_validateCmd.Short, flows_validate_jobs.Description, )
	flows_validateCmd.Long = flows_validateCmd.Short
}
