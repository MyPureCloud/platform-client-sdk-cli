package responsemanagement_libraries

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/responsemanagement_libraries_bulk"
)

func init() {
	responsemanagement_librariesCmd.AddCommand(responsemanagement_libraries_bulk.Cmdresponsemanagement_libraries_bulk())
	responsemanagement_librariesCmd.Short = utils.GenerateCustomDescription(responsemanagement_librariesCmd.Short, responsemanagement_libraries_bulk.Description, )
	responsemanagement_librariesCmd.Long = responsemanagement_librariesCmd.Short
}
