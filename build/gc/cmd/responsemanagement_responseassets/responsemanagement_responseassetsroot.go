package responsemanagement_responseassets

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/responsemanagement_responseassets_status"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/responsemanagement_responseassets_uploads"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/responsemanagement_responseassets_search"
)

func init() {
	responsemanagement_responseassetsCmd.AddCommand(responsemanagement_responseassets_status.Cmdresponsemanagement_responseassets_status())
	responsemanagement_responseassetsCmd.AddCommand(responsemanagement_responseassets_uploads.Cmdresponsemanagement_responseassets_uploads())
	responsemanagement_responseassetsCmd.AddCommand(responsemanagement_responseassets_search.Cmdresponsemanagement_responseassets_search())
	responsemanagement_responseassetsCmd.Short = utils.GenerateCustomDescription(responsemanagement_responseassetsCmd.Short, responsemanagement_responseassets_status.Description, responsemanagement_responseassets_uploads.Description, responsemanagement_responseassets_search.Description, )
	responsemanagement_responseassetsCmd.Long = responsemanagement_responseassetsCmd.Short
}
