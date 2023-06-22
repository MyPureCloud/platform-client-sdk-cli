package users_presences_purecloud

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_presences_purecloud_bulk"
)

func init() {
	users_presences_purecloudCmd.AddCommand(users_presences_purecloud_bulk.Cmdusers_presences_purecloud_bulk())
	users_presences_purecloudCmd.Short = utils.GenerateCustomDescription(users_presences_purecloudCmd.Short, users_presences_purecloud_bulk.Description, )
	users_presences_purecloudCmd.Long = users_presences_purecloudCmd.Short
}
