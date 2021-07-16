package responsemanagement

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/responsemanagement_libraries"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/responsemanagement_responses"
)

func init() {
	responsemanagementCmd.AddCommand(responsemanagement_libraries.Cmdresponsemanagement_libraries())
	responsemanagementCmd.AddCommand(responsemanagement_responses.Cmdresponsemanagement_responses())
	responsemanagementCmd.Short = utils.GenerateCustomDescription(responsemanagementCmd.Short, responsemanagement_libraries.Description, responsemanagement_responses.Description, )
	responsemanagementCmd.Long = responsemanagementCmd.Short
}
