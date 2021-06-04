package responsemanagement

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/responsemanagement_responses"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/responsemanagement_libraries"
)

func init() {
	responsemanagementCmd.AddCommand(responsemanagement_responses.Cmdresponsemanagement_responses())
	responsemanagementCmd.AddCommand(responsemanagement_libraries.Cmdresponsemanagement_libraries())
	responsemanagementCmd.Short = utils.GenerateCustomDescription(responsemanagementCmd.Short, responsemanagement_responses.Description, responsemanagement_libraries.Description, )
	responsemanagementCmd.Long = responsemanagementCmd.Short
}
