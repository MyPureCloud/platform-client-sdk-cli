package users_presences

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_presences_purecloud"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_presences_bulk"
)

func init() {
	users_presencesCmd.AddCommand(users_presences_purecloud.Cmdusers_presences_purecloud())
	users_presencesCmd.AddCommand(users_presences_bulk.Cmdusers_presences_bulk())
	users_presencesCmd.Short = utils.GenerateCustomDescription(users_presencesCmd.Short, users_presences_purecloud.Description, users_presences_bulk.Description, )
	users_presencesCmd.Long = users_presencesCmd.Short
}
